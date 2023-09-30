package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

// ripped from the Go
// src.
func callerPC(depth int) uintptr {
	var pcs [1]uintptr
	runtime.Callers(depth, pcs[:])
	return pcs[0]
}

func Err(err error) error {

	u := callerPC(3)
	fs := runtime.CallersFrames([]uintptr{u})
	f, _ := fs.Next()

	return errors.New(
		fmt.Sprintf(
			"error : %s [%s:%v [%s]]",
			err.Error(),
			f.File,
			f.Line,
			f.Function,
		),
	)
}

func main() {

	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	err1 := errors.New("Error 1")

	log.Println(
		Err(err1),
	)

	err2 := Err(errors.New("Error 2"))

	return err2
}
