package day11

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(l string) []int {
	ret := []int{}

	split := strings.Split(l, " ")
	for _, s := range split {
		v, _ := strconv.Atoi(s)
		ret = append(ret, v)

	}
	return ret
}

func digits(n int) int {
	return len(strconv.Itoa(n))
}

func split(n int) (bool, int, int) {
	s := strconv.Itoa(n)
	l := len(s)
	if (l % 2) == 1 {
		return false, 0, 0
	}
	n1, _ := strconv.Atoi(s[:l/2])
	n2, _ := strconv.Atoi(s[l/2:])
	return true, n1, n2
}

func blink(stones []int) []int {
	l := len(stones)
	for i := l - 1; i >= 0; i-- {
		s := stones[i]
		if s == 0 {
			stones[i] = 1
		} else {
			if isEven, n1, n2 := split(s); isEven {
				stones = slices.Replace(stones, i, i+1, n1, n2)
			} else {
				stones[i] = s * 2024
			}
		}

	}
	return stones
}

func Run() {
	file, _ := os.Open("../internal/day11/input")
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	stones := parse(line)
	fmt.Printf("%v\n", stones)
	for i := 0; i < 75; i++ {
		stones = blink(stones)
		//fmt.Printf("%v\n", stones)
		println(i, ": ", len(stones))
	}

}
