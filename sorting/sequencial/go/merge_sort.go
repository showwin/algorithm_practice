package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func merge_sort(bottom int, top int, ary []int) []int {
	if bottom >= top {
		return ary[bottom : bottom+1]
	}

	mid := (bottom + top) / 2
	ary1 := merge_sort(bottom, mid, ary)
	ary2 := merge_sort(mid+1, top, ary)

	i1 := 0
	l1 := len(ary1)
	i2 := 0
	l2 := len(ary2)

	ary_new := make([]int, l1+l2)
	for i := 0; i < l1+l2; i++ {
		if i1 < l1 && (i2 >= l2 || ary1[i1] <= ary2[i2]) {
			ary_new[i1+i2] = ary1[i1]
			i1++
		} else {
			ary_new[i1+i2] = ary2[i2]
			i2++
		}
	}

	return ary_new
}

var (
	size     = 10
	workload = 4
	store    = map[string][]int{}
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
	merge_sort(0, size-1, ary1)
	merge_sort(0, size-1, ary2)
	merge_sort(0, size-1, ary3)
	eTime := time.Now()
	fmt.Println(eTime.Sub(sTime) / 3)
}
