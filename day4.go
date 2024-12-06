package main

import (
	"bufio"
	"os"
)

type direction int

const (
	N direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func getChar(p []string, row, col int, dir direction, offset int) byte {
	var r, c int
	switch dir {
	case N:
		r = row - offset
		c = col
	case NE:
		r = row - offset
		c = col + offset
	case E:
		r = row
		c = col + offset
	case SE:
		r = row + offset
		c = col + offset
	case S:
		r = row + offset
		c = col
	case SW:
		r = row + offset
		c = col - offset
	case W:
		r = row
		c = col - offset
	case NW:
		r = row - offset
		c = col - offset
	}

	if (r < 0) || (r >= len(p)) || (c < 0) || (c >= len(p)) {
		return ' '
	}
	return p[r][c]
}

func getTranslations(r, c, offset int, d direction) (r1, c1 int, d1 direction, r2, c2 int, d2 direction) {

	switch d {
	case NE:
		r1, c1, d1, r2, c2, d2 = r-offset, c, SE, r, c+offset, NW
	case SE:
		r1, c1, d1, r2, c2, d2 = r+offset, c, NE, r, c+offset, SW
	case SW:
		r1, c1, d1, r2, c2, d2 = r, c-offset, SE, r+offset, c, NW
	case NW:
		r1, c1, d1, r2, c2, d2 = r, c-offset, NE, r-offset, c, SW
	}
	return
}

func checkWord(puzzle []string, word []byte, r, c int, d direction) bool {
	cnt := 0
	for i, letter := range word {
		if letter == getChar(puzzle, r, c, d, i) {
			cnt++
		}
	}
	if cnt == len(word) {
		return true
	}
	return false
}

func day4() {

	file, _ := os.Open("day4/input")
	defer file.Close()

	reader := bufio.NewReader(file)
	word := []byte{'X', 'M', 'A', 'S'}

	puzzle := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		puzzle = append(puzzle, string(line))
	}

	directions := []direction{N, NE, E, SE, S, SW, W, NW}
	cnt := 0
	for r, v := range puzzle {
		for c, _ := range v {
			for _, d := range directions {
				if checkWord(puzzle, word, r, c, d) {
					cnt++
				}
			}
		}
	}

	println(cnt)

	//part two
	directions = []direction{NE, SE, SW, NW}
	word = []byte{'M', 'A', 'S'}

	//fmt.Printf("%c\n", getChar([]string{"123", "456", "789"}, -1, -1, N, 0))
	//fmt.Printf("%c\n", getChar([]string{"123", "456", "789"}, -1, -1, N, 5))
	//fmt.Printf("%c\n", getChar([]string{"123", "456", "789"}, 0, 0, N, 0))

	cnt = 0
	for r, v := range puzzle {
		for c, _ := range v {
			for _, d := range directions {

				if checkWord(puzzle, word, r, c, d) {
					//we found one diagonal, now check the other side of the X
					r1, c1, d1, r2, c2, d2 := getTranslations(r, c, len(word)-1, d)
					if checkWord(puzzle, word, r1, c1, d1) {
						cnt++
					}
					if checkWord(puzzle, word, r2, c2, d2) {
						cnt++
					}
				}

			}
		}
	}
	print(cnt / 2)
}
