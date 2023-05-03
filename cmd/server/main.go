package main

import (
	"fmt"

	"github.com/bobhonores/planner/internal/comment"
	"github.com/bobhonores/planner/internal/db"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
