package utils

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func StudentToStudentForJson(fo, to interface{}) error {
	fmt.Println(fo)
	fmt.Println(to)
	b, err := jsoniter.Marshal(fo)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(b, to)
}
