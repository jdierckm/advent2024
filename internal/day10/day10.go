package day10

import (
	"aoc/internal/util"
	"bufio"
	"os"
	"strconv"
)

func parse(line string) []int {
	ret := []int{}
	for _, v := range line {
		i, _ := strconv.Atoi(string(v))
		ret = append(ret, i)
	}
	return ret
}

type Point struct {
	r, c int
}

func inRange(N, r, c int) bool {
	if (r >= 0) && (r < N) && (c >= 0) && (c < N) {
		return true
	}
	return false
}

func disc(topo [][]int, start Point) []Point {

	ret := []Point{}
	N := len(topo)
	v := topo[start.r][start.c]
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == j) || (i == -j) {
				continue
			}
			r, c := start.r+i, start.c+j
			if !inRange(N, start.r+i, start.c+j) {
				continue
			}
			if topo[r][c] == v+1 {
				if topo[r][c] == 9 {
					ret = append(ret, Point{r, c})
				} else {
					d := disc(topo, Point{r, c})
					ret = append(ret, d...)
				}
			}
		}
	}
	return ret
}

func discAll(topo [][]int) (int, int) {
	sum := 0
	sum2 := 0
	for r, v := range topo {
		for c, _ := range v {
			if topo[r][c] == 0 {
				pts := disc(topo, Point{r, c})
				sum2 += len(pts)
				s := util.NewSet[Point]()
				for _, p := range pts {
					s.Add(p)
				}
				sum += len(s.Elements())
			}
		}
	}
	return sum, sum2
}

func Run() {
	file, _ := os.Open("../internal/day10/input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	topo := [][]int{}

	for scanner.Scan() {
		topo = append(topo, parse(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	//fmt.Printf("%v\n", topo)
	println(discAll(topo))

}
