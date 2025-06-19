package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func Serialize(data any) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize data: %w", err)
	}
	return buf.Bytes(), nil
}

func Deserialize(data []byte, out any) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(out)
	if err != nil {
		return fmt.Errorf("failed to deserialize data: %w", err)
	}
	return nil
}
