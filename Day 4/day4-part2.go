package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := LoadPassportsFromFile()
	validPassports := 0

	validEyeColors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	validHairColors := regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)
	validHeight := regexp.MustCompile(`^([\d]+)(cm|in)$`)

	// Validate rules for passport fields
	for _, passport := range passports {
		byr, vByr := passport["byr"]
		iyr, vIyr := passport["iyr"]
		eyr, vEyr := passport["eyr"]
		hgt, vHgt := passport["hgt"]
		hcl, vHcl := passport["hcl"]
		ecl, vEcl := passport["ecl"]
		pid, vPid := passport["pid"]

		// Verify that it contains all eight fields
		if !(vByr && vIyr && vEyr && vHgt && vHcl && vEcl && vPid) {
			log.Printf("Invalid passport (Missing fields): %v", passport)
			continue
		}

		// (Birth Year) - four digits; at least 1920 and at most 2002.
		byrI, err := strconv.Atoi(byr)
		if err != nil || byrI < 1920 || byrI > 2002 {
			log.Printf("Invalid passport (byr): %v", passport)
			continue
		}

		// (Issue Year) - four digits; at least 2010 and at most 2020.
		iyrI, err := strconv.Atoi(iyr)
		if err != nil || iyrI < 2010 || iyrI > 2020 {
			log.Printf("Invalid passport (iyr): %v", passport)
			continue
		}

		// (Expiration Year) - four digits; at least 2020 and at most 2030.
		eyrI, err := strconv.Atoi(eyr)
		if err != nil || eyrI < 2020 || eyrI > 2030 {
			log.Printf("Invalid passport (eyr): %v", passport)
			continue
		}

		// (Passport ID) - a nine-digit number, including leading zeroes.
		_, err = strconv.Atoi(pid)
		if err != nil || len(pid) != 9 {
			log.Printf("Invalid passport (pid): %v", passport)
			continue
		}

		// (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		validEye := false
		for _, c := range validEyeColors {
			if ecl == c {
				validEye = true
				continue
			}
		}

		if !validEye {
			log.Printf("Invalid passport (pid): %v", passport)
			continue
		}

		// (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		matchedHair := validHairColors.Match([]byte(hcl))
		if !matchedHair {
			log.Printf("Invalid passport (hcl): %v", passport)
			continue
		}

		// (Height) - a number followed by either cm or in:
		matchedHeight := validHeight.Match([]byte(hgt))
		if !matchedHeight {
			log.Printf("Invalid passport (hgt 1): %v", passport)
			continue
		}

		h := validHeight.FindStringSubmatch(hgt)
		if h[2] == "in" {
			height, _ := strconv.Atoi(h[1])
			if height < 59 || height > 76 {
				log.Printf("Invalid passport (hgt 2): %v", passport)
				continue
			}
		} else { // cm
			height, _ := strconv.Atoi(h[1])
			if height < 150 || height > 193 {
				log.Printf("Invalid passport (hgt 3) (%v): %v", h[2], passport)
				continue
			}
		}

		validPassports++
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
