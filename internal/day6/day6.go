package day6

import (
	"aoc/internal/util"
	"bufio"
	"os"
	"strings"
)

func inArea(r, c, N int) bool {
	return (r >= 0) && (r < N) && (c >= 0) && (c < N)
}

func rotate90(d util.Direction) util.Direction {
	switch d {
	case util.N:
		return util.E
	case util.E:
		return util.S
	case util.S:
		return util.W
	case util.W:
		return util.N
	}
	return util.N //shouldn;t hit
}

func move(area []string, size, r, c int, d util.Direction) (int, int, util.Direction) {
	nr, nc := 0, 0
	switch d {
	case util.N:
		nr, nc = r-1, c
	case util.E:
		nr, nc = r, c+1
	case util.S:
		nr, nc = r+1, c
	case util.W:
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
// ri,ci = starting location, di = starting util.Direction
func patrol1(area []string, ri, ci int, di util.Direction) (int, []Point) {
	s := len(area)
	d := di
	visited := map[Point]bool{{ri, ci}: true}
	for r, c := ri, ci; inArea(r, c, s); r, c, d = move(area, s, r, c, d) {
		visited[Point{r, c}] = true
	}
	slc := make([]Point, len(visited))
	i := 0
	for k, _ := range visited {
		slc[i] = k
		i++
	}
	return len(visited), slc
}

type PointDir struct {
	x, y int
	d    util.Direction
}

func isLoopPath(area []string, ri, ci int, di util.Direction) bool {
	//we are in a loop if we come back to the same point
	//and traveling in same util.Direction
	s := len(area)
	d := di
	visited := map[PointDir]bool{}
	for r, c := ri, ci; inArea(r, c, s); r, c, d = move(area, s, r, c, d) {
		p := PointDir{r, c, d}
		_, ok := visited[p]
		if ok {
			return true
		}
		visited[PointDir{r, c, d}] = true
	}
	return false
}

func addObstacle(area []string, r, c int) {
	u := []rune(area[r])
	u[c] = '#'
	area[r] = string(u)
}

func removeObstacle(area []string, r, c int) {
	u := []rune(area[r])
	u[c] = '.'
	area[r] = string(u)
}

func patrol2(area []string, path []Point, ri, ci int, di util.Direction) int {
	sum := 0
	for _, p := range path {
		if (p.x == ri) && (p.y == ci) {
			continue
		}
		addObstacle(area, p.x, p.y)
		if isLoopPath(area, ri, ci, di) {
			sum++
		}
		removeObstacle(area, p.x, p.y)
	}
	return sum
}

func Run() {

	file, _ := os.Open("../internal/day6/input")
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
	cnt, path := patrol1(area, r, c, util.N)
	println(cnt)
	println(patrol2(area, path, r, c, util.N))

}
