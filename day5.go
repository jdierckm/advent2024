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

func day5() {

	file, _ := os.Open("day5/input")
	defer file.Close()

	reader := bufio.NewReader(file)

	rules := map[int][]int{}
	updates := [][]int{}

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
	m, e := isValid(rules, []int{75, 47, 61, 53, 29})
	fmt.Printf("%v %v\n", m, e)
	fmt.Printf("part1: %d\n", part1(rules, updates))
}
