package sort

import "testing"

func TestBubble(t *testing.T) {
	slice := []int{4, 8, 1, 5, 2, 9}
	expectedResult := []int{1, 2, 4, 5, 8, 9}

	Bubble(slice)

	for i, v := range expectedResult {
		if v != slice[i] {
			t.Errorf("expected [%d]%d and received [%d]%d", i, v, i, slice[i])
			break
		}
	}
}

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

func getElements(n int) (result []int) {
	result = make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}
	return
}

// slice := []int{4, 8, 1, 5, 2, 9}
// BenchmarkBubble-16      365.397.691                3.775 ns/op
// BenchmarkBubble-16      360.098.990                3.135 ns/op
// BenchmarkBubble-16      358.633.426                3.745 ns/op
// slice := getElements(10000) 10k
// BenchmarkBubble-16        209191              6927 ns/op
// BenchmarkBubble-16        166887              6372 ns/op
// BenchmarkBubble-16        277129              5629 ns/op
// slice := getElements(100000) 100k
// BenchmarkBubble-16             1        15039330.900 ns/op
// BenchmarkBubble-16             1        13167436300 ns/op
// BenchmarkBubble-16             1        13713001000 ns/op
func BenchmarkBubble(b *testing.B) {
	//slice := []int{4, 8, 1, 5, 2, 9}
	//slice := getElements(100000)
	slice := getElements(10000)

	for i := 0; i < b.N; i++ {
		Bubble(slice)
	}
}

// slice := []int{4, 8, 1, 5, 2, 9}
// BenchmarkSort-16        15.596.385                79.59 ns/op
// BenchmarkSort-16        12.949.593                92.92 ns/op
// BenchmarkSort-16        16.577.996                76.68 ns/
// slice := getElements(10000) 10k
// BenchmarkSort-16            2886            497202 ns/op
// BenchmarkSort-16            3091            552293 ns/op
// BenchmarkSort-16            2904            463438 ns/op
// slice := getElements(100000) 100k
// BenchmarkSort-16             196           6758528 ns/op
// BenchmarkSort-16             211           5571500 ns/op
// BenchmarkSort-16             235           5470146 ns/op
func BenchmarkSort(b *testing.B) {
	//slice := []int{4, 8, 1, 5, 2, 9}
	//slice := getElements(100000)
	slice := getElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}
