package utils

func Select[T any, TOut any](source []T, selector func(T) TOut) []TOut {
	res := make([]TOut, 0)

	for _, t := range source {
		res = append(res, selector(t))
	}

	return res
}
