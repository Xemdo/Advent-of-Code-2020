package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	values := readValuesFromFile()

	// Run tests
	for a := 0; a < len(values); a++ {
		for b := 0; b < len(values); b++ {
			sum := values[a] + values[b]
			if sum == 2020 {
				log.Printf("Values %v + %v = 2020", values[a], values[b])
				log.Printf("%v * %v = %v", values[a], values[b], values[a]*values[b])
				return
			}
		}
	}
}

func readValuesFromFile() []int {
	var values []int

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
		a, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, a)
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return values
}
