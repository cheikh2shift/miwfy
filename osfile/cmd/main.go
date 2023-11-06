package main

import (
	"os"

	file "github.com/cheikh2shift/miwfy/osfile"
)

func main() {

	f, err := os.OpenFile(
		"./test.txt",
		os.O_RDWR|os.O_CREATE,
		0755,
	)

	if err != nil {
		panic(err)
	}

	data, err := file.Get(f)

	str := string(data) + "+"

	err = file.Save(f, str)

	if err != nil {
		panic(err)
	}

}

func bruteMode() {

	f, err := os.OpenFile(
		"./test.txt",
		os.O_RDWR|os.O_CREATE,
		0755,
	)

	if err != nil {
		panic(err)
	}

	data, err := file.Get(f)

	str := string(data) + "+"

	err = file.Save(f, str)

	if err != nil {
		panic(err)
	}

}

func appendMode() {

	f, err := os.OpenFile(
		"./test.txt",
		os.O_RDWR,
		os.ModeAppend,
	)

	if err != nil {
		panic(err)
	}

	err = file.Save(f, "+")

	if err != nil {
		panic(err)
	}

}
