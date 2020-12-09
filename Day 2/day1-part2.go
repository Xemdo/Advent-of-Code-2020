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
		passRequire := []rune(r[3])[0]
		password := r[4]

		chars := []rune(password)

		if (chars[passRangeMin-1] == passRequire || chars[passRangeMax-1] == passRequire) && chars[passRangeMin-1] != chars[passRangeMax-1] {
			validPasswords++
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
