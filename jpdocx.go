package main

import (
	"encoding/json"
	"fmt"
	"jpdocx/docx"
	"os"
)

type Hash struct {
	Key   string
	Value string
}

func main() {
  params := os.Args[1]
  bytes := []byte(params)
	path := os.Args[2]
	path_to := os.Args[3]
	path_str := fmt.Sprintf("%v", path)
	path_to_str := fmt.Sprintf("%v", path_to)

	r, err := docx.ReadDocxFile(path_str)
	if err != nil {
		panic(err)
	}
	var hash_keys []Hash
	json.Unmarshal(bytes, &hash_keys)

	docx1 := r.Editable()
	for l := range hash_keys {
		str_key := fmt.Sprintf("{{%v}}", hash_keys[l].Key)
		docx1.Replace(str_key, hash_keys[l].Value, -1)
	}

	docx1.WriteToFile(path_to_str)

	r.Close()
	fmt.Println("ok")
}
