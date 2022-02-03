package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("idk")
	r := gin.Default()
	r.POST("/csv", readCSV)
	r.Run(os.Getenv("PORT"))
}
