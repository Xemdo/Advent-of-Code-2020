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
	for a := 0; a < len(values); a++ { // Grab first number
		for b := 0; b < len(values); b++ { // Grab second number
			for c := 0; c < len(values); c++ { // Check third number
				sum := values[a] + values[b] + values[c]
				if sum == 2020 {
					log.Printf("Values %v + %v + %v = 2020", values[a], values[b], values[c])
					log.Printf("%v * %v * %v = %v", values[a], values[b], values[c], values[a]*values[b]*values[c])
					return
				}
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
