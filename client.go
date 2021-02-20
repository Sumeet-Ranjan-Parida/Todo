package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type signupinput struct {
	username string `json:"username" binding:"required"`
	password string `json:"password" binding:"required"`
}

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
	r.POST("/signup", func(ctx *gin.Context) {
		var input signupinput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"username": input.username, "password": input.password})
	})

	r.Run()
}

func health(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{"status": 200, "data": "Testing API", "alive": true})
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/todoproj")
	if err != nil {
		panic(err.Error())
	}

	return db
}
