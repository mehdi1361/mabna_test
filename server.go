package main

import (
	"mabna_test/router"
	"github.com/gin-contrib/cors"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping/", router.PingHandler)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	port := os.Getenv("PORT")
	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
