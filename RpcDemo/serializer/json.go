package serializer

import (
	"encoding/json"
	"errors"
)

type JsonSerializer struct {
}

var _ Serializer = (*JsonSerializer)(nil)

func (j *JsonSerializer) Marshal(message interface{}) ([]byte, error) {
	var body json.RawMessage
	if message == nil {
		return []byte{}, nil
	}
	var ok bool
	if body, ok = message.(json.RawMessage); !ok {
		return nil, errors.New("message type error")
	}
	return json.Marshal(body)
}

func (j *JsonSerializer) Unmarshal(data []byte, message interface{}) error {
	if data == nil {
		return errors.New("message empty")
	}
	return json.Unmarshal(data, message)
}
