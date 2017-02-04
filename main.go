package main

import (
  "jpdocx/docx"
  "fmt"
  "encoding/json"
  "os"
  "io/ioutil"
)

type Language struct {
  Key  string
  Value string
}

func main() {
  bytes, err_stdin := ioutil.ReadAll(os.Stdin)
  if err_stdin != nil {
    panic(err_stdin)
  }
  path := os.Args[1]
  path_to := os.Args[2]
  path_str := fmt.Sprintf("%v", path)
  path_to_str := fmt.Sprintf("%v", path_to)

  r, err := docx.ReadDocxFile(path_str)
  if err != nil {
      panic(err)
  }

  var languages []Language
  json.Unmarshal(bytes, &languages)

  docx1 := r.Editable()
  for l := range languages {
    str_key := fmt.Sprintf("{{%v}}", languages[l].Key)
    docx1.Replace(str_key, languages[l].Value, -1)
  }

  docx1.WriteToFile(path_to_str)

  r.Close()
  fmt.Println("ok")
}
