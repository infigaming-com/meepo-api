// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: user/service/v1/error_reason.proto

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
	ErrorReason_UNSPECIFIED                         ErrorReason = 0
	ErrorReason_USER_NOT_FOUND                      ErrorReason = 100000
	ErrorReason_USER_AUTH_NOT_FOUND                 ErrorReason = 100001
	ErrorReason_USER_ALREADY_EXISTS                 ErrorReason = 100002
	ErrorReason_REVOKE_TOKEN_FAILED                 ErrorReason = 100003
	ErrorReason_GENERATE_REFRESH_TOKEN_ID_FAILED    ErrorReason = 100004
	ErrorReason_GENERATE_TOKEN_ID_FAILED            ErrorReason = 100005
	ErrorReason_SAVE_TOKEN_FAILED                   ErrorReason = 100006
	ErrorReason_GET_USER_AUTH_FAILED                ErrorReason = 100007
	ErrorReason_GET_USER_FAILED                     ErrorReason = 100008
	ErrorReason_LOCK_USER_AUTH_FAILED               ErrorReason = 100009
	ErrorReason_HASH_USER_PASSWORD_FAILED           ErrorReason = 100010
	ErrorReason_GENERATE_USER_ID_FAILED             ErrorReason = 100011
	ErrorReason_CREATE_USER_FAILED                  ErrorReason = 100012
	ErrorReason_CREATE_USER_AUTH_FAILED             ErrorReason = 100013
	ErrorReason_USER_DISABLED                       ErrorReason = 100014
	ErrorReason_USER_LOGIN_BANNED                   ErrorReason = 100015
	ErrorReason_INVALID_USER_PASSWORD               ErrorReason = 100016
	ErrorReason_GET_TOKEN_WITH_REFRESH_TOKEN_FAILED ErrorReason = 100017
	ErrorReason_GET_UNEXPIRED_TOKEN_FOR_USER_FAILED ErrorReason = 100018
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:      "UNSPECIFIED",
		100000: "USER_NOT_FOUND",
		100001: "USER_AUTH_NOT_FOUND",
		100002: "USER_ALREADY_EXISTS",
		100003: "REVOKE_TOKEN_FAILED",
		100004: "GENERATE_REFRESH_TOKEN_ID_FAILED",
		100005: "GENERATE_TOKEN_ID_FAILED",
		100006: "SAVE_TOKEN_FAILED",
		100007: "GET_USER_AUTH_FAILED",
		100008: "GET_USER_FAILED",
		100009: "LOCK_USER_AUTH_FAILED",
		100010: "HASH_USER_PASSWORD_FAILED",
		100011: "GENERATE_USER_ID_FAILED",
		100012: "CREATE_USER_FAILED",
		100013: "CREATE_USER_AUTH_FAILED",
		100014: "USER_DISABLED",
		100015: "USER_LOGIN_BANNED",
		100016: "INVALID_USER_PASSWORD",
		100017: "GET_TOKEN_WITH_REFRESH_TOKEN_FAILED",
		100018: "GET_UNEXPIRED_TOKEN_FOR_USER_FAILED",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":                         0,
		"USER_NOT_FOUND":                      100000,
		"USER_AUTH_NOT_FOUND":                 100001,
		"USER_ALREADY_EXISTS":                 100002,
		"REVOKE_TOKEN_FAILED":                 100003,
		"GENERATE_REFRESH_TOKEN_ID_FAILED":    100004,
		"GENERATE_TOKEN_ID_FAILED":            100005,
		"SAVE_TOKEN_FAILED":                   100006,
		"GET_USER_AUTH_FAILED":                100007,
		"GET_USER_FAILED":                     100008,
		"LOCK_USER_AUTH_FAILED":               100009,
		"HASH_USER_PASSWORD_FAILED":           100010,
		"GENERATE_USER_ID_FAILED":             100011,
		"CREATE_USER_FAILED":                  100012,
		"CREATE_USER_AUTH_FAILED":             100013,
		"USER_DISABLED":                       100014,
		"USER_LOGIN_BANNED":                   100015,
		"INVALID_USER_PASSWORD":               100016,
		"GET_TOKEN_WITH_REFRESH_TOKEN_FAILED": 100017,
		"GET_UNEXPIRED_TOKEN_FOR_USER_FAILED": 100018,
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
	return file_user_service_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_user_service_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_user_service_v1_error_reason_proto protoreflect.FileDescriptor

