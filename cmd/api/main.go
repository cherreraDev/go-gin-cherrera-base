package main

import (
	"go-gin-cherrera-base/cmd/api/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
