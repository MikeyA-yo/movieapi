package main

import (
	"encoding/json"
	"io"
)

func Add(a int, b int) int {
	return a + b
}

func DecodeJson(body io.ReadCloser, target *fiction) error {
	decode := json.NewDecoder(body)
	err := decode.Decode(&target)
	return err
}
