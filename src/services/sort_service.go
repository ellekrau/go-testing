package services

import "go-testing/utils/sort"

func Sort(elements []int) {
	if len(elements) < 10000 {
		sort.Bubble(elements)
		return
	}
	sort.Sort(elements)
}
