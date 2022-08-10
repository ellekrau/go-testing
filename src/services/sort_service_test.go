package services

import (
	"go-testing/utils/sort"
	"testing"
)

func TestSort(t *testing.T) {
	slice := []int{4, 8, 1, 5, 2, 9}
	expectedResult := []int{1, 2, 4, 5, 8, 9}

	Sort(slice)

	for i, v := range expectedResult {
		if v != slice[i] {
			t.Errorf("expected [%d]%d and received [%d]%d", i, v, i, slice[i])
			break
		}
	}
}

func TestSortLenBiggerThan10k(t *testing.T) {
	slice := sort.CreateSliceWithNElements(10001)

	Sort(slice)

	if slice[0] != 0 {
		t.Errorf("first slice element should be '0'")
	}

	if slice[10000] != 10000 {
		t.Errorf("last slice element should be '10000'")
	}
}

func TestSortLenEquals10k(t *testing.T) {
	slice := sort.CreateSliceWithNElements(10000)

	Sort(slice)

	if slice[0] != 0 {
		t.Errorf("first slice element should be '0'")
	}

	if slice[9999] != 9999 {
		t.Errorf("last slice element should be '10000'")
	}
}

func TestSortLenSmallerThan10k(t *testing.T) {
	slice := sort.CreateSliceWithNElements(9999)

	Sort(slice)

	if slice[0] != 0 {
		t.Errorf("first slice element should be '0'")
	}

	if slice[9998] != 9998 {
		t.Errorf("last slice element should be '10000'")
	}
}
