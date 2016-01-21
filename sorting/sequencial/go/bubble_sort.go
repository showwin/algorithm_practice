package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func bubble_sort(ary []int) []int {
	for n := size - 1; n > 0; n-- {
		for i := 0; i < n; i++ {
			if ary[i] > ary[i+1] {
				tmp := ary[i]
				ary[i] = ary[i+1]
				ary[i+1] = tmp
			}
		}
	}
	// fmt.Println(ary)
	return ary
}

var (
	size = 100000
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
	bubble_sort(ary1)
	bubble_sort(ary2)
	bubble_sort(ary3)
	eTime := time.Now()
	fmt.Println(eTime.Sub(sTime) / 3)
}
