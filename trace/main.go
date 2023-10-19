package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/trace"
	_ "golang.org/x/net/trace"
)

const APP = "TekDeckDude"

type myHandlerSign func(
	w http.ResponseWriter,
	req *http.Request,
	t trace.Trace,
)

func hello(
	w http.ResponseWriter,
	req *http.Request,
	t trace.Trace,
) {

	// simulate a call to db

	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		t.LazyPrintf("Running %v", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "hello\n")
}

func error(
	w http.ResponseWriter,
	req *http.Request,
	t trace.Trace,
) {

	// simulate a call to db
	time.Sleep(1 * time.Second)

	t.SetError()

	fmt.Fprintf(w, "error\n")
}

func Adapt(fn myHandlerSign) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		req *http.Request,
	) {

		tr := trace.New(APP, req.URL.Path)
		defer tr.Finish()

		tr.LazyPrintf("Request started")
		fn(w, req, tr)
		tr.LazyPrintf("Request Ended")
	}
}

func main() {

	http.HandleFunc(
		"/hello",
		Adapt(hello),
	)

	http.HandleFunc(
		"/error",
		Adapt(error),
	)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		log.Panic(err)
	}
}
