package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rules map[int][]int

func isBefore(rules Rules, a, b int) bool {
	aRules := rules[a]
	for _, v := range aRules {
		if v == b {
			return true
		}
	}
	return false
}

func isValid(rules Rules, update []int) (int, bool) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if !isBefore(rules, update[i], update[j]) {
				return 0, false
			}
		}
	}
	return update[len(update)/2], true
}

func part1(rules Rules, updates [][]int) int {
	sum := 0
	for _, u := range updates {
		mid, valid := isValid(rules, u)
		if valid {
			sum += mid
		}
	}
	return sum
}

func fix(aRules [][]int, update []int) int {
	cnts := make([]int, len(update))
	for _, v := range update {
		//get the column counts for rows in update
		s := 0
		for _, r := range update {
			s += aRules[r][v]
		}
		cnts[s] = v
	}
	return cnts[len(cnts)/2]
}

func isValid2(rules Rules, aRules [][]int, update []int) int {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if !isBefore(rules, update[i], update[j]) {
				return fix(aRules, update)
			}
		}
	}
	return 0
}

func part2(rules Rules, arules [][]int, updates [][]int) int {
	sum := 0
	for _, u := range updates {
		mid := isValid2(rules, arules, u)
		sum += mid
	}
	return sum
}

func rulesToArray(rules Rules, max int) [][]int {
	ret := make([][]int, max+1)
	for i := range ret {
		ret[i] = make([]int, max+1)
	}
	for k, vs := range rules {
		for _, v := range vs {
			ret[k][v] = 1
		}
	}
	return ret
}

func day5() {

	file, _ := os.Open("day5/input")
	defer file.Close()

	reader := bufio.NewReader(file)

	rules := map[int][]int{}
	updates := [][]int{}

	max := -1
	min := 1000000
	rulesMode := true
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if rulesMode {
			s := strings.Split(string(line), "|")
			if len(s) > 1 {
				k, _ := strconv.Atoi(s[0])
				v, _ := strconv.Atoi(s[1])
				_, ok := rules[k]
				if !ok {
					rules[k] = []int{v}
				} else {
					rules[k] = append(rules[k], v)
				}
				if k < min {
					min = k
				}
				if k > max {
					max = k
				}
				if v < min {
					min = v
				}
				if v > max {
					max = v
				}
			} else {
				rulesMode = false
			}
		} else {
			//reading updates
			pages := []int{}
			s := strings.Split(string(line), ",")
			for _, v := range s {
				p, _ := strconv.Atoi(v)
				pages = append(pages, p)
			}
			updates = append(updates, pages)
		}
	}
	fmt.Printf("min=%d max=%d\n", min, max)
	aRules := rulesToArray(rules, max)

	//fmt.Printf("r[98][43]=%d\n", aRules[98][43])
	//fmt.Printf("r[98][21]=%d\n", aRules[98][21])

	m, e := isValid(rules, []int{75, 47, 61, 53, 29})
	fmt.Printf("%v %v\n", m, e)
	fmt.Printf("part1: %d\n", part1(rules, updates))
	fmt.Printf("part2: %d\n", part2(rules, aRules, updates))
}
