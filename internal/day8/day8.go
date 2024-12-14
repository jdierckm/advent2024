package day8

import (
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

func Run() {
	file, _ := os.Open("../internal/day8/test")
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

	N := len(puzzle)
	println(N)
	showAntenna(puzzle)

}
