// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: backoffice/service/v1/error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_UNSPECIFIED                       ErrorReason = 0
	ErrorReason_CALL_WALLET_SERVICE_FAILED        ErrorReason = 60000
	ErrorReason_OPERATOR_IDS_NOT_FOUND_IN_CONTEXT ErrorReason = 60001
	ErrorReason_USER_INFO_NOT_FOUND_IN_CONTEXT    ErrorReason = 60002
	ErrorReason_REPORT_TIME_RANGE_ERROR           ErrorReason = 60003
	ErrorReason_REPORT_GET_DATA_ERROR             ErrorReason = 60004
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "UNSPECIFIED",
		60000: "CALL_WALLET_SERVICE_FAILED",
		60001: "OPERATOR_IDS_NOT_FOUND_IN_CONTEXT",
		60002: "USER_INFO_NOT_FOUND_IN_CONTEXT",
		60003: "REPORT_TIME_RANGE_ERROR",
		60004: "REPORT_GET_DATA_ERROR",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":                       0,
		"CALL_WALLET_SERVICE_FAILED":        60000,
		"OPERATOR_IDS_NOT_FOUND_IN_CONTEXT": 60001,
		"USER_INFO_NOT_FOUND_IN_CONTEXT":    60002,
		"REPORT_TIME_RANGE_ERROR":           60003,
		"REPORT_GET_DATA_ERROR":             60004,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_backoffice_service_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_backoffice_service_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_backoffice_service_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_backoffice_service_v1_error_reason_proto protoreflect.FileDescriptor

const file_backoffice_service_v1_error_reason_proto_rawDesc = "" +
	"\n" +
	"(backoffice/service/v1/error_reason.proto\x12\x19api.backoffice.service.v1\x1a\x13errors/errors.proto*\xd1\x01\n" +
	"\vErrorReason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12 \n" +
	"\x1aCALL_WALLET_SERVICE_FAILED\x10\xe0\xd4\x03\x12'\n" +
	"!OPERATOR_IDS_NOT_FOUND_IN_CONTEXT\x10\xe1\xd4\x03\x12$\n" +
	"\x1eUSER_INFO_NOT_FOUND_IN_CONTEXT\x10\xe2\xd4\x03\x12\x1d\n" +
	"\x17REPORT_TIME_RANGE_ERROR\x10\xe3\xd4\x03\x12\x1b\n" +
	"\x15REPORT_GET_DATA_ERROR\x10\xe4\xd4\x03\x1a\x04\xa0E\xf4\x03B[\n" +
	"\x19api.backoffice.service.v1P\x01Z<github.com/infigaming-com/meepo-api/backoffice/service/v1;v1b\x06proto3"

var (
	file_backoffice_service_v1_error_reason_proto_rawDescOnce sync.Once
	file_backoffice_service_v1_error_reason_proto_rawDescData []byte
)

func file_backoffice_service_v1_error_reason_proto_rawDescGZIP() []byte {
	file_backoffice_service_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_backoffice_service_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_backoffice_service_v1_error_reason_proto_rawDesc), len(file_backoffice_service_v1_error_reason_proto_rawDesc)))
	})
	return file_backoffice_service_v1_error_reason_proto_rawDescData
}

var file_backoffice_service_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_backoffice_service_v1_error_reason_proto_goTypes = []any{
	(ErrorReason)(0), // 0: api.backoffice.service.v1.ErrorReason
}
var file_backoffice_service_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backoffice_service_v1_error_reason_proto_init() }
func file_backoffice_service_v1_error_reason_proto_init() {
	if File_backoffice_service_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_backoffice_service_v1_error_reason_proto_rawDesc), len(file_backoffice_service_v1_error_reason_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backoffice_service_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_backoffice_service_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_backoffice_service_v1_error_reason_proto_enumTypes,
	}.Build()
	File_backoffice_service_v1_error_reason_proto = out.File
	file_backoffice_service_v1_error_reason_proto_goTypes = nil
	file_backoffice_service_v1_error_reason_proto_depIdxs = nil
}
