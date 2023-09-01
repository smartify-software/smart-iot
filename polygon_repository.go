package main

import "github.com/smartify-software/smart-iot/pb"

type GeoFenceRepository interface {
	Store(polygon *pb.Polygon) error
	GetPolygons() []*pb.Polygon
}

type InMemoryGeoFenceRepository struct {
	Polygons map[string]*pb.Polygon
}

func (repo *InMemoryGeoFenceRepository) Store(polygon *pb.Polygon) error {
	repo.Polygons[polygon.Id] = polygon
	return nil
}

func (repo *InMemoryGeoFenceRepository) GetPolygons() []*pb.Polygon {
	polygons := make([]*pb.Polygon, 0)
	for _, polygon := range repo.Polygons {
		polygons = append(polygons, polygon)
	}
	return polygons
}
