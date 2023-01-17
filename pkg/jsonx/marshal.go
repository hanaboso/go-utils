package jsonx

import "encoding/json"

func UnmarshalBytes[T any](data []byte) (*T, error) {
	out := new(T)
	if err := json.Unmarshal(data, out); err != nil {
		return nil, err
	}
	return out, nil
}

func UnmarshalString[T any](data string) (*T, error) {
	out := new(T)
	if err := json.Unmarshal([]byte(data), out); err != nil {
		return nil, err
	}
	return out, nil
}
