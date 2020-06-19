package main

import (
	"runtime"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	grLimit := runtime.NumCPU() - 1
	grLimitChan := make(chan struct{}, grLimit)
	temp := make([]int64, len(src))
	mergeSort(src, temp, grLimitChan)
}

func mergeSort(src, temp []int64, grLimitChan chan struct{}) {
	wg := sync.WaitGroup{}
	length := len(src)
	if length > 1 {
		mid := length / 2

		left, right := src[:mid], src[mid:]
		lTemp, rTemp := temp[:mid], temp[mid:]

		select {
		case grLimitChan <- struct{}{}:
			wg.Add(1)
			go func() {
				defer wg.Done()
				mergeSort(left, lTemp, grLimitChan)
			}()
		default:
			mergeSort(left, lTemp, grLimitChan)
		}
		mergeSort(right, rTemp, grLimitChan)
		wg.Wait()
		merge(src, temp, left, right)
	}
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
