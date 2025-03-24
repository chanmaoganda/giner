package main

import (
	"log"

	"github.com/chanmaoganda/giner/bootstrap"
)

func main() {
	err := bootstrap.Application()

	if err != nil {
		log.Fatal("Bootstrap Failed")
	}
}