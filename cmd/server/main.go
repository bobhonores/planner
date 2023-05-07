package main

import (
	"fmt"

	"github.com/bobhonores/planner/internal/action"
	"github.com/bobhonores/planner/internal/comment"
	"github.com/bobhonores/planner/internal/db"
	"github.com/bobhonores/planner/internal/transport/grpc"
	transportHttp "github.com/bobhonores/planner/internal/transport/http"
)

func main() {
	fmt.Println("Starting Planner server")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}

func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	actService := action.New(db)

	grpcHandler := grpc.New(actService)
	if err := grpcHandler.Serve(); err != nil {
		return err
	}

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
