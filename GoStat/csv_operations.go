package main

import (
	"encoding/csv"
	"fmt"
  "os"
  "io"
  //"reflect"
)

func main() {
  file, err := os.Open("tesco.csv")
  if err != nil {
    fmt.Println("Error", err)
    return
  }

  defer file.Close()
  reader := csv.NewReader(file)

	// Set FieldsPerRecord to -1 for csv with non-uniform columns
  reader.FieldsPerRecord = -1

  for {
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
        fmt.Println("Error", err)
        return
    }
    //fmt.Println(reflect.TypeOf(record))
    fmt.Println(record)
  }
}
