package util

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
)

func ToJson(obj any) []byte {
	rs, err := json.Marshal(obj)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return rs
}

func FromJson[T any](text []byte, target T) T {
	err := json.Unmarshal(text, target)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return target
}

func ToMsgpack(obj any) []byte {
	rs, err := msgpack.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}
	return rs
}

func FromMsgpack[T any](text []byte, target T) T {
	err := msgpack.Unmarshal(text, target)
	if err != nil {
		fmt.Println(err)
	}
	return target
}

func ToYaml(obj any) []byte {
	rs, err := yaml.Marshal(obj)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return rs
}

func FromYaml[T any](text []byte, target T) T {
	err := yaml.Unmarshal(text, target)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return target
}
