package main

import "fmt"

type SendMail interface {
	send(string)
}

type GMail struct{}

func (G GMail) send(s string) {
	fmt.Println("Sending mail to GMail:", s)
}

type Outlook struct{}

func (o Outlook) send(s string) {
	fmt.Println("Sending mail to Outlook:", s)
}

type Order struct {
}

func (order Order) process() {
	var mail SendMail
	mail = Outlook{} // How to know which struct to use? => bad
	//mail = GMail{}
	mail.send("Hello")
}

func main() {
	Order{}.process()
}
