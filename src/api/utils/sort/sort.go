package sort

import "sort"

func Bubble(slice []int) {
	next := true
	for next {
		next = false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				next = true
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}

func Sort(slice []int) {
	sort.Ints(slice)
}

func CreateSliceWithNElements(n int) (result []int) {
	result = make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return
}
