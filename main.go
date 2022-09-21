package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router := Router{routeBinder: r}
	if !router.BindRoute() {
		log.Println("bind url error")
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
