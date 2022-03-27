// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: portsvc/proto/ports.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SavePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Port *Port  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *SavePortRequest) Reset() {
	*x = SavePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portsvc_proto_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePortRequest) ProtoMessage() {}

func (x *SavePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portsvc_proto_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePortRequest.ProtoReflect.Descriptor instead.
func (*SavePortRequest) Descriptor() ([]byte, []int) {
	return file_portsvc_proto_ports_proto_rawDescGZIP(), []int{0}
}

func (x *SavePortRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SavePortRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type SavePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SavePortResponse) Reset() {
	*x = SavePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portsvc_proto_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePortResponse) ProtoMessage() {}

func (x *SavePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portsvc_proto_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePortResponse.ProtoReflect.Descriptor instead.
func (*SavePortResponse) Descriptor() ([]byte, []int) {
	return file_portsvc_proto_ports_proto_rawDescGZIP(), []int{1}
}

func (x *SavePortResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPortsRequest) Reset() {
	*x = GetPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portsvc_proto_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortsRequest) ProtoMessage() {}

func (x *GetPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portsvc_proto_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortsRequest.ProtoReflect.Descriptor instead.
func (*GetPortsRequest) Descriptor() ([]byte, []int) {
	return file_portsvc_proto_ports_proto_rawDescGZIP(), []int{2}
}

type PortList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports map[string]*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PortList) Reset() {
	*x = PortList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portsvc_proto_ports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortList) ProtoMessage() {}

func (x *PortList) ProtoReflect() protoreflect.Message {
	mi := &file_portsvc_proto_ports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortList.ProtoReflect.Descriptor instead.
func (*PortList) Descriptor() ([]byte, []int) {
	return file_portsvc_proto_ports_proto_rawDescGZIP(), []int{3}
}

func (x *PortList) GetPorts() map[string]*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Timezone    string    `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,8,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,9,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portsvc_proto_ports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_portsvc_proto_ports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_portsvc_proto_ports_proto_rawDescGZIP(), []int{4}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_portsvc_proto_ports_proto protoreflect.FileDescriptor

var file_portsvc_proto_ports_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x50, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x45, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xe2, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52,
	0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x79, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x39, 0x0a,
	0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x76, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portsvc_proto_ports_proto_rawDescOnce sync.Once
	file_portsvc_proto_ports_proto_rawDescData = file_portsvc_proto_ports_proto_rawDesc
)

func file_portsvc_proto_ports_proto_rawDescGZIP() []byte {
	file_portsvc_proto_ports_proto_rawDescOnce.Do(func() {
		file_portsvc_proto_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_portsvc_proto_ports_proto_rawDescData)
	})
	return file_portsvc_proto_ports_proto_rawDescData
}

var file_portsvc_proto_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_portsvc_proto_ports_proto_goTypes = []interface{}{
	(*SavePortRequest)(nil),  // 0: proto.SavePortRequest
	(*SavePortResponse)(nil), // 1: proto.SavePortResponse
	(*GetPortsRequest)(nil),  // 2: proto.GetPortsRequest
	(*PortList)(nil),         // 3: proto.PortList
	(*Port)(nil),             // 4: proto.Port
	nil,                      // 5: proto.PortList.PortsEntry
}
var file_portsvc_proto_ports_proto_depIdxs = []int32{
	4, // 0: proto.SavePortRequest.port:type_name -> proto.Port
	5, // 1: proto.PortList.ports:type_name -> proto.PortList.PortsEntry
	4, // 2: proto.PortList.PortsEntry.value:type_name -> proto.Port
	0, // 3: proto.Ports.Save:input_type -> proto.SavePortRequest
	2, // 4: proto.Ports.GetPorts:input_type -> proto.GetPortsRequest
	1, // 5: proto.Ports.Save:output_type -> proto.SavePortResponse
	3, // 6: proto.Ports.GetPorts:output_type -> proto.PortList
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_portsvc_proto_ports_proto_init() }
func file_portsvc_proto_ports_proto_init() {
	if File_portsvc_proto_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portsvc_proto_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePortRequest); i {
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
		file_portsvc_proto_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePortResponse); i {
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
		file_portsvc_proto_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortsRequest); i {
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
		file_portsvc_proto_ports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortList); i {
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
		file_portsvc_proto_ports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_portsvc_proto_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portsvc_proto_ports_proto_goTypes,
		DependencyIndexes: file_portsvc_proto_ports_proto_depIdxs,
		MessageInfos:      file_portsvc_proto_ports_proto_msgTypes,
	}.Build()
	File_portsvc_proto_ports_proto = out.File
	file_portsvc_proto_ports_proto_rawDesc = nil
	file_portsvc_proto_ports_proto_goTypes = nil
	file_portsvc_proto_ports_proto_depIdxs = nil
}
