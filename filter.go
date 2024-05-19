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
