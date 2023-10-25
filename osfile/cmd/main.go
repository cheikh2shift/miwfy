package main

import (
	"fmt"
	"os"

	file "github.com/cheikh2shift/miwfy/osfile"
)

func main() {

	f, err := os.OpenFile(
		"./test.txt",
		os.O_RDWR,
		os.ModeAppend,
	)

	if err != nil {
		panic(err)
	}

	data, err := file.Get(f)

	fmt.Printf("Current Entries:\n%s\n", data)

	err = file.Save(f, "+")

	if err != nil {
		panic(err)
	}

}
