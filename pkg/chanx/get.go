package chanx

func TryGet[T interface{}](channel chan T) (T, bool) {
	select {
	case value := <-channel:
		return value, true
	default:
		var empty T
		return empty, false
	}
}
