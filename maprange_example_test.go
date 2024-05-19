package functools_test

import (
	"fmt"
	"iter"

	"github.com/serverhorror/functools"
)

func fakeIterator[T any](s []T) iter.Seq[T] {
	sliceIter := func(slice []T) iter.Seq[T] {
		return func(yield func(T) bool) {
			for _, v := range slice {
				if !yield(v) {
					break
				}
			}
		}
	}
	return sliceIter(s)

}

func ExampleMapRange() {
	seq := []int{1, 2, 3, 4, 5}
	expect := []int{2, 4, 6, 8, 10}
	got := make([]int, 0, len(expect))
	double := func(x int) (int, error) {
		return x * 2, nil
	}
	mappedSeq := functools.MapRange(double, fakeIterator(seq))
	for v := range mappedSeq {
		got = append(got, v)
	}
	fmt.Printf("%#v", got)
	// Output: []int{2, 4, 6, 8, 10}
}
