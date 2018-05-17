package util

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigDefault

func MarshalJSON(data interface{}) ([]byte, error) {
	if serialized, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return serialized, nil
	}
}

func UnmarshalJSON(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
