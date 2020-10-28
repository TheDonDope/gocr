package main

import (
	gocrHttp "github.com/TheDonDope/gocr/pkg/http"
)

func main() {

	r := gocrHttp.NewServer()
	err := r.Run(":4242")
	if err != nil {
		panic(err)
	}

}
