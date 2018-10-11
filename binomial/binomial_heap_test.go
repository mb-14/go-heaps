package binomial

import (
	"sort"
	"testing"

	"github.com/theodesp/go-heaps"
)

func TestLeftistHeapInteger(t *testing.T) {
	heap := &BinomialHeap{}

	numbers := []int{4, 3, 2, 5}

	for _, number := range numbers {
		heap.Insert(Int(number))
	}

	sort.Ints(numbers)

	for _, number := range numbers {
		if Int(number) != heap.DeleteMin().(go_heaps.Integer) {
			t.Fail()
		}
	}
}

func TestLeftistHeapString(t *testing.T) {
	heap := &BinomialHeap{}

	strs := []string{"a", "ccc", "bb", "d"}

	for _, str := range strs {
		heap.Insert(Str(str))
	}

	sort.Strings(strs)

	for _, str := range strs {
		if Str(str) != heap.DeleteMin().(go_heaps.String) {
			t.Fail()
		}
	}
}

func Int(value int) go_heaps.Integer {
	return go_heaps.Integer(value)
}

func Str(value string) go_heaps.String {
	return go_heaps.String(value)
}
