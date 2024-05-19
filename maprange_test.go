//go:build goexperiment.rangefunc

package functools_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/serverhorror/functools"
)

func prepareSliceIter[T any](t *testing.T, s []T) iter.Seq[T] {
	t.Helper()
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

func TestMapRange(t *testing.T) {

	mapFn := func(x int) (string, error) {
		m := map[int]string{
			1: "one",
			2: "two",
			3: "three",
			4: "four",
			5: "five",
		}
		return m[x], nil
	}

	slice := []int{1, 2, 3, 4, 5}
	expected := []string{"one", "two", "three", "four", "five"}
	seq := prepareSliceIter(t, slice)

	got := make([]string, 0, len(expected))

	for v := range functools.MapRange(mapFn, seq) {
		got = append(got, v)
	}
	if eq := slices.Compare(got, expected); eq != 0 {
		t.Errorf("got %v, expected %v", got, expected)
	}

}

func TestMapRangePull(t *testing.T) {

	mapFn := func(x int) (string, error) {
		m := map[int]string{
			1: "one",
			2: "two",
			3: "three",
			4: "four",
			5: "five",
		}
		return m[x], nil
	}

	slice := []int{1, 2, 3, 4, 5}
	expected := []string{"one", "two", "three", "four", "five"}
	seq := prepareSliceIter(t, slice)

	pull, stop := functools.MapRangePull(mapFn, seq)
	defer stop()

	got := make([]string, 0, len(expected))
	for v, ok := pull(); ok; v, ok = pull() {
		got = append(got, v)
	}
	if eq := slices.Compare(got, expected); eq != 0 {
		t.Errorf("got %v, expected %v", got, expected)
	}

}
