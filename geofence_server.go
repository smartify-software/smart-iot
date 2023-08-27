package main

import (
	"context"
	"github.com/smartify-software/smart-iot/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GeoFenceServer struct {
	pb.UnimplementedGeoFenceServiceServer
	Repo GeoFenceRepository
}

func (s *GeoFenceServer) AddPolygons(ctx context.Context, req *pb.AddPolygonsRequest) (*emptypb.Empty, error) {
	for _, polygon := range req.Polygons {
		if err := s.Repo.AddPolygon(polygon); err != nil {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *GeoFenceServer) StreamFilteredLocations(_ *emptypb.Empty, stream pb.GeoFenceService_StreamFilteredLocationsServer) error {
	filteredLocations, err := s.Repo.StreamFilteredLocations()
	if err != nil {
		return err
	}

	for _, locations := range filteredLocations {
		if err := stream.Send(locations); err != nil {
			return err
		}
	}

	return nil
}
