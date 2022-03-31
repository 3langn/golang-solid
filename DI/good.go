package main

import "fmt"

type sendMail interface {
	send(string)
}

type gmail struct {
	email string
}

func (g gmail) send(s string) {
	fmt.Println("Sending mail to gmail")
}

type outlook struct {
	email string
}

type order struct {
	mail sendMail
}

func (order order) process(email string) {
	order.mail.send(email)
}

func main() {
	// Dependency Injection
	order := order{mail: gmail{}}
	order.process("hi")
}
