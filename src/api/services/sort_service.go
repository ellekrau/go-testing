package services

import "go-testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) < 10000 {
		sort.Bubble(elements)
		return
	}
	sort.Sort(elements)
}
