package day11

import (
	"bufio"
	"container/list"
	"math"
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

func toList(stones []int) *list.List {
	l := list.New()
	for _, v := range stones {
		l.PushBack(v)
	}
	return l
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

func split2(n int) (bool, int, int) {
	l := 1
	d := 1
	for {
		r := n / d
		if r < 10 {
			break
		}
		d = d * 10
		l++
	}
	if (l % 2) == 1 {
		return false, 0, 0
	}
	p := int(math.Pow10(l / 2))
	n1 := n / p
	n2 := n - n1*p
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

func blink2(stones *list.List) {
	for se := stones.Front(); se != nil; se = se.Next() {
		s := se.Value.(int)
		if s == 0 {
			se.Value = 1
		} else {
			if isEven, n1, n2 := split(s); isEven {
				stones.InsertBefore(n1, se)
				h := stones.InsertBefore(n2, se)
				stones.Remove(se)
				se = h
				//stones = slices.Replace(stones, i, i+1, n1, n2)
			} else {
				se.Value = s * 2024
			}
		}

	}
}

func printStones(i int, s *list.List) {
	print(i, ": ")
	for s := s.Front(); s != nil; s = s.Next() {
		print(s.Value.(int), " ")
	}
	println()
}

func Run() {
	b, n1, n2 := split2(253000)
	println(b, n1, n2)

	file, _ := os.Open("../internal/day11/test")
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	stones := parse(line)
	sl := toList(stones)
	printStones(0, sl)

	for i := 0; i < 25; i++ {
		blink2(sl)
		//printStones(i+1, sl)
		println(i, ": ", sl.Len())
	}
	println(sl.Len())

}
