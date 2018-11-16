package helper

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

func Dump(data ...interface{}) {
	DumpYAML(data...)
}

func DumpJSON(data ...interface{}) {
	for _, item := range data {
		d, _ := json.MarshalIndent(item, " ", "  ")
		fmt.Println(string(d))
	}
}

func DumpYAML(data ...interface{}) {
	for _, item := range data {
		d, _ := yaml.Marshal(item)
		fmt.Println(string(d))
	}
}
