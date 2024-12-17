package day9

import (
	"bufio"
	"os"
	"strconv"
)

func parseDisk(d []byte) []int {
	ret := make([]int, len(d))

	for i, b := range d {
		s := string(b)
		v, _ := strconv.Atoi(s)
		ret[i] = v

	}
	return ret
}

func disk(diskMap []int) []int {
	ret := []int{}
	isFile := true
	fileId := 0
	var v int
	for i := 0; i < len(diskMap); i++ {
		if isFile {
			v = fileId
			fileId++
		} else {
			v = -1
		}
		for j := 0; j < diskMap[i]; j++ {
			ret = append(ret, v)
		}
		isFile = !isFile
	}
	return ret
}

func compact(d []int) []int {
	for f, b := 0, len(d)-1; f <= b; {
		if d[f] != -1 {
			f++
		} else {
			if d[b] == -1 {
				b--
			} else {
				//swap
				d[f], d[b] = d[b], -1
			}
		}
	}
	return d
}

func checksum(compact []int) int64 {
	var s int64
	s = 0
	for i, v := range compact {
		if v != -1 {
			s += int64(i * v)
		}
	}
	return s
}

func Run() {
	file, _ := os.Open("../internal/day9/input")
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	diskMap := parseDisk(line)
	disk := disk(diskMap)
	compact := compact(disk)
	println(checksum(compact))
}
