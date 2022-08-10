package sort

import (
	"testing"
	"time"
)

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

func TestBubbleWithTimeout(t *testing.T) {
	reachedTimeoutChannel := make(chan bool, 1)
	defer close(reachedTimeoutChannel)

	go func() {
		Bubble(CreateSliceWithNElements(10))
		reachedTimeoutChannel <- false
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		reachedTimeoutChannel <- true
	}()

	if <-reachedTimeoutChannel {
		t.Errorf("the bubblesort execution reached 50 ms timeout")
	}
}

// go test --bench=. -count 5
// slice := []int{4, 8, 1, 5, 2, 9}
// BenchmarkBubble-16      365.397.691                3.775 ns/op
// BenchmarkBubble-16      360.098.990                3.135 ns/op
// BenchmarkBubble-16      358.633.426                3.745 ns/op
// slice := getElements(10000) 10k
// BenchmarkBubble-16        209191              6927 ns/op
// BenchmarkBubble-16        166887              6372 ns/op
// BenchmarkBubble-16        277129              5629 ns/op
// slice := getElements(100000) 100k
// BenchmarkBubble-16             1        15039330900 ns/op
// BenchmarkBubble-16             1        13167436300 ns/op
// BenchmarkBubble-16             1        13713001000 ns/op
func BenchmarkBubble(b *testing.B) {
	//slice := []int{4, 8, 1, 5, 2, 9}
	//slice := getElements(100000)
	slice := CreateSliceWithNElements(10000)

	for i := 0; i < b.N; i++ {
		Bubble(slice)
	}
}

// go test --bench=. -count 5
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
	slice := CreateSliceWithNElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}
