package main

import (
	"errors"
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := HogAuth("localhost")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("localhost:1025", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

type hogAuth struct {
	host string
}

func HogAuth(host string) smtp.Auth {
	return &hogAuth{host}
}

func (a *hogAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {

	if server.Name != a.host {
		return "", nil, errors.New("wrong host name")
	}
	resp := []byte("HELLO\x00HELLO\x00HELLO")
	return "PLAIN", resp, nil
}

func (a *hogAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// We've already sent everything.
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}
