package codec

import (
	"github.com/json-iterator/go"
)

var JSON = new(jsonIterCodec)

var json = jsoniter.ConfigFastest

type jsonIterCodec int

func (j jsonIterCodec) Name() string {
	return "jsoniter"
}

func (j jsonIterCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsonIterCodec) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
