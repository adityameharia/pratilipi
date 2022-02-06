package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := "http://localhost:8000/find/" + c.Param("userid")
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(502, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		if resp.Status != "200 OK" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Unable to validate user",
			})
			return
		}
		defer resp.Body.Close()
		var user UserApiResponse
		json.NewDecoder(resp.Body).Decode(&user)
		c.Set("user", user)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
