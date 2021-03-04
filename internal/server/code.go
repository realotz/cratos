package server

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
)

// codec is a Codec implementation with json.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return jsoniter.Marshal(m)
	}
	return json.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return jsoniter.Unmarshal(data, m)
	}
	return json.Unmarshal(data, v)
}

func (codec) Name() string {
	return "json"
}
