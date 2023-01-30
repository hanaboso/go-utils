package arrayx

func NthItems[T any](haystack []T, step int) []T {
	return NthItemsFrom(haystack, step, step-1)
}

func NthItemsFrom[T any](haystack []T, step, firstIndex int) []T {
	result := []T{}
	for i := firstIndex; i < len(haystack); i += step {
		result = append(result, haystack[i])
	}

	return result
}
