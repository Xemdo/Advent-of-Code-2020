package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	seats := LoadInputFromFile()

	log.Printf("Values: %v", seats)

	// Get highest value
	maxSeat := max(seats)
	log.Printf("Highest seat ID: %v", maxSeat)

	yourSeat := int64(-1)

	for a := int64(1); a < maxSeat-1; a++ {
		if _, ok := seats[a]; ok {
			// Value exists; Not ours
			continue
		} else {
			// Value doesn't exist. Check plus and minus one for existing.
			_, ok1 := seats[a-1]
			_, ok2 := seats[a+1]

			if ok1 && ok2 {
				yourSeat = a
				break
			}
		}
	}

	log.Printf("Your seat: %v", yourSeat)
}

func max(values map[int64][]int64) int64 {
	var maxNumber int64
	for maxNumber = range values {
		break
	}

	for n := range values {
		if n > maxNumber {
			maxNumber = n
		}
	}

	return maxNumber
}

func LoadInputFromFile() map[int64][]int64 {
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

	seats := make(map[int64][]int64)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		rowBinaryStr := string([]rune(line)[:7])
		rowBinaryStr = strings.ReplaceAll(rowBinaryStr, "F", "0")
		rowBinaryStr = strings.ReplaceAll(rowBinaryStr, "B", "1")
		rowBinaryInt, err := strconv.ParseInt(rowBinaryStr, 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		colBinaryStr := string([]rune(line)[7:10])
		colBinaryStr = strings.ReplaceAll(colBinaryStr, "L", "0")
		colBinaryStr = strings.ReplaceAll(colBinaryStr, "R", "1")
		colBinaryInt, err := strconv.ParseInt(colBinaryStr, 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		seatId := rowBinaryInt*8 + colBinaryInt

		seat := make([]int64, 2)
		seat[0] = colBinaryInt
		seat[1] = rowBinaryInt

		seats[seatId] = seat
	}

	return seats
}
