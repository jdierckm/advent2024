package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(input string) bool {
	str := string(input)
	split := strings.Split(str, " ")
	p := 0
	isInc := false
	for i, vs := range split {
		v, _ := strconv.Atoi(vs)
		if i == 1 {
			isInc = v > p
		}
		if i > 0 {
			if isInc {
				if (v <= p) || (v > p+3) {
					return false
				}
			}
			if !isInc {
				if (v >= p) || (v < p-3) {
					return false
				}
			}
		}
		p = v
	}
	return true
}

func anySafe(input string) bool {
	for i, _ := range input {
		if isSafe2(input, i) {
			return true
		}
	}
	return false
}

func isSafe2(input string, drop int) bool {
	str := string(input)
	split := strings.Split(str, " ")
	p := 0
	isInc := false
	i := -1
	for j, vs := range split {
		v, _ := strconv.Atoi(vs)

		if j != drop {
			i++
		} else {
			continue
		}

		if i == 1 {
			isInc = v > p
		}
		if i > 0 {
			if isInc {
				if (v <= p) || (v > p+3) {
					return false
				}
			}
			if !isInc {
				if (v >= p) || (v < p-3) {
					return false
				}
			}
		}
		p = v
	}
	return true
}

func day2() {
	//fmt.Printf("1 2 7 8 9 - %v\n", anySafe("1 2 7 8 9"))
	//fmt.Printf("1 3 2 4 5 - %v\n", anySafe("1 3 2 4 5"))
	//fmt.Printf("10 1 2 4 5 - %v\n", anySafe("10 1 2 4 5"))

	file, _ := os.Open("day2/input")
	defer file.Close()

	reader := bufio.NewReader(file)

	cnt := 0
	cnt2 := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		if isSafe(string(line)) {
			cnt += 1
		} else {
			if anySafe(string(line)) {
				cnt2 += 1
			}
		}
	}
	fmt.Printf("%d\n", cnt)
	fmt.Printf("%d\n", cnt2)

}
