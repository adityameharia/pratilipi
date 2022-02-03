package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func readCSV(c *gin.Context) {
	records := readCsvFile("./test.csv")
	fmt.Println(records)
}
