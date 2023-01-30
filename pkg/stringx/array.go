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
