package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/matsunaga-project-2022/shanks/internal/controllers"
	"github.com/matsunaga-project-2022/shanks/internal/domain/repository"
	"github.com/matsunaga-project-2022/shanks/internal/domain/service"
	"github.com/matsunaga-project-2022/shanks/internal/pkg/grpc"
	"github.com/matsunaga-project-2022/shanks/internal/proto"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	run(ctx)
}

func run(ctx context.Context) {

	// Repository 組み立て
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/shanks?parseTime=true")
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

	stopServer, err := grpc.StartGRPCServer(ctx, server, 8080)
	if err != nil {
		log.Println(err)
		ctx.Done()
	}

	<-ctx.Done()
	stopServer()
}
