package main

import (
	"log"
	"log/syslog"
	"time"
)

func main() {

	d := 2 * time.Second
	progName := "syslog-test"

	logger, err := syslog.New(
		syslog.LOG_INFO,
		progName,
	)

	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(d)
	err = logger.Err("Error entry")

	if err != nil {
		log.Println(err)
	}
	time.Sleep(d)

	err = logger.Notice("Retrying")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(d)

	err = logger.Alert("Op. complete")

	if err != nil {
		log.Println(err)
	}
}
