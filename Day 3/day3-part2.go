package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	sledMap := LoadMapFromFile()

	var trial1, trial2, trial3, trial4, trial5 int

	for i := 0; i < sledMap.Height-1; i++ {
		trial1 += sledMap.MoveCursor(1, 1)
	}
	sledMap.ResetCursor()
	log.Printf("Trial 1: %v", trial1)

	for i := 0; i < sledMap.Height-1; i++ {
		trial2 += sledMap.MoveCursor(3, 1)
	}
	sledMap.ResetCursor()
	log.Printf("Trial 2: %v", trial2)

	for i := 0; i < sledMap.Height-1; i++ {
		trial3 += sledMap.MoveCursor(5, 1)
	}
	sledMap.ResetCursor()
	log.Printf("Trial 3: %v", trial3)

	for i := 0; i < sledMap.Height-1; i++ {
		trial4 += sledMap.MoveCursor(7, 1)
	}
	sledMap.ResetCursor()
	log.Printf("Trial 4: %v", trial4)

	for i := 0; i < sledMap.Height-1; i += 2 {
		trial5 += sledMap.MoveCursor(1, 2)
	}
	sledMap.ResetCursor()
	log.Printf("Trial 5: %v", trial5)

	log.Printf("%v * %v * %v * %v * %v = %v", trial1, trial2, trial3, trial4, trial5, trial1*trial2*trial3*trial4*trial5)

	//log.Printf("%v trees hit. Current position: [%v, %v]", treesHit, sledMap.Cursor.X, sledMap.Cursor.Y)
}

type Position struct {
	X int
	Y int
}

type Map struct {
	Width         int
	Height        int
	MapData       []string
	RepeatingData []string
	Cursor        Position
}

// Moves the cursor and returns how many trees are hit
func (mapObj *Map) MoveCursor(moveX int, moveY int) int {
	treeHit := 0
	oldCursor := mapObj.Cursor
	failedMoveX, failedMoveY := -999, -999

	if moveX == 0 && moveY == 0 {
		return treeHit
	}

	// Check x-axis

	if mapObj.Cursor.X+moveX >= mapObj.Width {
		mapObj.ExtendMapRight()
	}

	if mapObj.Cursor.X+moveX > 0 { // Check for moving left outside of bounds
		mapObj.Cursor.X += moveX
	} else {
		failedMoveX = mapObj.Cursor.X
	}

	// Check y-axis

	if mapObj.Cursor.Y+moveY >= mapObj.Height {
		mapObj.Cursor.Y = mapObj.Height - 1
	} else if mapObj.Cursor.Y+moveY < 0 {
		failedMoveY = mapObj.Cursor.Y
	} else {
		mapObj.Cursor.Y += moveY
	}

	// Check for bad moves

	if failedMoveX != -999 && failedMoveY != -999 {
		log.Printf("Invalid Move on X+Y: [%v, %v] from [%v, %v]. Resetting position.", moveX, moveY, failedMoveX, failedMoveY)
		mapObj.Cursor.X = oldCursor.X
		mapObj.Cursor.Y = oldCursor.Y
		return treeHit
	} else if failedMoveX != -999 {
		log.Printf("Invalid Move on X: [%v, %v] from [%v, %v]. Resetting position.", moveX, moveY, failedMoveX, oldCursor.Y)
		mapObj.Cursor.X = failedMoveX
		mapObj.Cursor.Y = oldCursor.Y
		return treeHit
	} else if failedMoveY != -999 {
		log.Printf("Invalid Move on Y: [%v, %v] from [%v, %v]. Resetting position.", moveX, moveY, oldCursor.X, failedMoveY)
		mapObj.Cursor.X = oldCursor.X
		mapObj.Cursor.Y = failedMoveY
		return treeHit
	}

	// Check for hit tree

	if []rune(mapObj.MapData[mapObj.Cursor.Y])[mapObj.Cursor.X] == '#' {
		treeHit++
	}

	return treeHit
}

func (mapObj *Map) ExtendMapRight() {
	mapObj.Width += len(mapObj.RepeatingData[0])

	newMapData := mapObj.MapData
	for i, line := range mapObj.MapData {
		newLine := line + mapObj.RepeatingData[i]
		newMapData[i] = newMapData[i] + newLine
	}
	mapObj.MapData = newMapData
}

func (mapObj *Map) ResetCursor() {
	mapObj.Cursor.X = 0
	mapObj.Cursor.Y = 0
}

func LoadMapFromFile() Map {
	fptr := flag.String("fpath", "map", "./")
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

	width := 0
	height := 0

	var mapData []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		if width == 0 {
			width = len(s.Text())
		}

		var line string
		for _, c := range s.Text() {
			line = line + string(c)
		}
		mapData = append(mapData, line)
		height++
	}

	if s.Err() != nil {
		log.Fatal(err)
	}

	mapObj := Map{
		Width:         width,
		Height:        height,
		MapData:       mapData,
		RepeatingData: mapData,
	}

	return mapObj

	//log.Printf("%v | %v | %v", mapObj.Width, mapObj.Height, mapObj.MapData)
}
