// Package functools provides utility functions for functional programming in Go.
package functools

// FilterFn is a function type that takes a value of type T and returns a boolean value
// indicating whether the value should be filtered or not.
//
// It also returns an error if any error occurs during the filtering process.
type FilterFn[T any] func(T) (bool, error)

// Filter applies a filter function to each element in the input slice and returns a new slice containing only the elements for which the filter function returns true.
//
// The filter function takes an element of type T as input and returns a boolean value indicating whether the element should be included in the result slice.
// If the filter function returns an error, the function immediately returns nil and the error.
func Filter[T any](f FilterFn[T], s []T) ([]T, error) {
	result := make([]T, 0, len(s))
	var ok bool
	var err error
	for _, v := range s {
		ok, err = f(v)
		if err != nil {
			return nil, err
		}
		if ok {
			result = append(result, v)
		}
	}
	return result, nil
}

// MapFn is a function type that takes an input of type T and returns an output of type Out.
//
// It is used to define the mapping function for the `Map` operation in the `functools` package.
type MapFn[T any, Out any] func(T) (Out, error)

// Map applies the given function to each element in the input slice and
// returns a new slice containing the results.
//
// The function must take an element of type T as input and return a value of type Out.
// If any error occurs during the mapping process, the function returns an empty slice and the error.
func Map[T any, Out any](m MapFn[T, Out], s []T) ([]Out, error) {
	result := make([]Out, 0, len(s))
	var r Out
	var err error
	for _, v := range s {
		r, err = m(v)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// ReduceFn takes an accumulator of type `Out` and an element of type `T`, and returns the updated accumulator and an error (if any).
type ReduceFn[T any, Out any] func(Out, T) (Out, error)

// Reduce applies a function to each element in the slice and returns a single value.
//
// The function takes two arguments: the current accumulator value and the current element from the slice.
// It returns the updated accumulator value and an error, if any.
// The initial value of the accumulator is the zero value of the specified type.
// If an error occurs during the reduction process, the function stops and returns the zero value of the accumulator and the error.
// Otherwise, it returns the final accumulator value and a nil error.
func Reduce[T any, Out any](r ReduceFn[T, Out], s []T) (Out, error) {
	var accumulator Out
	var err error
	for _, v := range s {
		accumulator, err = r(accumulator, v)
		if err != nil {
			var acc Out
			return acc, err
		}
	}
	return accumulator, nil
}
