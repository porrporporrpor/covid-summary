package main

import (
	"log"

	"github.com/porrporporrpor/covid-summary/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := router.Boot(gin.New())
	if err != nil {
		log.Fatal(err)
	}
}
