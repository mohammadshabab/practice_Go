package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Parent struct {
	key1 string `json: "key1"`
	key2 string `json: "key2"`
	key3 string `json: "key3"`
}
type JsonData struct {
	JsonData []Parent `json: "parent_node"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to provide json path and excel path")
		fmt.Println("[Usage] [file1.json] [file2.csv]")
		return
	}

	jdata := JsonData{}

	jsonFile := os.Args[1]

	jfile, _ := ioutil.ReadFile(jsonFile)

	fmt.Println("jfile ", string(jfile))
	_ = json.Unmarshal([]byte(jfile), &jdata)

	fmt.Println(jdata)
	//fmt.Println(jsonData)
	for i := 0; i < len(jdata.JsonData); i++ {
		fmt.Println("key1: ", jdata.JsonData[i].key1)
		fmt.Println("key2: ", jdata.JsonData[i].key2)
		fmt.Println("key3: ", jdata.JsonData[i].key3)
	}

//	fmt.Println(Parent.key1)

}
