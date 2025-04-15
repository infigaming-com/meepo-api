// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: system/service/v1/system.proto

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

type Currency struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Symbol        string                 `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Icon          string                 `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	DecimalPlaces int32                  `protobuf:"varint,5,opt,name=decimal_places,json=decimalPlaces,proto3" json:"decimal_places,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Currency) Reset() {
	*x = Currency{}
	mi := &file_system_service_v1_system_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{0}
}

func (x *Currency) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Currency) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Currency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Currency) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Currency) GetDecimalPlaces() int32 {
	if x != nil {
		return x.DecimalPlaces
	}
	return 0
}

type AddCurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      *Currency              `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCurrencyRequest) Reset() {
	*x = AddCurrencyRequest{}
	mi := &file_system_service_v1_system_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCurrencyRequest) ProtoMessage() {}

func (x *AddCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCurrencyRequest.ProtoReflect.Descriptor instead.
func (*AddCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{1}
}

func (x *AddCurrencyRequest) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

type AddCurrencyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      *Currency              `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCurrencyResponse) Reset() {
	*x = AddCurrencyResponse{}
	mi := &file_system_service_v1_system_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCurrencyResponse) ProtoMessage() {}

func (x *AddCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCurrencyResponse.ProtoReflect.Descriptor instead.
func (*AddCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{2}
}

func (x *AddCurrencyResponse) GetCurrency() *Currency {
	if x != nil {
		return x.Currency
	}
	return nil
}

type GetCurrenciesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currency      string                 `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrenciesRequest) Reset() {
	*x = GetCurrenciesRequest{}
	mi := &file_system_service_v1_system_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesRequest) ProtoMessage() {}

func (x *GetCurrenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*GetCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrenciesRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currencies    *Currency              `protobuf:"bytes,1,opt,name=currencies,proto3" json:"currencies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	mi := &file_system_service_v1_system_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{4}
}

func (x *GetCurrenciesResponse) GetCurrencies() *Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

type ListCurrenciesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCurrenciesRequest) Reset() {
	*x = ListCurrenciesRequest{}
	mi := &file_system_service_v1_system_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrenciesRequest) ProtoMessage() {}

func (x *ListCurrenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*ListCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{5}
}

type ListCurrenciesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Currencies    []*Currency            `protobuf:"bytes,1,rep,name=currencies,proto3" json:"currencies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCurrenciesResponse) Reset() {
	*x = ListCurrenciesResponse{}
	mi := &file_system_service_v1_system_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrenciesResponse) ProtoMessage() {}

func (x *ListCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_service_v1_system_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*ListCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_system_service_v1_system_proto_rawDescGZIP(), []int{6}
}

func (x *ListCurrenciesResponse) GetCurrencies() []*Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

var File_system_service_v1_system_proto protoreflect.FileDescriptor

const file_system_service_v1_system_proto_rawDesc = "" +
	"\n" +
	"\x1esystem/service/v1/system.proto\x12\x11system.service.v1\x1a\x1cgoogle/api/annotations.proto\"\x8d\x01\n" +
	"\bCurrency\x12\x1a\n" +
	"\bcurrency\x18\x01 \x01(\tR\bcurrency\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x16\n" +
	"\x06symbol\x18\x03 \x01(\tR\x06symbol\x12\x12\n" +
	"\x04icon\x18\x04 \x01(\tR\x04icon\x12%\n" +
	"\x0edecimal_places\x18\x05 \x01(\x05R\rdecimalPlaces\"M\n" +
	"\x12AddCurrencyRequest\x127\n" +
	"\bcurrency\x18\x01 \x01(\v2\x1b.system.service.v1.CurrencyR\bcurrency\"N\n" +
	"\x13AddCurrencyResponse\x127\n" +
	"\bcurrency\x18\x01 \x01(\v2\x1b.system.service.v1.CurrencyR\bcurrency\"2\n" +
	"\x14GetCurrenciesRequest\x12\x1a\n" +
	"\bcurrency\x18\x01 \x01(\tR\bcurrency\"T\n" +
	"\x15GetCurrenciesResponse\x12;\n" +
	"\n" +
	"currencies\x18\x01 \x01(\v2\x1b.system.service.v1.CurrencyR\n" +
	"currencies\"\x17\n" +
	"\x15ListCurrenciesRequest\"U\n" +
	"\x16ListCurrenciesResponse\x12;\n" +
	"\n" +
	"currencies\x18\x01 \x03(\v2\x1b.system.service.v1.CurrencyR\n" +
	"currencies2\xa7\x03\n" +
	"\x06System\x12\x82\x01\n" +
	"\vAddCurrency\x12%.system.service.v1.AddCurrencyRequest\x1a&.system.service.v1.AddCurrencyResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v1/system/currencies/add\x12\x88\x01\n" +
	"\rGetCurrencies\x12'.system.service.v1.GetCurrenciesRequest\x1a(.system.service.v1.GetCurrenciesResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v1/system/currencies/get\x12\x8c\x01\n" +
	"\x0eListCurrencies\x12(.system.service.v1.ListCurrenciesRequest\x1a).system.service.v1.ListCurrenciesResponse\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/v1/system/currencies/listBO\n" +
	"\x11system.service.v1P\x01Z8github.com/infigaming-com/meepo-api/system/service/v1;v1b\x06proto3"

