package main

import (
	"context"
	"log"
	"time"

	"github.com/smartify-software/smart-iot/pb" // replace with the actual package
	"google.golang.org/grpc"
)

type GeoFenceClient struct {
	client pb.GeoFenceServiceClient
}

func NewGeoFenceClient(cc *grpc.ClientConn) *GeoFenceClient {
	return &GeoFenceClient{
		client: pb.NewGeoFenceServiceClient(cc),
	}
}

func (gfc *GeoFenceClient) AddPolygons() {
	polygons := []*pb.Polygon{
		{
			Id:   "1",
			Name: "Polygon1",
			Vertices: []*pb.Point{
				{Latitude: 40.0, Longitude: 40.0},
				{Latitude: 42.0, Longitude: 40.0},
				{Latitude: 42.0, Longitude: 42.0},
				{Latitude: 40.0, Longitude: 42.0},
			},
		},
		{
			Id:   "2",
			Name: "Polygon2",
			Vertices: []*pb.Point{
				{Latitude: 50.0, Longitude: 50.0},
				{Latitude: 52.0, Longitude: 50.0},
				{Latitude: 52.0, Longitude: 52.0},
				{Latitude: 50.0, Longitude: 52.0},
			},
		},
	}

	req := &pb.AddPolygonsRequest{Polygons: polygons}

	_, err := gfc.client.AddPolygons(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add polygons: %v", err)
	}

	log.Println("Polygons added successfully.")
}

func (gfc *GeoFenceClient) StreamFilteredLocations() {
	stream, err := gfc.client.StreamFilteredLocations(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	for {
		filteredLocations, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving from stream: %v", err)
		}

		log.Printf("Received Filtered Locations: %v", filteredLocations)
	}
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}(cc)

	client := NewGeoFenceClient(cc)

	// Add Polygons
	client.AddPolygons()

	// Let's wait a bit before streaming (simulating some delay)
	time.Sleep(2 * time.Second)

	// Stream Filtered Locations
	go client.StreamFilteredLocations()

	// Stop the client after some time
	time.Sleep(10 * time.Second)
}
