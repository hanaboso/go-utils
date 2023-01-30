package arrayx

func RemoveAt[T any](haystack []T, index int) []T {
	result := make([]T, 0)
	result = append(result, haystack[:index]...)

	return append(result, haystack[index+1:]...)
}
