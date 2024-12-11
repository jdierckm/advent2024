package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// returns a map of the number of times each number appears in the sorted slice of ints
func rcounts(sorted []int) map[int]int {
	cnts := make(map[int]int)
	p := -1
	pi := 0
	l := len(sorted)
	for i, v := range sorted {
		if p != v {
			cnts[p] = i - pi
			p = v
			pi = i
		}
		if i == l-1 {
			cnts[v] = i - pi + 1
		}
	}
	return cnts
}

func Run() {
	file, _ := os.Open("../internal/day1/day1-input")
	defer file.Close()

	reader := bufio.NewReader(file)

	l := []int{}
	r := []int{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			//fmt.Printf("ReadLine: %q\n", line)
			str := strings.Replace(string(line), " ", "", 2)
			split := strings.Split(str, " ")
			li, _ := strconv.Atoi(split[0])
			l = append(l, li)
			ri, _ := strconv.Atoi(split[1])
			r = append(r, ri)
		}
	}
	slices.Sort(l)
	slices.Sort(r)
	s := 0.0
	for i, _ := range l {
		s += math.Abs(float64(l[i] - r[i]))
	}
	fmt.Printf("%f\n", s)

	rcounts([]int{123, 123, 125, 128})
	rcounts([]int{123, 123, 125, 128, 128})

	rcnts := rcounts(r)
	sim := 0
	for _, v := range l {
		sim += v * rcnts[v]
	}
	fmt.Printf("sim=%d\n", sim)
}
