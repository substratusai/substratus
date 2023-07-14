package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/substratusai/substratus/internal/gcpmanager"
	"github.com/substratusai/substratus/internal/sci"
	"google.golang.org/grpc"
)

func main() {
	// Create a storage client
	storageClient, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create storage client: %v", err)
	}

	s := grpc.NewServer()
	sci.RegisterControllerServer(s, &gcpmanager.Server{
		StorageClient: storageClient,
	})

	port := 10443
	fmt.Printf("gcpmanager server listening on port %v...", port)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
