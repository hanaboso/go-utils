package mapx

func Values[M ~map[K]V, K comparable, V any](data M) []V {
	result := make([]V, 0, len(data))
	for _, value := range data {
		result = append(result, value)
	}

	return result
}

func Keys[M ~map[K]V, K comparable, V any](data M) []K {
	result := make([]K, 0, len(data))
	for key := range data {
		result = append(result, key)
	}

	return result
}
