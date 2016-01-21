package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

func merge(bottom int, top int, wg *sync.WaitGroup) {
	mid := (bottom + top) / 2
	for _, ok1 := store[strconv.Itoa(bottom)+","+strconv.Itoa(mid)]; !ok1; {
	}
	for _, ok2 := store[strconv.Itoa(mid+1)+","+strconv.Itoa(top)]; !ok2; {
		if mid+1 >= size && top >= size {
			break
		}
	}
	ary1 := store[strconv.Itoa(bottom)+","+strconv.Itoa(mid)]
	ary2 := store[strconv.Itoa(mid+1)+","+strconv.Itoa(top)]
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

	delete(store, strconv.Itoa(bottom)+","+strconv.Itoa(mid))
	delete(store, strconv.Itoa(mid+1)+","+strconv.Itoa(top))
	store[strconv.Itoa(bottom)+","+strconv.Itoa(top)] = ary_new
	// fmt.Println(store)
	defer wg.Done()
}

func merge_sort() {
	depth := int(math.Ceil(math.Log2(float64(size))))
	// fmt.Println(depth)

	for i := 0; i < depth; i++ {
		width := int(math.Exp2(float64(i + 1)))
		var wg sync.WaitGroup
		// var mu = &sync.Mutex{}
		for j := 0; j < size-1; j = j + width {
			wg.Add(1)
			// fmt.Println(strconv.Itoa(i)+","+strconv.Itoa(j))
			go merge(j, j+width-1, &wg)
		}
		wg.Wait()
	}
	// 完成
	// fmt.Println(store)
}

var (
	size     = 100000
	workload = 4
	store    = map[string][]int{}
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
	for i := 0; i < len(ary1); i++ {
		store[strconv.Itoa(i)+","+strconv.Itoa(i)] = ary1[i : i+1]
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
	//fmt.Println(store)
	merge_sort()
	// fmt.Println(store["0,"+strconv.Itoa(size-1)])
	//store = map[string][]int{}
	//merge_sort(0, size - 1, ary2)
	// fmt.Println(store["0,"+strconv.Itoa(size-1)])
	//store = map[string][]int{}
	//merge_sort(0, size - 1, ary3)
	eTime := time.Now()
	fmt.Println(eTime.Sub(sTime))
}
