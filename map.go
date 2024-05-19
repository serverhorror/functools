package functools

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