const file_user_service_v1_error_reason_proto_rawDesc = "" +
	"\n" +
	"\"user/service/v1/error_reason.proto\x12\x13api.user.service.v1\x1a\x13errors/errors.proto*\xfa\x04\n" +
	"\vErrorReason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x0eUSER_NOT_FOUND\x10\xa0\x8d\x06\x1a\x04\xa8E\x94\x03\x12\x1f\n" +
	"\x13USER_AUTH_NOT_FOUND\x10\xa1\x8d\x06\x1a\x04\xa8E\x94\x03\x12\x1f\n" +
	"\x13USER_ALREADY_EXISTS\x10\xa2\x8d\x06\x1a\x04\xa8E\x99\x03\x12\x19\n" +
	"\x13REVOKE_TOKEN_FAILED\x10\xa3\x8d\x06\x12&\n" +
	" GENERATE_REFRESH_TOKEN_ID_FAILED\x10\xa4\x8d\x06\x12\x1e\n" +
	"\x18GENERATE_TOKEN_ID_FAILED\x10\xa5\x8d\x06\x12\x17\n" +
	"\x11SAVE_TOKEN_FAILED\x10\xa6\x8d\x06\x12\x1a\n" +
	"\x14GET_USER_AUTH_FAILED\x10\xa7\x8d\x06\x12\x15\n" +
	"\x0fGET_USER_FAILED\x10\xa8\x8d\x06\x12\x1b\n" +
	"\x15LOCK_USER_AUTH_FAILED\x10\xa9\x8d\x06\x12\x1f\n" +
	"\x19HASH_USER_PASSWORD_FAILED\x10\xaa\x8d\x06\x12\x1d\n" +
	"\x17GENERATE_USER_ID_FAILED\x10\xab\x8d\x06\x12\x18\n" +
	"\x12CREATE_USER_FAILED\x10\xac\x8d\x06\x12\x1d\n" +
	"\x17CREATE_USER_AUTH_FAILED\x10\xad\x8d\x06\x12\x19\n" +
	"\rUSER_DISABLED\x10\xae\x8d\x06\x1a\x04\xa8E\x91\x03\x12\x1d\n" +
	"\x11USER_LOGIN_BANNED\x10\xaf\x8d\x06\x1a\x04\xa8E\x91\x03\x12!\n" +
	"\x15INVALID_USER_PASSWORD\x10\xb0\x8d\x06\x1a\x04\xa8E\x91\x03\x12)\n" +
	"#GET_TOKEN_WITH_REFRESH_TOKEN_FAILED\x10\xb1\x8d\x06\x12)\n" +
	"#GET_UNEXPIRED_TOKEN_FOR_USER_FAILED\x10\xb2\x8d\x06\x1a\x04\xa0E\xf4\x03BO\n" +
	"\x13api.user.service.v1P\x01Z6github.com/infigaming-com/meepo-api/user/service/v1;v1b\x06proto3"

var (
	file_user_service_v1_error_reason_proto_rawDescOnce sync.Once
	file_user_service_v1_error_reason_proto_rawDescData []byte
)

func file_user_service_v1_error_reason_proto_rawDescGZIP() []byte {
	file_user_service_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_user_service_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_v1_error_reason_proto_rawDesc), len(file_user_service_v1_error_reason_proto_rawDesc)))
	})
	return file_user_service_v1_error_reason_proto_rawDescData
}

var file_user_service_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_v1_error_reason_proto_goTypes = []any{
	(ErrorReason)(0), // 0: api.user.service.v1.ErrorReason
}
var file_user_service_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_v1_error_reason_proto_init() }
func file_user_service_v1_error_reason_proto_init() {
	if File_user_service_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_v1_error_reason_proto_rawDesc), len(file_user_service_v1_error_reason_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_service_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_user_service_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_user_service_v1_error_reason_proto_enumTypes,
	}.Build()
	File_user_service_v1_error_reason_proto = out.File
	file_user_service_v1_error_reason_proto_goTypes = nil
	file_user_service_v1_error_reason_proto_depIdxs = nil
}
