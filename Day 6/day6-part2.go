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

	answerSum := 0

	for _, group := range surveys { // Go through each group
		if len(group) == 1 { // Only one person
			answerSum += len(group[0])
			continue
		}

		// Check first person, and then check each of those answers again each person
		answers := []rune(group[0])
		for _, r := range answers {
			notFound := false

			// Checking each person now
			for i, person := range group {
				if i == 0 { // Skip if first person
					continue
				}

				if !contains(person, r) {
					notFound = true
					continue
				}
			}

			if !notFound {
				answerSum++
			}
		}
	}

	log.Printf("Sum: %v", answerSum)
}

func LoadAnswersFromFile() [][][]rune {
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

	// This is one of the laziest use of 3D arrays I've ever done
	// [ ][ ][x] = Answer for a person; One rune
	// [ ][x][ ] = Array of answers for a person
	// [x][ ][ ] = Array of people; One group
	var surveys [][][]rune = make([][][]rune, 0)

	group := 0
	person := 0

	s := bufio.NewScanner(f)
	for s.Scan() {

		if strings.TrimSpace(s.Text()) == "" {
			// Empty line; Close group; Reset person counter
			group++
			person = 0
			continue
		}

		if len(surveys)-1 < group {
			newGroup := make([][]rune, 0)
			surveys = append(surveys, newGroup)
		}

		surveys[group] = append(surveys[group], make([]rune, 0))

		for _, r := range []rune(s.Text()) {
			surveys[group][person] = append(surveys[group][person], r)
		}

		person++
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
