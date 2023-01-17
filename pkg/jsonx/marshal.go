package jsonx

import "encoding/json"

func unmarshalBytes[T any](data []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func unmarshalString[T any](data string) (*T, error) {
	out := new(T)
	if err := json.Unmarshal([]byte(data), out); err != nil {
		return nil, err
	}
	return out, nil
}
