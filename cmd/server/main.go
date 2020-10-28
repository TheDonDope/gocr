package main

import (
	gocrHttp "github.com/TheDonDope/gocr/pkg/http"
	"github.com/gin-gonic/gin"
)

func main() {
	gocrHttp.SetAuthVariables()
	r := gin.Default()
	r.Use(gocrHttp.CORSMiddleware())

	authorized := r.Group("/gocr")
	// authorized.Use(gocrHttp.AuthRequired())
	authorized.GET("healthy", gocrHttp.Healthy)

	err := r.Run(":4242")
	if err != nil {
		panic(err)
	}

}
