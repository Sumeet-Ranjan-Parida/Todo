package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect: %s", err)
	}

	defer conn.Close()

	//client := proto.NewContactClient(conn)

	r := gin.Default()
	r.GET("/health", health)

	r.Run()
}

func health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"data": "Testing API", "alive": true})
}
