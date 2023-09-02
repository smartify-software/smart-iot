package main

import "github.com/smartify-software/smart-iot/pb"

type GeoFenceClientInterface interface {
	AddPolygons(polygons []*pb.Polygon) error
	StreamFilteredLocations() (<-chan *pb.FencedLocation, error)
}
