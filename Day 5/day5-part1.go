package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	seats := LoadInputFromFile()

	// Sort descending
	sort.Slice(seats, func(i, j int) bool {
		if len(seats[i]) == 0 && len(seats[j]) == 0 {
			return false
		}

		if len(seats[i]) == 0 || len(seats[j]) == 0 {
			return len(seats[i]) == 0
		}

		return seats[i][0] > seats[j][0]
	})

	// Get highest value
	log.Printf("Highest seat ID: %v", seats[0][0])
}

func LoadInputFromFile() [][]int64 {
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

	var seats [][]int64

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

		seat := make([]int64, 3)
		seat[0] = seatId
		seat[1] = colBinaryInt
		seat[2] = rowBinaryInt

		seats = append(seats, seat)

	}

	return seats
}
