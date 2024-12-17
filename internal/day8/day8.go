package day8

import (
	"aoc/internal/util"
	"bufio"
	"os"
)

func isAntenna(b byte) bool {
	if ((b >= 48) && (b <= 57)) ||
		((b >= 97) && (b <= 122)) ||
		((b >= 65) && (b <= 90)) {
		return true
	}
	return false
}

func showAntenna(p []string) {

	for _, s := range p {
		for _, b := range s {
			if isAntenna(byte(b)) {
				print("A,")
			} else {
				print(".,")
			}
		}
		println()
	}
}

type Point struct {
	x, y int
}

func inRange(N, r, c int) bool {
	if (r >= 0) && (r < N) && (c >= 0) && (c < N) {
		return true
	}
	return false
}

func isAntennaRange(puzzle []string, N, r, c int) bool {
	return inRange(N, r, c) && isAntenna(puzzle[r][c])
}

func antennaMap(p []string) (int, map[rune][]Point) {
	m := make(map[rune][]Point)
	for x, l := range p {
		for y, v := range l {
			if v != '.' {
				_, ok := m[v]
				if !ok {
					m[v] = []Point{}
				}
				m[v] = append(m[v], Point{x, y})
			}
		}
	}
	return len(p), m
}

func antiNodes(size int, m map[rune][]Point) map[rune]*util.Set[Point] {
	n := make(map[rune]*util.Set[Point])
	for k, v := range m {
		n[k] = util.NewSet[Point]()
		for _, p := range v {
			for _, p2 := range v {
				if p != p2 {
					deltaX := p.x - p2.x
					deltaY := p.y - p2.y

					pt := Point{p.x + deltaX, p.y + deltaY}
					if inRange(size, pt.x, pt.y) {
						n[k].Add(pt)
					}

					pt = Point{p2.x - deltaX, p2.y - deltaY}
					if inRange(size, pt.x, pt.y) {
						n[k].Add(pt)
					}
				}
			}
		}
	}
	return n
}

func antiNodes2(size int, m map[rune][]Point) map[rune]*util.Set[Point] {
	n := make(map[rune]*util.Set[Point])
	for k, v := range m {
		n[k] = util.NewSet[Point]()
		for _, p := range v {
			for _, p2 := range v {
				if p != p2 {

					for d := -size; d < size+1; d++ {
						deltaX := d * (p.x - p2.x)
						deltaY := d * (p.y - p2.y)

						pt := Point{p.x + deltaX, p.y + deltaY}
						if inRange(size, pt.x, pt.y) {
							n[k].Add(pt)
						}

						pt = Point{p2.x - deltaX, p2.y - deltaY}
						if inRange(size, pt.x, pt.y) {
							n[k].Add(pt)
						}
					}
				}
			}
		}
	}
	return n
}

func uniq(m map[rune]*util.Set[Point]) int {
	s := util.NewSet[Point]()
	for _, nodes := range m {
		for _, p := range nodes.Elements() {
			s.Add(p)
		}
	}
	return len(s.Elements())
}

func Run() {
	file, _ := os.Open("../internal/day8/input")
	defer file.Close()

	reader := bufio.NewReader(file)
	puzzle := []string{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		puzzle = append(puzzle, string(line))
	}

	//N := len(puzzle)
	//println(N)
	//showAntenna(puzzle)
	N, m := antennaMap(puzzle)
	aNodes := antiNodes(N, m)
	println(uniq(aNodes))

	aNodes2 := antiNodes2(N, m)
	println(uniq(aNodes2))

	//r := antiNodes(puzzle, N, 2, 5)
	//println(r)
	//println(allAntiNodes(puzzle))

}
