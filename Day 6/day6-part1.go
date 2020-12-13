package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	surveys := LoadAnswersFromFile()

	count := 0
	for _, survey := range surveys {
		count += len(survey)
	}
	log.Printf("Sum of answer count: %v", count)
}

func LoadAnswersFromFile() [][]rune {
	fptr := flag.String("fpath", "answers", "./")
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

	var surveys [][]rune
	var survey []rune = nil

	s := bufio.NewScanner(f)
	for s.Scan() {
		if strings.TrimSpace(s.Text()) == "" {
			// Empty line; Close survey if open, and proceed to the next line
			if survey != nil {
				surveys = append(surveys, survey)
				survey = nil
			}
			continue
		}

		if survey == nil {
			survey = make([]rune, 0)
		}

		for _, r := range []rune(s.Text()) {
			if !contains(survey, r) {
				survey = append(survey, r)
			}
		}
	}

	if survey != nil {
		surveys = append(surveys, survey)
		survey = nil
	}

	return surveys
}

func contains(array []rune, search rune) bool {
	for _, r := range array {
		if r == search {
			return true
		}
	}

	return false
}
