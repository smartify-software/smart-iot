// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: geofence.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A single point on the map represented by latitude and longitude.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// A polygon represented as a series of points.
type Polygon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // Unique identifier for this polygon
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`         // Friendly name for this polygon
	Vertices []*Point `protobuf:"bytes,3,rep,name=vertices,proto3" json:"vertices,omitempty"` // Vertices of the polygon
}

func (x *Polygon) Reset() {
	*x = Polygon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Polygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Polygon) ProtoMessage() {}

func (x *Polygon) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Polygon.ProtoReflect.Descriptor instead.
func (*Polygon) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{1}
}

func (x *Polygon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Polygon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Polygon) GetVertices() []*Point {
	if x != nil {
		return x.Vertices
	}
	return nil
}

// Request message to add one or more polygons for filtering.
type AddPolygonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polygons []*Polygon `protobuf:"bytes,1,rep,name=polygons,proto3" json:"polygons,omitempty"` // The polygons to be added
}

func (x *AddPolygonsRequest) Reset() {
	*x = AddPolygonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPolygonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPolygonsRequest) ProtoMessage() {}

func (x *AddPolygonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPolygonsRequest.ProtoReflect.Descriptor instead.
func (*AddPolygonsRequest) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{2}
}

func (x *AddPolygonsRequest) GetPolygons() []*Polygon {
	if x != nil {
		return x.Polygons
	}
	return nil
}

// A single location represented by latitude and longitude, along with the ID of the polygon it belongs to.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	PolygonId string  `protobuf:"bytes,3,opt,name=polygon_id,json=polygonId,proto3" json:"polygon_id,omitempty"` // The ID of the polygon this location falls within.
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetPolygonId() string {
	if x != nil {
		return x.PolygonId
	}
	return ""
}

// The streaming response containing locations within polygons.
type FilteredLocations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"` // The filtered locations within the polygons.
}

func (x *FilteredLocations) Reset() {
	*x = FilteredLocations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geofence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilteredLocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteredLocations) ProtoMessage() {}

func (x *FilteredLocations) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteredLocations.ProtoReflect.Descriptor instead.
func (*FilteredLocations) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{4}
}

func (x *FilteredLocations) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

var File_geofence_proto protoreflect.FileDescriptor

var file_geofence_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65,
	0x6e, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x22, 0x63, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x11, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x39, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x67,
	0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xba, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x6f, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x59, 0x0a,
	0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x67, 0x65, 0x6f, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geofence_proto_rawDescOnce sync.Once
	file_geofence_proto_rawDescData = file_geofence_proto_rawDesc
)

func file_geofence_proto_rawDescGZIP() []byte {
	file_geofence_proto_rawDescOnce.Do(func() {
		file_geofence_proto_rawDescData = protoimpl.X.CompressGZIP(file_geofence_proto_rawDescData)
	})
	return file_geofence_proto_rawDescData
}

var file_geofence_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_geofence_proto_goTypes = []interface{}{
	(*Point)(nil),              // 0: smartify.geofence.Point
	(*Polygon)(nil),            // 1: smartify.geofence.Polygon
	(*AddPolygonsRequest)(nil), // 2: smartify.geofence.AddPolygonsRequest
	(*Location)(nil),           // 3: smartify.geofence.Location
	(*FilteredLocations)(nil),  // 4: smartify.geofence.FilteredLocations
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_geofence_proto_depIdxs = []int32{
	0, // 0: smartify.geofence.Polygon.vertices:type_name -> smartify.geofence.Point
	1, // 1: smartify.geofence.AddPolygonsRequest.polygons:type_name -> smartify.geofence.Polygon
	3, // 2: smartify.geofence.FilteredLocations.locations:type_name -> smartify.geofence.Location
	2, // 3: smartify.geofence.GeoFenceService.AddPolygons:input_type -> smartify.geofence.AddPolygonsRequest
	5, // 4: smartify.geofence.GeoFenceService.StreamFilteredLocations:input_type -> google.protobuf.Empty
	5, // 5: smartify.geofence.GeoFenceService.AddPolygons:output_type -> google.protobuf.Empty
	4, // 6: smartify.geofence.GeoFenceService.StreamFilteredLocations:output_type -> smartify.geofence.FilteredLocations
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_geofence_proto_init() }
func file_geofence_proto_init() {
	if File_geofence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geofence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_geofence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Polygon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_geofence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPolygonsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_geofence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_geofence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilteredLocations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geofence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geofence_proto_goTypes,
		DependencyIndexes: file_geofence_proto_depIdxs,
		MessageInfos:      file_geofence_proto_msgTypes,
	}.Build()
	File_geofence_proto = out.File
	file_geofence_proto_rawDesc = nil
	file_geofence_proto_goTypes = nil
	file_geofence_proto_depIdxs = nil
}