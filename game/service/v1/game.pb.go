// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: game/service/v1/game.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOperatorRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	AggregatorOperatorId string                 `protobuf:"bytes,1,opt,name=aggregator_operator_id,json=aggregatorOperatorId,proto3" json:"aggregator_operator_id,omitempty"`
	OperatorId           string                 `protobuf:"bytes,2,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	ApiKey               string                 `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string                 `protobuf:"bytes,4,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Enabled              bool                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *CreateOperatorRequest) Reset() {
	*x = CreateOperatorRequest{}
	mi := &file_game_service_v1_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOperatorRequest) ProtoMessage() {}

func (x *CreateOperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOperatorRequest.ProtoReflect.Descriptor instead.
func (*CreateOperatorRequest) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOperatorRequest) GetAggregatorOperatorId() string {
	if x != nil {
		return x.AggregatorOperatorId
	}
	return ""
}

func (x *CreateOperatorRequest) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *CreateOperatorRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *CreateOperatorRequest) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *CreateOperatorRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type CreateOperatorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOperatorResponse) Reset() {
	*x = CreateOperatorResponse{}
	mi := &file_game_service_v1_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOperatorResponse) ProtoMessage() {}

func (x *CreateOperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOperatorResponse.ProtoReflect.Descriptor instead.
func (*CreateOperatorResponse) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOperatorResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOperatorRequest struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	OperatorId           string                 `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	AggregatorOperatorId string                 `protobuf:"bytes,2,opt,name=aggregator_operator_id,json=aggregatorOperatorId,proto3" json:"aggregator_operator_id,omitempty"`
	ApiKey               string                 `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string                 `protobuf:"bytes,4,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Enabled              bool                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *UpdateOperatorRequest) Reset() {
	*x = UpdateOperatorRequest{}
	mi := &file_game_service_v1_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOperatorRequest) ProtoMessage() {}

func (x *UpdateOperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOperatorRequest.ProtoReflect.Descriptor instead.
func (*UpdateOperatorRequest) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOperatorRequest) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *UpdateOperatorRequest) GetAggregatorOperatorId() string {
	if x != nil {
		return x.AggregatorOperatorId
	}
	return ""
}

func (x *UpdateOperatorRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *UpdateOperatorRequest) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *UpdateOperatorRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateOperatorResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	AggregatorOperatorId string                 `protobuf:"bytes,1,opt,name=aggregator_operator_id,json=aggregatorOperatorId,proto3" json:"aggregator_operator_id,omitempty"`
	OperatorId           string                 `protobuf:"bytes,2,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	ApiKey               string                 `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string                 `protobuf:"bytes,4,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Enabled              bool                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *UpdateOperatorResponse) Reset() {
	*x = UpdateOperatorResponse{}
	mi := &file_game_service_v1_game_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOperatorResponse) ProtoMessage() {}

func (x *UpdateOperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOperatorResponse.ProtoReflect.Descriptor instead.
func (*UpdateOperatorResponse) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOperatorResponse) GetAggregatorOperatorId() string {
	if x != nil {
		return x.AggregatorOperatorId
	}
	return ""
}

func (x *UpdateOperatorResponse) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *UpdateOperatorResponse) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *UpdateOperatorResponse) GetApiSecret() string {
	if x != nil {
		return x.ApiSecret
	}
	return ""
}

func (x *UpdateOperatorResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type DeleteOperatorRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OperatorId    string                 `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOperatorRequest) Reset() {
	*x = DeleteOperatorRequest{}
	mi := &file_game_service_v1_game_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOperatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperatorRequest) ProtoMessage() {}

func (x *DeleteOperatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperatorRequest.ProtoReflect.Descriptor instead.
func (*DeleteOperatorRequest) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOperatorRequest) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

type DeleteOperatorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteOperatorResponse) Reset() {
	*x = DeleteOperatorResponse{}
	mi := &file_game_service_v1_game_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteOperatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperatorResponse) ProtoMessage() {}

