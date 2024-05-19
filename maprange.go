//go:build goexperiment.rangefunc

package functools

import (
	"iter"
	"log/slog"
)

// MapRange applies a mapping function to each element in a sequence and returns a new sequence
// containing the mapped values.
//
// The mapping function `m` takes an element of type `T` from the input sequence and returns
// a value of type `Out`. If an error occurs during the mapping process, the error is logged
// and the iteration is stopped.
//
// The input sequence `seq` should implement the `iter.Seq` interface, which allows iteration
// over its elements.
//
// The returned sequence is an iterator function that can be used to iterate over the mapped
// values. The iterator function takes a `yield` function as an argument, which is called for
// each mapped value. If the `yield` function returns `false`, the iteration is stopped.
//
// Example usage:
//
//	seq := []int{1, 2, 3, 4, 5}
//	double := func(x int) (int, error) {
//	  return x * 2, nil
//	}
//	mappedSeq := MapRange(double, seq)
//	for v := range mappedSeq {
//	  fmt.Println(v)
//	}
//
// Output:
//
//	2
//	4
//	6
//	8
//	10
func MapRange[T any, Out any](m MapFn[T, Out], seq iter.Seq[T]) iter.Seq[Out] {

	iterator := func(yield func(Out) bool) {
		for v := range seq {
			r, err := m(v)
			if err != nil {
				slog.Error("error in iterator", "err", err)
				break
			}
			ok := yield(r)
			if !ok {
				slog.Info("yielded from iterator", "r", r, "ok", ok)
				return
			}
		}
	}
	return iterator
}

// MapRangePull applies a mapping function to each element in a sequence and returns a pull-based iterator.
//
// The mapping function `m` takes an element of type `T` from the sequence and returns an element of type `Out`.
// The sequence `seq` is an iterable sequence of elements of type `T`.
// The returned pull-based iterator consists of two functions:
//
//   - The first function is a generator that yields the mapped elements one by one until the sequence is exhausted or the generator is stopped.
//   - The second function is a closer that can be called to stop the generator and perform any necessary cleanup.
func MapRangePull[T any, Out any](m MapFn[T, Out], seq iter.Seq[T]) (func() (Out, bool), func()) {

	iterator := func(yield func(Out) bool) {
		for v := range seq {
			r, err := m(v)
			if err != nil {
				slog.Error("error in iterator", "err", err)
				break
			}
			ok := yield(r)
			if !ok {
				slog.Info("yielded from iterator", "r", r, "ok", ok)
				return
			}
		}
	}
	slog.Info("created the iterator", "iterator", iterator)
	return iter.Pull(iterator)
}
