package main

import (
	"fmt"
	"log"
	"net"

	"github.com/smartify-software/smart-iot/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := &InMemoryGeoFenceRepository{Polygons: make(map[string]*pb.Polygon)}
	server := grpc.NewServer()
	geoFenceServer := &GeoFenceServer{Repo: repo}

	pb.RegisterGeoFenceServiceServer(server, geoFenceServer)

	fmt.Println("GeoFenceService running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
