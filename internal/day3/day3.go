package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isEnabled(dos, donts [][]int, n int) bool {
	do_i := 0
	dont_i := 0
	for i := 0; i < len(dos) && dos[i][0] < n; i++ {
		do_i = dos[i][0]
	}

	for i := 0; i < len(donts) && donts[i][0] < n; i++ {
		dont_i = donts[i][0]
	}

	return (do_i > dont_i) || (dont_i == 0)
}

func Run() {
	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	s, err := os.ReadFile("../internal/day3/input")
	if err != nil {
		panic(err)
	}
	//s := []byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)
	matches := re.FindAllSubmatch(s, -1)
	matches_i := re.FindAllSubmatchIndex(s, -1)

	sum := 0
	for _, m := range matches {
		x, _ := strconv.Atoi(string(m[1]))
		y, _ := strconv.Atoi(string(m[2]))
		sum += x * y
	}
	fmt.Printf("%d\n", sum)

	dos := regexp.MustCompile(`do\(\)`)
	donts := regexp.MustCompile(`don't\(\)`)

	dos_i := dos.FindAllSubmatchIndex(s, -1)
	donts_i := donts.FindAllSubmatchIndex(s, -1)

	ns := 0
	for i, m := range matches {
		mi := matches_i[i][0]
		if isEnabled(dos_i, donts_i, mi) {
			x, _ := strconv.Atoi(string(m[1]))
			y, _ := strconv.Atoi(string(m[2]))
			ns += x * y
		}

	}
	fmt.Printf("%d\n", ns)
	//fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`), -1))
	//[["mul(2,4)" "2" "4"] ["mul(5,5)" "5" "5"] ["mul(11,8)" "11" "8"] ["mul(8,5)" "8" "5"]]
	//fmt.Println(re.FindAllSubmatchIndex([]byte(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`), -1))
	//fmt.Println(re.FindAllSubmatchIndex([]byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`), -1))
	//fmt.Println(dos.FindAllSubmatchIndex([]byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`), -1))
	//fmt.Println(donts.FindAllSubmatchIndex([]byte(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`), -1))
	//[[1 9 5 6 7 8] [28 36 32 33 34 35] [48 57 52 54 55 56] [64 72 68 69 70 71]]
	//[[59 63]]
	//[[20 27]]
}