func (x *DeleteOperatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_service_v1_game_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperatorResponse.ProtoReflect.Descriptor instead.
func (*DeleteOperatorResponse) Descriptor() ([]byte, []int) {
	return file_game_service_v1_game_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOperatorResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_game_service_v1_game_proto protoreflect.FileDescriptor

const file_game_service_v1_game_proto_rawDesc = "" +
	"\n" +
	"\x1agame/service/v1/game.proto\x12\x0fgame.service.v1\x1a\x1cgoogle/api/annotations.proto\"\xc0\x01\n" +
	"\x15CreateOperatorRequest\x124\n" +
	"\x16aggregator_operator_id\x18\x01 \x01(\tR\x14aggregatorOperatorId\x12\x1f\n" +
	"\voperator_id\x18\x02 \x01(\tR\n" +
	"operatorId\x12\x17\n" +
	"\aapi_key\x18\x03 \x01(\tR\x06apiKey\x12\x1d\n" +
	"\n" +
	"api_secret\x18\x04 \x01(\tR\tapiSecret\x12\x18\n" +
	"\aenabled\x18\x05 \x01(\bR\aenabled\"0\n" +
	"\x16CreateOperatorResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"\xc0\x01\n" +
	"\x15UpdateOperatorRequest\x12\x1f\n" +
	"\voperator_id\x18\x01 \x01(\tR\n" +
	"operatorId\x124\n" +
	"\x16aggregator_operator_id\x18\x02 \x01(\tR\x14aggregatorOperatorId\x12\x17\n" +
	"\aapi_key\x18\x03 \x01(\tR\x06apiKey\x12\x1d\n" +
	"\n" +
	"api_secret\x18\x04 \x01(\tR\tapiSecret\x12\x18\n" +
	"\aenabled\x18\x05 \x01(\bR\aenabled\"\xc1\x01\n" +
	"\x16UpdateOperatorResponse\x124\n" +
	"\x16aggregator_operator_id\x18\x01 \x01(\tR\x14aggregatorOperatorId\x12\x1f\n" +
	"\voperator_id\x18\x02 \x01(\tR\n" +
	"operatorId\x12\x17\n" +
	"\aapi_key\x18\x03 \x01(\tR\x06apiKey\x12\x1d\n" +
	"\n" +
	"api_secret\x18\x04 \x01(\tR\tapiSecret\x12\x18\n" +
	"\aenabled\x18\x05 \x01(\bR\aenabled\"8\n" +
	"\x15DeleteOperatorRequest\x12\x1f\n" +
	"\voperator_id\x18\x01 \x01(\tR\n" +
	"operatorId\"0\n" +
	"\x16DeleteOperatorResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xb5\x02\n" +
	"\x04Game\x12c\n" +
	"\x0eCreateOperator\x12&.game.service.v1.CreateOperatorRequest\x1a'.game.service.v1.CreateOperatorResponse\"\x00\x12c\n" +
	"\x0eUpdateOperator\x12&.game.service.v1.UpdateOperatorRequest\x1a'.game.service.v1.UpdateOperatorResponse\"\x00\x12c\n" +
	"\x0eDeleteOperator\x12&.game.service.v1.DeleteOperatorRequest\x1a'.game.service.v1.DeleteOperatorResponse\"\x00BK\n" +
	"\x0fgame.service.v1P\x01Z6github.com/infigaming-com/meepo-api/game/service/v1;v1b\x06proto3"

var (
	file_game_service_v1_game_proto_rawDescOnce sync.Once
	file_game_service_v1_game_proto_rawDescData []byte
)

func file_game_service_v1_game_proto_rawDescGZIP() []byte {
	file_game_service_v1_game_proto_rawDescOnce.Do(func() {
		file_game_service_v1_game_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_game_service_v1_game_proto_rawDesc), len(file_game_service_v1_game_proto_rawDesc)))
	})
	return file_game_service_v1_game_proto_rawDescData
}

var file_game_service_v1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_game_service_v1_game_proto_goTypes = []any{
	(*CreateOperatorRequest)(nil),  // 0: game.service.v1.CreateOperatorRequest
	(*CreateOperatorResponse)(nil), // 1: game.service.v1.CreateOperatorResponse
	(*UpdateOperatorRequest)(nil),  // 2: game.service.v1.UpdateOperatorRequest
	(*UpdateOperatorResponse)(nil), // 3: game.service.v1.UpdateOperatorResponse
	(*DeleteOperatorRequest)(nil),  // 4: game.service.v1.DeleteOperatorRequest
	(*DeleteOperatorResponse)(nil), // 5: game.service.v1.DeleteOperatorResponse
}
var file_game_service_v1_game_proto_depIdxs = []int32{
	0, // 0: game.service.v1.Game.CreateOperator:input_type -> game.service.v1.CreateOperatorRequest
	2, // 1: game.service.v1.Game.UpdateOperator:input_type -> game.service.v1.UpdateOperatorRequest
	4, // 2: game.service.v1.Game.DeleteOperator:input_type -> game.service.v1.DeleteOperatorRequest
	1, // 3: game.service.v1.Game.CreateOperator:output_type -> game.service.v1.CreateOperatorResponse
	3, // 4: game.service.v1.Game.UpdateOperator:output_type -> game.service.v1.UpdateOperatorResponse
	5, // 5: game.service.v1.Game.DeleteOperator:output_type -> game.service.v1.DeleteOperatorResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_service_v1_game_proto_init() }
func file_game_service_v1_game_proto_init() {
	if File_game_service_v1_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_game_service_v1_game_proto_rawDesc), len(file_game_service_v1_game_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_service_v1_game_proto_goTypes,
		DependencyIndexes: file_game_service_v1_game_proto_depIdxs,
		MessageInfos:      file_game_service_v1_game_proto_msgTypes,
	}.Build()
	File_game_service_v1_game_proto = out.File
	file_game_service_v1_game_proto_goTypes = nil
	file_game_service_v1_game_proto_depIdxs = nil
}
