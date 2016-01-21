package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func bubble_sort(ary []int) []int {
	done := make([][]bool, size)
	for i := 0; i < size; i++ {
		done[i] = make([]bool, size)
	}
	tasks := make([]int, size-1)
	for i := 0; i < size-1; i++ {
		tasks[i] = i + 1
	}
	var mu = &sync.Mutex{}
	var wg sync.WaitGroup

	for i := 0; i < workload; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				var k int
				mu.Lock()
				if len(tasks) < 1 {
					mu.Unlock()
					wg.Done()
					return
				}
				k, tasks = tasks[len(tasks)-1], tasks[:len(tasks)-1]
				mu.Unlock()
				// fmt.Println(k)
				for j := 0; j < k; j++ {
					// fmt.Println("doing["+strconv.Itoa(k)+", "+strconv.Itoa(j+1)+"]")
					for done[k][j+1] == false && k != size-1 {
					}
					if ary[j] > ary[j+1] {
						tmp := ary[j]
						ary[j] = ary[j+1]
						ary[j+1] = tmp
					}
					done[k-1][j] = true
				}
			}
		}(i)
	}
	wg.Wait()
	// fmt.Println(ary)
	return ary
}

var (
	size     = 100000
	workload = 4
)

func main() {
	fp, _ := os.Open("../data/" + strconv.Itoa(size) + "/seed0.txt")
	scanner := bufio.NewScanner(fp)
	ary1 := make([]int, size)
	i := 0
	for scanner.Scan() {
		ary1[i], _ = strconv.Atoi(scanner.Text())
		i++
	}
	fp, _ = os.Open("../data/" + strconv.Itoa(size) + "/seed1.txt")
	scanner = bufio.NewScanner(fp)
	ary2 := make([]int, size)
	i = 0
	for scanner.Scan() {
		ary2[i], _ = strconv.Atoi(scanner.Text())
		i++
	}
	fp, _ = os.Open("../data/" + strconv.Itoa(size) + "/seed2.txt")
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
