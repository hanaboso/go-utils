package intx

import "strconv"

func Parse(key string) int {
	val, _ := strconv.Atoi(key)
	return val
}

func ParseDefault(key string, defaultValue int) int {
	val, err := strconv.Atoi(key)
	if err != nil {
		return defaultValue
	}

	return val
}
