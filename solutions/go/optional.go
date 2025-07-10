package main

import "encoding/json"

type Optional[T any] struct {
	Value T
	Set   bool
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.Set = false
		return nil
	}
	o.Set = true
	return json.Unmarshal(data, &o.Value)
}
