package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	port = flag.String("port", ":9000", "")
	workerAddr = flag.String("worker address", "127.0.0.1:9001", "")
)


func main()  {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(RequestIDMiddleWare)
	r.Use(AuthMiddleWare)

	r.GET("/signal", Serve)

	err := r.Run(*port)
	if err != nil {
		log.Println("Signal Server boot failed")
	} else {
		log.Println("Signal Server run in port: ", *port)
	}
}