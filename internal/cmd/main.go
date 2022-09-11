package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/matsunaga-project-2022/shanks/internal/controllers"
	"github.com/matsunaga-project-2022/shanks/internal/domain/repository"
	"github.com/matsunaga-project-2022/shanks/internal/domain/service"
	"github.com/matsunaga-project-2022/shanks/internal/pkg/grpc"
	"github.com/matsunaga-project-2022/shanks/internal/proto"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	run(ctx)
}

func run(ctx context.Context) {

	// database
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	host := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// server
	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortStr)

	// Repository 組み立て
	db, err := sqlx.Connect("mysql", host)
	if err != nil {
		log.Println(err)
	}
	repo := repository.NewWordRepositoryMysqlImp(db)

	// Word Service 組み立て
	wordService := service.NewWord(repo)
	// Word Controller 組み立て
	wordController := controllers.NewWordController(wordService)

	server := grpc.NewGRPCServer()

	// Word proto 組み立て
	proto.RegisterWordServiceServer(server, wordController)

	stopServer, err := grpc.StartGRPCServer(ctx, server, uint(serverPort))
	if err != nil {
		log.Println(err)
		ctx.Done()
	}

	<-ctx.Done()
	stopServer()
}
