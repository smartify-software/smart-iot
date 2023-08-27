package main

import "github.com/smartify-software/smart-iot/pb"

type GeoFenceRepository interface {
	AddPolygon(polygon *pb.Polygon) error
	StreamFilteredLocations(locations []*pb.Point) ([]*pb.FilteredLocations, error)
}

type InMemoryGeoFenceRepository struct {
	Polygons map[string]*pb.Polygon
}

func (repo *InMemoryGeoFenceRepository) AddPolygon(polygon *pb.Polygon) error {
	repo.Polygons[polygon.Id] = polygon
	return nil
}

func (repo *InMemoryGeoFenceRepository) StreamFilteredLocations(locations []*pb.Point) ([]*pb.FilteredLocations, error) {
	var filtered []*pb.FilteredLocations

	for polygonID, polygon := range repo.Polygons {
		var locsInPolygon []*pb.Location

		for _, location := range locations {
			if pointInPolygon(location, polygon) {
				locsInPolygon = append(locsInPolygon, &pb.Location{
					Latitude:  location.Latitude,
					Longitude: location.Longitude,
					PolygonId: polygonID,
				})
			}
		}

		if len(locsInPolygon) > 0 {
			filtered = append(filtered, &pb.FilteredLocations{
				Locations: locsInPolygon,
			})
		}
	}

	return filtered, nil
}

func pointInPolygon(point *pb.Point, polygon *pb.Polygon) bool {
	// Initialize count and previous vertex
	count, prev := 0, polygon.Vertices[len(polygon.Vertices)-1]

	for _, vertex := range polygon.Vertices {
		if ((vertex.Latitude > point.Latitude) != (prev.Latitude > point.Latitude)) &&
			(point.Longitude < (prev.Longitude-vertex.Longitude)*(point.Latitude-vertex.Latitude)/(prev.Latitude-vertex.Latitude)+vertex.Longitude) {
			count++
		}
		prev = vertex
	}

	return count%2 == 1
}
