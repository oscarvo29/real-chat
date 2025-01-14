package utils

import (
	"encoding/json"
	"io"
)

func ParseJsonObject[T any](jsonData io.ReadCloser, dataPointer *T) error {
	body, err := io.ReadAll(jsonData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &dataPointer)
	return err
}
