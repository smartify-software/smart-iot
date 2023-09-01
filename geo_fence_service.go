package main

import (
	"context"
	"fmt"
	"github.com/smartify-software/smart-iot/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"time"
)

type GeoFenceServer struct {
	pb.UnimplementedGeoFenceLabelingServerServer
	Repo GeoFenceRepository
}

func NewGeoFenceServer(
	repo GeoFenceRepository,
) *GeoFenceServer {
	return &GeoFenceServer{Repo: repo}
}

func (s *GeoFenceServer) AddPolygonsForGeoFencing(
	ctx context.Context,
	req *pb.AddPolygonsRequest,
) (*emptypb.Empty, error) {
	for _, polygon := range req.Polygons {
		if err := s.Repo.Store(polygon); err != nil {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *GeoFenceServer) StreamFencedLocations(
	_ *emptypb.Empty,
	stream pb.GeoFenceLabelingServer_StreamFencedLocationsServer,
) error {
	// receive locations from the client
	// and send fenced locations to the downstream client
	for location := range GetRandomLocation() {

		// Check if the received location is inside any of the polygons
		for _, polygon := range s.Repo.GetPolygons() {
			if IsPointInPolygon(location.Latitude, location.Longitude, polygon) {
				// Send FencedLocation to the downstream
				fencedLocation := &pb.FencedLocation{
					Latitude:  location.Latitude,
					Longitude: location.Longitude,
					PolygonId: polygon.Id,
				}
				if err := stream.Send(fencedLocation); err != nil {
					return fmt.Errorf("could not send fenced location: %v", err)
				}
			}
		}
	}
	return nil
}

func GetRandomLocation() chan *pb.Location {
	ch := make(chan *pb.Location)
	go func() {
		for {
			lat := randomNoise(80)  // Random latitude around San Francisco, for example
			long := randomNoise(80) // Random longitude around San Francisco, for example
			location := &pb.Location{Latitude: lat, Longitude: long}
			ch <- location
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

// IsPointInPolygon checks if a point (x, y) is inside a given polygon

// IsPointInPolygon checks if a point (x, y) is inside a given polygon
func IsPointInPolygon(x, y float64, polygon *pb.Polygon) bool {
	n := len(polygon.Vertices)
	if n < 3 {
		return false
	}

	inside := false
	for i, j := 0, n-1; i < n; j, i = i, i+1 {
		xi, yi := polygon.Vertices[i].Latitude, polygon.Vertices[i].Longitude
		xj, yj := polygon.Vertices[j].Latitude, polygon.Vertices[j].Longitude

		intersect := ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi)+xi)
		if intersect {
			inside = !inside
		}
	}
	return inside
}

func randomNoise(coord float64) float64 {
	return coord + rand.Float64()/100
}
