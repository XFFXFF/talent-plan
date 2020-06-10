package main

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	length := len(src)
	if length > 1 {
		mid := length / 2
		MergeSort(src[:mid])
		MergeSort(src[mid:])
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
