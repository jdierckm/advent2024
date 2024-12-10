package main

import (
	"bufio"
	"os"
	"strings"
)

func inArea(r, c, N int) bool {
	return (r >= 0) && (r < N) && (c >= 0) && (c < N)
}

func rotate90(d direction) direction {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return N //shouldn;t hit
}

func move(area []string, size, r, c int, d direction) (int, int, direction) {
	nr, nc := 0, 0
	switch d {
	case N:
		nr, nc = r-1, c
	case E:
		nr, nc = r, c+1
	case S:
		nr, nc = r+1, c
	case W:
		nr, nc = r, c-1
	}

	if inArea(nr, nc, size) {
		if area[nr][nc] != '#' {
			return nr, nc, d
		} else {
			return r, c, rotate90(d)
		}
	}
	return nr, nc, d
}

type Point struct {
	x, y int
}

// return the total number of spaces walked before leaving area
// ri,ci = starting location, di = starting direction
func patrol1(area []string, ri, ci int, di direction) int {
	s := len(area)
	d := di
	visited := map[Point]bool{Point{ri, ci}: true}
	for r, c := ri, ci; inArea(r, c, s); r, c, d = move(area, s, r, c, d) {
		visited[Point{r, c}] = true
	}
	return len(visited)
}

func day6() {

	file, _ := os.Open("day6/input")
	defer file.Close()

	reader := bufio.NewReader(file)

	area := []string{}
	i, r, c := 0, 0, 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		area = append(area, string(line))
		if c <= 0 {
			c = strings.Index(string(line), "^")
			if c >= 0 {
				r = i
			}
		}
		i++
	}
	println(r, ":", c)
	println(patrol1(area, r, c, N))

}
