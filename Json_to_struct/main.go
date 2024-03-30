package main

import (
	"encoding/json"
	"fmt"
)

type mystruct struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	jsondata := []byte(`{"name":"kanagavalli","age":25}`)
	obj := mystruct{}
	err := json.Unmarshal(jsondata, &obj)
	fmt.Println(err, "error")
	if err != nil {
		fmt.Println(err)
	}

	jsondata, err = json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsondata))
}