var (
	file_system_service_v1_system_proto_rawDescOnce sync.Once
	file_system_service_v1_system_proto_rawDescData []byte
)

func file_system_service_v1_system_proto_rawDescGZIP() []byte {
	file_system_service_v1_system_proto_rawDescOnce.Do(func() {
		file_system_service_v1_system_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_system_service_v1_system_proto_rawDesc), len(file_system_service_v1_system_proto_rawDesc)))
	})
	return file_system_service_v1_system_proto_rawDescData
}

var file_system_service_v1_system_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_system_service_v1_system_proto_goTypes = []any{
	(*Currency)(nil),               // 0: system.service.v1.Currency
	(*AddCurrencyRequest)(nil),     // 1: system.service.v1.AddCurrencyRequest
	(*AddCurrencyResponse)(nil),    // 2: system.service.v1.AddCurrencyResponse
	(*GetCurrenciesRequest)(nil),   // 3: system.service.v1.GetCurrenciesRequest
	(*GetCurrenciesResponse)(nil),  // 4: system.service.v1.GetCurrenciesResponse
	(*ListCurrenciesRequest)(nil),  // 5: system.service.v1.ListCurrenciesRequest
	(*ListCurrenciesResponse)(nil), // 6: system.service.v1.ListCurrenciesResponse
}
var file_system_service_v1_system_proto_depIdxs = []int32{
	0, // 0: system.service.v1.AddCurrencyRequest.currency:type_name -> system.service.v1.Currency
	0, // 1: system.service.v1.AddCurrencyResponse.currency:type_name -> system.service.v1.Currency
	0, // 2: system.service.v1.GetCurrenciesResponse.currencies:type_name -> system.service.v1.Currency
	0, // 3: system.service.v1.ListCurrenciesResponse.currencies:type_name -> system.service.v1.Currency
	1, // 4: system.service.v1.System.AddCurrency:input_type -> system.service.v1.AddCurrencyRequest
	3, // 5: system.service.v1.System.GetCurrencies:input_type -> system.service.v1.GetCurrenciesRequest
	5, // 6: system.service.v1.System.ListCurrencies:input_type -> system.service.v1.ListCurrenciesRequest
	2, // 7: system.service.v1.System.AddCurrency:output_type -> system.service.v1.AddCurrencyResponse
	4, // 8: system.service.v1.System.GetCurrencies:output_type -> system.service.v1.GetCurrenciesResponse
	6, // 9: system.service.v1.System.ListCurrencies:output_type -> system.service.v1.ListCurrenciesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_system_service_v1_system_proto_init() }
func file_system_service_v1_system_proto_init() {
	if File_system_service_v1_system_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_system_service_v1_system_proto_rawDesc), len(file_system_service_v1_system_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_service_v1_system_proto_goTypes,
		DependencyIndexes: file_system_service_v1_system_proto_depIdxs,
		MessageInfos:      file_system_service_v1_system_proto_msgTypes,
	}.Build()
	File_system_service_v1_system_proto = out.File
	file_system_service_v1_system_proto_goTypes = nil
	file_system_service_v1_system_proto_depIdxs = nil
}
