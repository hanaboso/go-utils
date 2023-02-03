package intx

import (
	"strconv"
)

func Parse(key string) int {
	val, _ := strconv.Atoi(key)
	return val
}

func ParseAny(value interface{}) int {
	switch t := value.(type) {
	case int:
		return t
	case string:
		val, _ := strconv.Atoi(t)
		return val
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	default:
		return 0
	}
}

func ParseDefault(key string, defaultValue int) int {
	val, err := strconv.Atoi(key)
	if err != nil {
		return defaultValue
	}

	return val
}
