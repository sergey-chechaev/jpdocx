package main

import (
	"encoding/json"
  "fmt"
	"jpdocx/docx"
  "os"
)

func main(){
  params := os.Args[1]
	path := os.Args[2]
	path_to := os.Args[3]
	path_str := fmt.Sprintf("%v", path)
	path_to_str := fmt.Sprintf("%v", path_to)
  data := []byte(params)

  objmap := make(map[string]interface{})
  err := json.Unmarshal(data, &objmap)
  if err != nil {
    fmt.Println(err)
  }

	r, err := docx.ReadDocxFile(path_str)
	if err != nil {
		panic(err)
	}

	docx1 := r.Editable()
  for key, value := range objmap {
		str_key := fmt.Sprintf("{{%v}}", key)
		str_value := fmt.Sprintf("%v", value)
		docx1.Replace(str_key, str_value, -1)
  }
	docx1.WriteToFile(path_to_str)

	r.Close()
	fmt.Println("ok")
}
