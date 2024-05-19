package functools

type FilterFn[T any] func(T) (bool, error)
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

type MapFn[T any, Out any] func(T) (Out, error)
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

type ReduceFn[T any, Out any] func(Out, T) (Out, error)
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
