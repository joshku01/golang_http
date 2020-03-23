package main

import (
	"github.com/gin-gonic/gin"
	"golang_http/Route"
)

func main() {

	r := gin.Default()
	Route.Router(r)
	r.Run(":8082")
}
