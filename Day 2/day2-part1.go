package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	values := readValuesFromFile()

	reg := regexp.MustCompile(`^(\d{1,})-(\d{1,}) (\w): (\w+)`)

	validPasswords := 0

	for _, line := range values {
		r := reg.FindStringSubmatch(line)
		passRangeMin, _ := strconv.Atoi(r[1])
		passRangeMax, _ := strconv.Atoi(r[2])
		passRequire := r[3]
		password := r[4]

		reg2 := regexp.MustCompile(passRequire)
		count := len(reg2.FindAllStringSubmatch(password, -1))
		if count >= passRangeMin && count <= passRangeMax {
			validPasswords++
			//log.Printf("[%v, %v, %v, %v]", r[1], r[2], r[3], r[4])
		}
	}

	log.Printf("Total number of valid passwords: %v", validPasswords)
}

func readValuesFromFile() []string {
	var values []string

	fptr := flag.String("fpath", "input", "./")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		values = append(values, s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return values
}
