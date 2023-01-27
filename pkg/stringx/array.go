package stringx

func InArray(haystack []string, needle string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}

func FromArray(haystack []string, index int) string {
	if len(haystack) > index {
		return haystack[index]
	}

	return ""
}

func FromArrayDefault(haystack []string, index int, defaultValue string) string {
	if len(haystack) > index {
		return haystack[index]
	}

	return defaultValue
}

func NthItems(haystack []string, step int) []string {
	return NthItemsFrom(haystack, step, step-1)
}

func NthItemsFrom(haystack []string, step, firstIndex int) []string {
	result := []string{}
	for i := firstIndex; i < len(haystack); i += step {
		result = append(result, haystack[i])
	}

	return result
}
