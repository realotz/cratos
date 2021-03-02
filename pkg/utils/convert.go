package utils

import jsoniter "github.com/json-iterator/go"

func StudentToStudentForJson(fo, to interface{}) error {
	b, err := jsoniter.Marshal(fo)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(b, to)
}
