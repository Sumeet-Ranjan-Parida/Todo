package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"std/Proj/proto"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Sumeet-Ranjan-Parida/Todo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Failed to listen on port 4040: %v", err)
	}

	gs := grpc.NewServer()
	proto.RegisterContactServer(gs, &server{})
	reflection.Register(gs)

	if err := gs.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server on port 4040: %v", err)
	}
}

func (s *server) GetCred(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	db := dbConn()

	defer db.Close()

	username, password := request.GetUsername(), request.GetPassword()

	insert, err := db.Prepare("INSERT INTO accounts(username, password) VALUES (?,?)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(username, password)

}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:sumeet@tcp(127.0.0.1:3306)/todoproj")
	if err != nil {
		panic(err.Error())
	}

	return db
}
