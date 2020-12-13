package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	passports := LoadPassportsFromFile()
	validPassports := 0

	// Verify that it contains all eight fields
	for _, passport := range passports {
		_, byr := passport["byr"]
		_, iyr := passport["iyr"]
		_, eyr := passport["eyr"]
		_, hgt := passport["hgt"]
		_, hcl := passport["hcl"]
		_, ecl := passport["ecl"]
		_, pid := passport["pid"]

		if byr && iyr && eyr && hgt && hcl && ecl && pid {
			validPassports++
		} else {
			log.Printf("Invalid passport: %v", passport)
		}
	}

	log.Printf("Valid passports: %v", validPassports)
}

func LoadPassportsFromFile() []map[string]string {
	fptr := flag.String("fpath", "passports", "./")
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

	var passports []map[string]string
	var passport map[string]string = nil

	s := bufio.NewScanner(f)
	for s.Scan() {
		if strings.TrimSpace(s.Text()) == "" {
			// Empty line; Close passport if open, and proceed to next line
			if passport != nil {
				passports = append(passports, passport)
				passport = nil
			}
			continue
		}

		if passport == nil {
			// Open new passport
			passport = make(map[string]string)
		}

		keyValuePairs := strings.Split(s.Text(), " ")
		for _, kv := range keyValuePairs {
			kvSplit := strings.Split(kv, ":")

			val, dup := passport[kvSplit[0]]
			if dup {
				log.Printf("Key already exists in passport [%v:%v]; Updating to latest value [%v:%v]", kvSplit[0], val, kvSplit[0], kvSplit[1])
			}

			passport[kvSplit[0]] = kvSplit[1]
		}

	}

	if passport != nil {
		// EOF; Close passport if open
		passports = append(passports, passport)
		passport = nil
	}

	log.Printf("Read %v passports", len(passports))

	return passports
}
