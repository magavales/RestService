package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var port *string

func init() {
	port = flag.String("port", "8080", "Port on which server will listen for requests")
}

func main() {
	flag.Parse()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ping: Well done", "well done")
		log.Println("msg accepted")
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Printf("[server error]: %s\n", err)
	}
}
