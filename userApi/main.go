package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"time"

	"github.com/adityameharia/pratilipi/userApi/docs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	docs.SwaggerInfo.Title = "UserAPI"
	docs.SwaggerInfo.Description = "This server responds to the userApi requests"
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOps := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("successfully connected")
	}
	collection = client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("COLLECTION"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.Use(CORSMiddleware())
	r.POST("/signup", Signup)
	r.GET("/like/:cmd/:userId/:bookId", Like)
	r.GET("/find/:userId", FindUserRoute)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
