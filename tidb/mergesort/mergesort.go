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
	mergeSort(src, grLimitChan)
}

func mergeSort(src []int64, grLimitChan chan struct{}) {
	wg := sync.WaitGroup{}
	length := len(src)
	if length > 1 {
		mid := length / 2

		select {
		case grLimitChan <- struct{}{}:
			wg.Add(1)
			go func() {
				defer wg.Done()
				mergeSort(src[:mid], grLimitChan)
			}()
		default:
			mergeSort(src[:mid], grLimitChan)
		}
		mergeSort(src[mid:], grLimitChan)
		wg.Wait()
		merge(src, mid)
	}
}

func merge(src []int64, mid int) {
	leftSrc := make([]int64, mid)
	rightSrc := make([]int64, len(src)-mid)
	copy(leftSrc, src[:mid])
	copy(rightSrc, src[mid:])

	i := 0
	for len(leftSrc) > 0 && len(rightSrc) > 0 {
		if leftSrc[0] < rightSrc[0] {
			src[i] = leftSrc[0]
			leftSrc = leftSrc[1:]
		} else {
			src[i] = rightSrc[0]
			rightSrc = rightSrc[1:]
		}
		i++
	}

	for j := 0; j < len(leftSrc); j++ {
		src[i] = leftSrc[j]
		i++
	}

	for j := 0; j < len(rightSrc); j++ {
		src[i] = rightSrc[j]
		i++
	}
}
