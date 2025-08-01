// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: payment/service/v1/error_reason.proto

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
	ErrorReason_UNSPECIFIED                      ErrorReason = 0
	ErrorReason_GET_PAYMENT_METHOD_LIST_FAILED   ErrorReason = 50001
	ErrorReason_CREATE_PAYMENT_CHANNEL_FAILED    ErrorReason = 50002
	ErrorReason_GET_PAYMENT_CHANNEL_PAGE_FAILED  ErrorReason = 50003
	ErrorReason_INITIATE_DEPOSIT_FAILED          ErrorReason = 50004
	ErrorReason_INITIATE_WITHDRAW_FAILED         ErrorReason = 50005
	ErrorReason_DEPOSIT_CALLBACK_FAILED          ErrorReason = 50006
	ErrorReason_WITHDRAW_CALLBACK_FAILED         ErrorReason = 50007
	ErrorReason_GET_TRANSACTION_PAGE_FAILED      ErrorReason = 50008
	ErrorReason_GET_TRANSACTION_DETAIL_FAILED    ErrorReason = 50009
	ErrorReason_GET_CHANNEL_LIST_FAILED          ErrorReason = 50010
	ErrorReason_GET_ADDRESS_FAILED               ErrorReason = 50011
	ErrorReason_OPERATOR_ID_NOT_FOUND_IN_CONTEXT ErrorReason = 50012
	ErrorReason_UPDATE_PAYMENT_METHOD_FAILED     ErrorReason = 50013
	ErrorReason_UPDATE_PAYMENT_CHANNEL_FAILED    ErrorReason = 50014
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "UNSPECIFIED",
		50001: "GET_PAYMENT_METHOD_LIST_FAILED",
		50002: "CREATE_PAYMENT_CHANNEL_FAILED",
		50003: "GET_PAYMENT_CHANNEL_PAGE_FAILED",
		50004: "INITIATE_DEPOSIT_FAILED",
		50005: "INITIATE_WITHDRAW_FAILED",
		50006: "DEPOSIT_CALLBACK_FAILED",
		50007: "WITHDRAW_CALLBACK_FAILED",
		50008: "GET_TRANSACTION_PAGE_FAILED",
		50009: "GET_TRANSACTION_DETAIL_FAILED",
		50010: "GET_CHANNEL_LIST_FAILED",
		50011: "GET_ADDRESS_FAILED",
		50012: "OPERATOR_ID_NOT_FOUND_IN_CONTEXT",
		50013: "UPDATE_PAYMENT_METHOD_FAILED",
		50014: "UPDATE_PAYMENT_CHANNEL_FAILED",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"GET_PAYMENT_METHOD_LIST_FAILED":   50001,
		"CREATE_PAYMENT_CHANNEL_FAILED":    50002,
		"GET_PAYMENT_CHANNEL_PAGE_FAILED":  50003,
		"INITIATE_DEPOSIT_FAILED":          50004,
		"INITIATE_WITHDRAW_FAILED":         50005,
		"DEPOSIT_CALLBACK_FAILED":          50006,
		"WITHDRAW_CALLBACK_FAILED":         50007,
		"GET_TRANSACTION_PAGE_FAILED":      50008,
		"GET_TRANSACTION_DETAIL_FAILED":    50009,
		"GET_CHANNEL_LIST_FAILED":          50010,
		"GET_ADDRESS_FAILED":               50011,
		"OPERATOR_ID_NOT_FOUND_IN_CONTEXT": 50012,
		"UPDATE_PAYMENT_METHOD_FAILED":     50013,
		"UPDATE_PAYMENT_CHANNEL_FAILED":    50014,
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
	return file_payment_service_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_payment_service_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_payment_service_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_payment_service_v1_error_reason_proto protoreflect.FileDescriptor

const file_payment_service_v1_error_reason_proto_rawDesc = "" +
	"\n" +
	"%payment/service/v1/error_reason.proto\x12\x16api.payment.service.v1\x1a\x13errors/errors.proto*\x86\x04\n" +
	"\vErrorReason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12$\n" +
	"\x1eGET_PAYMENT_METHOD_LIST_FAILED\x10ц\x03\x12#\n" +
	"\x1dCREATE_PAYMENT_CHANNEL_FAILED\x10҆\x03\x12%\n" +
	"\x1fGET_PAYMENT_CHANNEL_PAGE_FAILED\x10ӆ\x03\x12\x1d\n" +
	"\x17INITIATE_DEPOSIT_FAILED\x10Ԇ\x03\x12\x1e\n" +
	"\x18INITIATE_WITHDRAW_FAILED\x10Ն\x03\x12\x1d\n" +
	"\x17DEPOSIT_CALLBACK_FAILED\x10ֆ\x03\x12\x1e\n" +
	"\x18WITHDRAW_CALLBACK_FAILED\x10׆\x03\x12!\n" +
	"\x1bGET_TRANSACTION_PAGE_FAILED\x10؆\x03\x12#\n" +
	"\x1dGET_TRANSACTION_DETAIL_FAILED\x10ن\x03\x12\x1d\n" +
	"\x17GET_CHANNEL_LIST_FAILED\x10چ\x03\x12\x18\n" +
	"\x12GET_ADDRESS_FAILED\x10ۆ\x03\x12&\n" +
	" OPERATOR_ID_NOT_FOUND_IN_CONTEXT\x10܆\x03\x12\"\n" +
	"\x1cUPDATE_PAYMENT_METHOD_FAILED\x10݆\x03\x12#\n" +
	"\x1dUPDATE_PAYMENT_CHANNEL_FAILED\x10ކ\x03\x1a\x04\xa0E\xf4\x03BU\n" +
	"\x16api.payment.service.v1P\x01Z9github.com/infigaming-com/meepo-api/payment/service/v1;v1b\x06proto3"

var (
	file_payment_service_v1_error_reason_proto_rawDescOnce sync.Once
	file_payment_service_v1_error_reason_proto_rawDescData []byte
)

func file_payment_service_v1_error_reason_proto_rawDescGZIP() []byte {
	file_payment_service_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_payment_service_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payment_service_v1_error_reason_proto_rawDesc), len(file_payment_service_v1_error_reason_proto_rawDesc)))
	})
	return file_payment_service_v1_error_reason_proto_rawDescData
}

var file_payment_service_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_service_v1_error_reason_proto_goTypes = []any{
	(ErrorReason)(0), // 0: api.payment.service.v1.ErrorReason
}
var file_payment_service_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_service_v1_error_reason_proto_init() }
func file_payment_service_v1_error_reason_proto_init() {
	if File_payment_service_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payment_service_v1_error_reason_proto_rawDesc), len(file_payment_service_v1_error_reason_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_service_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_payment_service_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_payment_service_v1_error_reason_proto_enumTypes,
	}.Build()
	File_payment_service_v1_error_reason_proto = out.File
	file_payment_service_v1_error_reason_proto_goTypes = nil
	file_payment_service_v1_error_reason_proto_depIdxs = nil
}
