//go:build e2e

package tests

import (
	"log"

	api "github.com/bobhonores/planner/api/v1"
	"google.golang.org/grpc"
)

func GetClient() api.ActionServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	actionClient := api.NewActionServiceClient(conn)
	return actionClient
}
