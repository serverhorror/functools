package functools

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
