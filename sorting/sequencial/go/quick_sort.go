package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func quick_sort(bottom int, top int, ary []int) []int {
	if bottom >= top {
		return ary
	}
	div := ary[bottom]
	lower := bottom
	upper := top
	for lower < upper {
		for lower <= upper && ary[lower] <= div {
			lower++
		}
		for lower <= upper && ary[upper] > div {
			upper--
		}
		if lower < upper {
			tmp := ary[lower]
			ary[lower] = ary[upper]
			ary[upper] = tmp
		}
	}
	ary[bottom] = ary[upper]
	ary[upper] = div
	quick_sort(bottom, upper-1, ary)
	quick_sort(upper+1, top, ary)
	return ary
}

var (
	size = 1000000
)

func main() {
	fp, _ := os.Open("../../../data/" + strconv.Itoa(size) + "/seed0.txt")
	scanner := bufio.NewScanner(fp)
	ary1 := make([]int, size)
	i := 0
	for scanner.Scan() {
		ary1[i], _ = strconv.Atoi(scanner.Text())
		i++
	}
	fp, _ = os.Open("../../../data/" + strconv.Itoa(size) + "/seed1.txt")
	scanner = bufio.NewScanner(fp)
	ary2 := make([]int, size)
	i = 0
	for scanner.Scan() {
		ary2[i], _ = strconv.Atoi(scanner.Text())
		i++
	}
	fp, _ = os.Open("../../../data/" + strconv.Itoa(size) + "/seed2.txt")
	scanner = bufio.NewScanner(fp)
	ary3 := make([]int, size)
	i = 0
	for scanner.Scan() {
		ary3[i], _ = strconv.Atoi(scanner.Text())
		i++
	}

	sTime := time.Now()
	quick_sort(0, size-1, ary1)
	quick_sort(0, size-1, ary2)
	quick_sort(0, size-1, ary3)
	eTime := time.Now()
	fmt.Println(eTime.Sub(sTime) / 3)
}
