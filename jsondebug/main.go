package main

import (
	"encoding/json"
	"log"
)

type RowData struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (m *RowData) UnmarshalJSON(d []byte) error {

	log.Println("Data passed:", string(d))
	// do decrypting...
	m.Name = "test"
	m.Count = 100

	return nil
}

func main() {

	blob := `[
		"SCRAMBLEDDATA",
		"SCRAMBLEDDATA"	
	]`

	var inventory []RowData

	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	log.Println(inventory)

}
