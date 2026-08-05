package main

import "log"

const email = ""
const password = ""

func init() {
	if email == "" || password == "" {
		log.Fatalln("Please set your email and password in credentials.go")
	}
}