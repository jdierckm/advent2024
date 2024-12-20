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

type slot struct {
	loc, size int
}

func openSlots(d []int) []slot {
	ret := []slot{}
	size := 0
	loc := 0
	for i, v := range d {
		if v != -1 {
			if size > 0 {
				ret = append(ret, slot{loc: loc, size: size})
				size = 0
			}
		} else {
			if size == 0 {
				loc = i
			}
			size++
		}
	}
	return ret
}

type file struct {
	num, len, loc int
}

func fileLocs(d []int) []file {
	ret := []file{}
	l := 0
	last := -1
	for i := len(d) - 1; i >= 0; i-- {
		di := d[i]
		if di != last {
			if last != -1 {
				ret = append(ret, file{num: last, len: l, loc: i + 1})
			}
			l = 0
		}
		last = d[i]
		l++
	}
	return ret
}

func getSlot(slots []slot, length int) int {
	for i, v := range slots {
		if v.size >= length {
			r := v.loc
			slots[i] = slot{loc: v.loc + length, size: v.size - length}
			return r
		}
	}
	return -1
}

func compact2(d []int, slots []slot, fileLocs []file) {
	for _, f := range fileLocs {
		idx := getSlot(slots, f.len)
		if idx > 0 {
			for i := 0; i < f.len; i++ {
				d[idx+i] = f.num
				d[f.loc+i] = -1
			}
		}
	}
}

func compact2_2(slots []slot, fileLocs []file) []file {
	ret := make([]file, len(fileLocs))
	copy(ret, fileLocs)

	for i, f := range ret {
		idx := getSlot(slots, f.len)
		if idx > 0 {
			ret[i] = file{len: f.len, loc: idx, num: f.num}
		}
	}
	return ret
}

func checksum2(files []file) int {

	s := 0
	for _, f := range files {
		for i := 0; i < f.len; i++ {
			s += f.num * (f.loc + i)
		}
	}
	return s
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
	line, _ := reader.ReadString('\n')
	diskMap := parseDisk([]byte(line))
	dsk := disk(diskMap)
	compact := compact(dsk)
	println(checksum(compact))

	//part2
	diskMap = parseDisk([]byte(line))
	dsk = disk(diskMap)
	//fmt.Printf("map: %v\n", dsk)
	slots := openSlots(dsk)
	//fmt.Printf("slots: %v\n", slots)
	fileLocs := fileLocs(dsk)
	//fmt.Printf("files: %v\n", fileLocs)
	compact2(dsk, slots, fileLocs)
	println(checksum(dsk))

	// c2 := compact2_2(slots, fileLocs)
	// fmt.Printf("c2: %v\n", c2[:20])
	// println(checksum2(c2))
}
