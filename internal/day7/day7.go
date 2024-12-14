package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Total  int
	Inputs []int
}

func parse(line []byte) Problem {
	l := string(line)
	s := strings.Split(l, ": ")
	operands := strings.Split(s[1], " ")
	retOps := make([]int, len(operands))
	for i, v := range operands {
		x, _ := strconv.Atoi(v)
		retOps[i] = x
	}
	r, _ := strconv.Atoi(s[0])
	return Problem{r, retOps}
}

type operator struct {
	n  string
	op func(int, int) int
}

func plus() operator {
	return operator{n: "+", op: func(a, b int) int { return a + b }}
}

func mult() operator {
	return operator{n: "*", op: func(a, b int) int { return a * b }}
}

func concat() operator {
	return operator{n: "||", op: func(a, b int) int {
		as := strconv.Itoa(a)
		bs := strconv.Itoa(b)
		n, _ := strconv.Atoi(as + bs)
		return n
	},
	}
}

var operators = []operator{
	operator{n: "+", op: func(a, b int) int { return a + b }},
	operator{n: "*", op: func(a, b int) int { return a * b }}}

func initOps(opSet []operator) [][]operator {
	ret := [][]operator{}
	for _, o := range opSet {
		ret = append(ret, []operator{o})
	}
	return ret
}

func addAnotherOrig(ops [][]operator, opSet []operator) [][]operator {
	ret := [][]operator{}
	for _, v := range ops {
		for _, o := range opSet {
			t := append(v, o)
			ret = append(ret, t)
		}
	}
	return ret
}

func addAnother(ops [][]operator, opSet []operator) [][]operator {
	l := len(opSet)
	ret := make([][]operator, l*len(ops))
	for i, v := range ops {
		for j, o := range opSet {
			t := make([]operator, len(v)+1)
			copy(t, v)
			t[len(v)] = o
			ret[l*i+j] = t
		}
	}
	return ret
}

func check(p Problem) bool {

	ops := initOps(operators)

	for i := 1; i < len(p.Inputs)-1; i++ {
		ops = addAnother(ops, operators)
	}

	for _, op := range ops {
		s := p.Inputs[0]
		//print(i, " : ", s)
		for n := 1; n < len(p.Inputs) && s <= p.Total; n++ {
			//print(op[n-1].n, p.Inputs[n])
			s = op[n-1].op(s, p.Inputs[n])
		}
		//println("=", s)
		if s == p.Total {
			return true
		}
	}

	return false
}

func checkAll(problems []Problem) int {
	s := 0
	for _, p := range problems {
		if check(p) {
			s += p.Total
		}
	}
	return s
}

func Run() {
	file, _ := os.Open("../internal/day7/input")
	defer file.Close()

	reader := bufio.NewReader(file)

	input := []Problem{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		input = append(input, parse(line))
	}

	//println(check(input[784]))
	println(checkAll(input))

	operators = append(operators, concat())
	//println(operators[2].op(12, 34))
	println(checkAll(input))
}
