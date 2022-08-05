package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	oldSlice := []int{4, 8, 1, 5, 2, 9}
	expectedResult := []int{9, 8, 5, 4, 2, 1}
	slice := oldSlice

	BubbleSort(slice)

	for i, v := range expectedResult {
		if v != slice[i] {
			t.Errorf("expected [%d]%d and received [%d]%d", i, v, i, slice[i])
			break
		}
	}
}
