package main

import (
	"sort"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	temp := make([]int64, len(src))
	mergeSort(src, temp)
}

func mergeSort(src, temp []int64) {
	wg := sync.WaitGroup{}
	length := len(src)
	if length <= 10000 {
		sort.Slice(src, func(i, j int) bool { return src[i] <= src[j] })
		return
	}
	mid := length / 2

	left, right := src[:mid], src[mid:]
	lTemp, rTemp := temp[:mid], temp[mid:]

	wg.Add(1)
	go func() {
		defer wg.Done()
		mergeSort(left, lTemp)
	}()
	mergeSort(right, rTemp)
	wg.Wait()
	merge(src, temp, left, right)
}

func merge(src, temp, left, right []int64) {
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			temp[i] = left[0]
			left = left[1:]
		} else {
			temp[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		temp[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		temp[i] = right[j]
		i++
	}

	copy(src, temp)
}
