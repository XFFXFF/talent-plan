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
		merge(src, lTemp, rTemp, mid)
	}
}

func merge(src, lTemp, rTemp []int64, mid int) {
	copy(lTemp, src[:mid])
	copy(rTemp, src[mid:])

	i := 0
	for len(lTemp) > 0 && len(rTemp) > 0 {
		if lTemp[0] < rTemp[0] {
			src[i] = lTemp[0]
			lTemp = lTemp[1:]
		} else {
			src[i] = rTemp[0]
			rTemp = rTemp[1:]
		}
		i++
	}

	for j := 0; j < len(lTemp); j++ {
		src[i] = lTemp[j]
		i++
	}

	for j := 0; j < len(rTemp); j++ {
		src[i] = rTemp[j]
		i++
	}
}
