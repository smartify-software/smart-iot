package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/smartify-software/smart-iot/pb" // replace with the actual package
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GeoFenceClient struct {
	client pb.GeoFenceLabelingServerClient
}

func (gfc *GeoFenceClient) StreamFilteredLocations() (<-chan *pb.FencedLocation, error) {
	stream, err := gfc.client.StreamFencedLocations(context.Background(), &emptypb.Empty{})
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

func NewGeoFenceClient(client *grpc.ClientConn) *GeoFenceClient {
	return &GeoFenceClient{client: pb.NewGeoFenceLabelingServerClient(client)}
}

func (gfc *GeoFenceClient) AddPolygons(polygons []*pb.Polygon) error {
	//polygons := []*pb.Polygon{
	//	{
	//		Id:   "0",
	//		Name: "WashingtonDC_BoundingBox",
	//		Vertices: []*pb.Point{
	//			{Latitude: 38.8072, Longitude: -77.1369}, // Southwest corner
	//			{Latitude: 38.8072, Longitude: -76.9369}, // Southeast corner
	//			{Latitude: 39.0072, Longitude: -76.9369}, // Northeast corner
	//			{Latitude: 39.0072, Longitude: -77.1369}, // Northwest corner
	//		},
	//	},
	//	{
	//		Id:   "1",
	//		Name: "Polygon1",
	//		Vertices: []*pb.Point{
	//			{Latitude: 40.0, Longitude: 40.0},
	//			{Latitude: 42.0, Longitude: 40.0},
	//			{Latitude: 42.0, Longitude: 42.0},
	//			{Latitude: 40.0, Longitude: 42.0},
	//		},
	//	},
	//	{
	//		Id:   "2",
	//		Name: "Polygon2",
	//		Vertices: []*pb.Point{
	//			{Latitude: 50.0, Longitude: 50.0},
	//			{Latitude: 52.0, Longitude: 50.0},
	//			{Latitude: 52.0, Longitude: 52.0},
	//			{Latitude: 50.0, Longitude: 52.0},
	//		},
	//	},
	//}

	req := &pb.AddPolygonsRequest{Polygons: polygons}

	_, err := gfc.client.AddPolygonsForGeoFencing(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not add polygons: %v", err)
	}

	log.Println("Polygons added successfully.")
	return nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := &InMemoryGeoFenceRepository{Polygons: make(map[string]*pb.Polygon)}
	server := grpc.NewServer()
	geoFenceServer := NewGeoFenceServer(repo)

	pb.RegisterGeoFenceLabelingServerServer(server, geoFenceServer)

	fmt.Println("GeoFenceService running on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	go StartServer()

	time.Sleep(2 * time.Second)

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
	polygons := []*pb.Polygon{
		{
			Id:   "0",
			Name: "WashingtonDC_BoundingBox",
			Vertices: []*pb.Point{
				{Latitude: 38.8072, Longitude: -77.1369}, // Southwest corner
				{Latitude: 38.8072, Longitude: -76.9369}, // Southeast corner
				{Latitude: 39.0072, Longitude: -76.9369}, // Northeast corner
				{Latitude: 39.0072, Longitude: -77.1369}, // Northwest corner
			},
		},
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

	err = client.AddPolygons(polygons)
	if err != nil {
		log.Printf("Error while adding polygons: %v", err)
	}

	// Let's wait a bit before streaming (simulating some delay)
	time.Sleep(2 * time.Second)

	// Stream Filtered Locations
	go func() {
		_, err := client.StreamFilteredLocations()
		if err != nil {
			log.Printf("Error while streaming filtered locations: %v", err)
		}
	}()

	// Stop the client after some time
	time.Sleep(10 * time.Second)
}
