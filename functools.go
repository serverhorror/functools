package functools

func Filter[T any](f func(T) (bool, error), s []T) ([]T, error) {
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

func Map[T any, Out any](m func(T) (Out, error), s []T) ([]Out, error) {
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

func Reduce[T any, Out any](r func(Out, T) (Out, error), s []T) (Out, error) {
	var accumulator Out
	var err error
	for _, v := range s {
		accumulator, err = r(accumulator, v)
		if err != nil {
			return accumulator, err
		}
	}
	return accumulator, nil
}
