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
	ErrorReason_UNSPECIFIED                              ErrorReason = 0
	ErrorReason_USER_INFO_NOT_FOUND_IN_CONTEXT           ErrorReason = 10000
	ErrorReason_REQUEST_INFO_NOT_FOUND_IN_CONTEXT        ErrorReason = 10001
	ErrorReason_USER_NOT_FOUND                           ErrorReason = 10002
	ErrorReason_USER_AUTH_NOT_FOUND                      ErrorReason = 10003
	ErrorReason_USER_ALREADY_EXISTS                      ErrorReason = 10004
	ErrorReason_REVOKE_TOKEN_FAILED                      ErrorReason = 10005
	ErrorReason_GENERATE_REFRESH_TOKEN_ID_FAILED         ErrorReason = 10006
	ErrorReason_GENERATE_TOKEN_ID_FAILED                 ErrorReason = 10007
	ErrorReason_SAVE_TOKEN_FAILED                        ErrorReason = 10008
	ErrorReason_GET_USER_AUTH_FAILED                     ErrorReason = 10009
	ErrorReason_GET_USER_FAILED                          ErrorReason = 10010
	ErrorReason_UPDATE_USER_FAILED                       ErrorReason = 10011
	ErrorReason_LOCK_USER_AUTH_FAILED                    ErrorReason = 10012
	ErrorReason_HASH_USER_PASSWORD_FAILED                ErrorReason = 10013
	ErrorReason_GENERATE_USER_ID_FAILED                  ErrorReason = 10014
	ErrorReason_ADD_USER_FAILED                          ErrorReason = 10015
	ErrorReason_ADD_USER_AUTH_FAILED                     ErrorReason = 10016
	ErrorReason_USER_DISABLED                            ErrorReason = 10017
	ErrorReason_USER_LOGIN_BANNED                        ErrorReason = 10018
	ErrorReason_INVALID_USER_PASSWORD                    ErrorReason = 10019
	ErrorReason_GET_TOKEN_WITH_REFRESH_TOKEN_FAILED      ErrorReason = 10020
	ErrorReason_GET_UNEXPIRED_TOKEN_FOR_USER_FAILED      ErrorReason = 10021
	ErrorReason_OAUTH_PROVIDER_NOT_SUPPORTED             ErrorReason = 10022
	ErrorReason_USERNAME_OR_PASSWORD_INVALID             ErrorReason = 10023
	ErrorReason_ADD_USER_TO_WALLET_FAILED                ErrorReason = 10024
	ErrorReason_GET_USERS_BY_OPERATOR_IDS_FAILED         ErrorReason = 10025
	ErrorReason_GET_OPERATOR_ID_BY_ORIGIN_FAILED         ErrorReason = 10026
	ErrorReason_USER_TAG_ALREADY_EXISTS                  ErrorReason = 10027
	ErrorReason_USER_TAG_NOT_EXIST                       ErrorReason = 10030
	ErrorReason_VERIFY_GOOGLE_TOKEN_FAILED               ErrorReason = 10031
	ErrorReason_MARSHAL_REQUEST_INFO_FAILED              ErrorReason = 10032
	ErrorReason_ADD_USER_EVENT_FAILED                    ErrorReason = 10033
	ErrorReason_GENERATE_USER_EVENT_ID_FAILED            ErrorReason = 10034
	ErrorReason_FOLLOW_PARENT_ENABLED                    ErrorReason = 10035 // Operator is in follow_parent=true mode
	ErrorReason_OPERATOR_TAGS_ALREADY_EXISTS             ErrorReason = 10036
	ErrorReason_OPERATOR_TAG_NOT_FOUND                   ErrorReason = 10037
	ErrorReason_OPERATOR_PARENT_NOT_FOUND                ErrorReason = 10038
	ErrorReason_NON_FOLLOW_PARENT_OPERATOR_NOT_FOUND     ErrorReason = 10039
	ErrorReason_SYSTEM_OPERATOR                          ErrorReason = 10040
	ErrorReason_LIST_USERS_FAILED                        ErrorReason = 10041
	ErrorReason_ROLE_NOT_FOUND                           ErrorReason = 10042
	ErrorReason_FAILED_TO_SEND_EMAIL                     ErrorReason = 10043
	ErrorReason_VERIFICATION_CODE_SEND_TOO_FREQUENTLY    ErrorReason = 10044
	ErrorReason_EMAIL_VERIFICATION_FAILED                ErrorReason = 10045
	ErrorReason_EMAIL_ALREADY_REGISTERED                 ErrorReason = 10046
	ErrorReason_GENERATE_COMMENT_ID_FAILED               ErrorReason = 10047
	ErrorReason_ADD_COMMENT_FAILED                       ErrorReason = 10048
	ErrorReason_GET_COMMENTS_BY_USER_ID_FAILED           ErrorReason = 10049
	ErrorReason_USER_TAGS_NOT_MATCH_OPERATOR_TAGS        ErrorReason = 10050
	ErrorReason_INVALID_OPERATOR_TYPE                    ErrorReason = 10051
	ErrorReason_ADD_USER_DAILY_ACTIVITY_FAILED           ErrorReason = 10052
	ErrorReason_OPERATOR_ID_NOT_FOUND_BY_ORIGIN          ErrorReason = 10053
	ErrorReason_OPERATOR_ID_NOT_FOUND_IN_CONTEXT         ErrorReason = 10054
	ErrorReason_LIST_OPERATORS_FAILED                    ErrorReason = 10055
	ErrorReason_GET_OPERATORS_BY_IDS_FAILED              ErrorReason = 10056
	ErrorReason_OPERATOR_HIERARCHY_TOO_DEEP              ErrorReason = 10057
	ErrorReason_GET_PARENT_OPERATOR_IDS_FAILED           ErrorReason = 10058
	ErrorReason_CREATE_ORIGIN_TO_OPERATOR_MAPPING_FAILED ErrorReason = 10059
	ErrorReason_ORIGIN_ALREADY_EXISTS                    ErrorReason = 10060
	ErrorReason_GENERATE_BUSINESS_ID_FAILED              ErrorReason = 10061
	ErrorReason_CREATE_BUSINESS_FAILED                   ErrorReason = 10062
	ErrorReason_OPERATOR_IDS_NOT_FOUND_BY_ORIGIN         ErrorReason = 10063
	ErrorReason_OPERATOR_IDS_NOT_FOUND_IN_CONTEXT        ErrorReason = 10064
	ErrorReason_ACTUAL_OPERATOR_ID_NOT_GET_IN_CONTEXT    ErrorReason = 10065
	ErrorReason_INVALID_OPERATOR_KEY                     ErrorReason = 10066
	ErrorReason_GET_OPERATOR_BY_KEY_FAILED               ErrorReason = 10067
	ErrorReason_OPERATOR_KEY_ALREADY_EXISTS              ErrorReason = 10068
	ErrorReason_EMPTY_PASSWORD                           ErrorReason = 10069
	ErrorReason_EMAIL_ALREADY_EXISTS                     ErrorReason = 10070
	ErrorReason_INVALID_OPERATOR_STATUS                  ErrorReason = 10071
	ErrorReason_DOMAIN_POOL_EMPTY                        ErrorReason = 10072
	ErrorReason_TARGET_OPERATOR_NOT_BELONG_TO_OPERATOR   ErrorReason = 10073
	ErrorReason_OPERATOR_NOT_FOUND_BY_ID                 ErrorReason = 10074
	ErrorReason_GET_OPERATOR_BY_ID_FAILED                ErrorReason = 10075
	ErrorReason_INVALID_OPERATOR_TRANSITION_ACTION       ErrorReason = 10076
	ErrorReason_UPDATE_OPERATOR_STATUS_FAILED            ErrorReason = 10077
	ErrorReason_INVALID_OPERATOR_CONTEXT                 ErrorReason = 10078
	ErrorReason_INVALID_OPERATOR_ID                      ErrorReason = 10079
	ErrorReason_OPERATOR_CONTEXT_PERMISSION_DENIED       ErrorReason = 10080
	ErrorReason_OPERATOR_CONTEXT_NOT_FOUND_IN_CONTEXT    ErrorReason = 10081
	ErrorReason_GET_ADMIN_USERS_BY_EMAIL_FAILED          ErrorReason = 10082
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:     "UNSPECIFIED",
		10000: "USER_INFO_NOT_FOUND_IN_CONTEXT",
		10001: "REQUEST_INFO_NOT_FOUND_IN_CONTEXT",
		10002: "USER_NOT_FOUND",
		10003: "USER_AUTH_NOT_FOUND",
		10004: "USER_ALREADY_EXISTS",
		10005: "REVOKE_TOKEN_FAILED",
		10006: "GENERATE_REFRESH_TOKEN_ID_FAILED",
		10007: "GENERATE_TOKEN_ID_FAILED",
		10008: "SAVE_TOKEN_FAILED",
		10009: "GET_USER_AUTH_FAILED",
		10010: "GET_USER_FAILED",
		10011: "UPDATE_USER_FAILED",
		10012: "LOCK_USER_AUTH_FAILED",
		10013: "HASH_USER_PASSWORD_FAILED",
		10014: "GENERATE_USER_ID_FAILED",
		10015: "ADD_USER_FAILED",
		10016: "ADD_USER_AUTH_FAILED",
		10017: "USER_DISABLED",
		10018: "USER_LOGIN_BANNED",
		10019: "INVALID_USER_PASSWORD",
		10020: "GET_TOKEN_WITH_REFRESH_TOKEN_FAILED",
		10021: "GET_UNEXPIRED_TOKEN_FOR_USER_FAILED",
		10022: "OAUTH_PROVIDER_NOT_SUPPORTED",
		10023: "USERNAME_OR_PASSWORD_INVALID",
		10024: "ADD_USER_TO_WALLET_FAILED",
		10025: "GET_USERS_BY_OPERATOR_IDS_FAILED",
		10026: "GET_OPERATOR_ID_BY_ORIGIN_FAILED",
		10027: "USER_TAG_ALREADY_EXISTS",
		10030: "USER_TAG_NOT_EXIST",
		10031: "VERIFY_GOOGLE_TOKEN_FAILED",
		10032: "MARSHAL_REQUEST_INFO_FAILED",
		10033: "ADD_USER_EVENT_FAILED",
		10034: "GENERATE_USER_EVENT_ID_FAILED",
		10035: "FOLLOW_PARENT_ENABLED",
		10036: "OPERATOR_TAGS_ALREADY_EXISTS",
		10037: "OPERATOR_TAG_NOT_FOUND",
		10038: "OPERATOR_PARENT_NOT_FOUND",
		10039: "NON_FOLLOW_PARENT_OPERATOR_NOT_FOUND",
		10040: "SYSTEM_OPERATOR",
		10041: "LIST_USERS_FAILED",
		10042: "ROLE_NOT_FOUND",
		10043: "FAILED_TO_SEND_EMAIL",
		10044: "VERIFICATION_CODE_SEND_TOO_FREQUENTLY",
		10045: "EMAIL_VERIFICATION_FAILED",
		10046: "EMAIL_ALREADY_REGISTERED",
		10047: "GENERATE_COMMENT_ID_FAILED",
		10048: "ADD_COMMENT_FAILED",
		10049: "GET_COMMENTS_BY_USER_ID_FAILED",
		10050: "USER_TAGS_NOT_MATCH_OPERATOR_TAGS",
		10051: "INVALID_OPERATOR_TYPE",
		10052: "ADD_USER_DAILY_ACTIVITY_FAILED",
		10053: "OPERATOR_ID_NOT_FOUND_BY_ORIGIN",
		10054: "OPERATOR_ID_NOT_FOUND_IN_CONTEXT",
		10055: "LIST_OPERATORS_FAILED",
		10056: "GET_OPERATORS_BY_IDS_FAILED",
		10057: "OPERATOR_HIERARCHY_TOO_DEEP",
		10058: "GET_PARENT_OPERATOR_IDS_FAILED",
		10059: "CREATE_ORIGIN_TO_OPERATOR_MAPPING_FAILED",
		10060: "ORIGIN_ALREADY_EXISTS",
		10061: "GENERATE_BUSINESS_ID_FAILED",
		10062: "CREATE_BUSINESS_FAILED",
		10063: "OPERATOR_IDS_NOT_FOUND_BY_ORIGIN",
		10064: "OPERATOR_IDS_NOT_FOUND_IN_CONTEXT",
		10065: "ACTUAL_OPERATOR_ID_NOT_GET_IN_CONTEXT",
		10066: "INVALID_OPERATOR_KEY",
		10067: "GET_OPERATOR_BY_KEY_FAILED",
		10068: "OPERATOR_KEY_ALREADY_EXISTS",
		10069: "EMPTY_PASSWORD",
		10070: "EMAIL_ALREADY_EXISTS",
		10071: "INVALID_OPERATOR_STATUS",
		10072: "DOMAIN_POOL_EMPTY",
		10073: "TARGET_OPERATOR_NOT_BELONG_TO_OPERATOR",
		10074: "OPERATOR_NOT_FOUND_BY_ID",
		10075: "GET_OPERATOR_BY_ID_FAILED",
		10076: "INVALID_OPERATOR_TRANSITION_ACTION",
		10077: "UPDATE_OPERATOR_STATUS_FAILED",
		10078: "INVALID_OPERATOR_CONTEXT",
		10079: "INVALID_OPERATOR_ID",
		10080: "OPERATOR_CONTEXT_PERMISSION_DENIED",
		10081: "OPERATOR_CONTEXT_NOT_FOUND_IN_CONTEXT",
		10082: "GET_ADMIN_USERS_BY_EMAIL_FAILED",
	}
	ErrorReason_value = map[string]int32{
		"UNSPECIFIED":                              0,
		"USER_INFO_NOT_FOUND_IN_CONTEXT":           10000,
		"REQUEST_INFO_NOT_FOUND_IN_CONTEXT":        10001,
		"USER_NOT_FOUND":                           10002,
		"USER_AUTH_NOT_FOUND":                      10003,
		"USER_ALREADY_EXISTS":                      10004,
		"REVOKE_TOKEN_FAILED":                      10005,
		"GENERATE_REFRESH_TOKEN_ID_FAILED":         10006,
		"GENERATE_TOKEN_ID_FAILED":                 10007,
		"SAVE_TOKEN_FAILED":                        10008,
		"GET_USER_AUTH_FAILED":                     10009,
		"GET_USER_FAILED":                          10010,
		"UPDATE_USER_FAILED":                       10011,
		"LOCK_USER_AUTH_FAILED":                    10012,
		"HASH_USER_PASSWORD_FAILED":                10013,
		"GENERATE_USER_ID_FAILED":                  10014,
		"ADD_USER_FAILED":                          10015,
		"ADD_USER_AUTH_FAILED":                     10016,
		"USER_DISABLED":                            10017,
		"USER_LOGIN_BANNED":                        10018,
		"INVALID_USER_PASSWORD":                    10019,
		"GET_TOKEN_WITH_REFRESH_TOKEN_FAILED":      10020,
		"GET_UNEXPIRED_TOKEN_FOR_USER_FAILED":      10021,
		"OAUTH_PROVIDER_NOT_SUPPORTED":             10022,
		"USERNAME_OR_PASSWORD_INVALID":             10023,
		"ADD_USER_TO_WALLET_FAILED":                10024,
		"GET_USERS_BY_OPERATOR_IDS_FAILED":         10025,
		"GET_OPERATOR_ID_BY_ORIGIN_FAILED":         10026,
		"USER_TAG_ALREADY_EXISTS":                  10027,
		"USER_TAG_NOT_EXIST":                       10030,
		"VERIFY_GOOGLE_TOKEN_FAILED":               10031,
		"MARSHAL_REQUEST_INFO_FAILED":              10032,
		"ADD_USER_EVENT_FAILED":                    10033,
		"GENERATE_USER_EVENT_ID_FAILED":            10034,
		"FOLLOW_PARENT_ENABLED":                    10035,
		"OPERATOR_TAGS_ALREADY_EXISTS":             10036,
		"OPERATOR_TAG_NOT_FOUND":                   10037,
		"OPERATOR_PARENT_NOT_FOUND":                10038,
		"NON_FOLLOW_PARENT_OPERATOR_NOT_FOUND":     10039,
		"SYSTEM_OPERATOR":                          10040,
		"LIST_USERS_FAILED":                        10041,
		"ROLE_NOT_FOUND":                           10042,
		"FAILED_TO_SEND_EMAIL":                     10043,
		"VERIFICATION_CODE_SEND_TOO_FREQUENTLY":    10044,
		"EMAIL_VERIFICATION_FAILED":                10045,
		"EMAIL_ALREADY_REGISTERED":                 10046,
		"GENERATE_COMMENT_ID_FAILED":               10047,
		"ADD_COMMENT_FAILED":                       10048,
		"GET_COMMENTS_BY_USER_ID_FAILED":           10049,
		"USER_TAGS_NOT_MATCH_OPERATOR_TAGS":        10050,
		"INVALID_OPERATOR_TYPE":                    10051,
		"ADD_USER_DAILY_ACTIVITY_FAILED":           10052,
		"OPERATOR_ID_NOT_FOUND_BY_ORIGIN":          10053,
		"OPERATOR_ID_NOT_FOUND_IN_CONTEXT":         10054,
		"LIST_OPERATORS_FAILED":                    10055,
		"GET_OPERATORS_BY_IDS_FAILED":              10056,
		"OPERATOR_HIERARCHY_TOO_DEEP":              10057,
		"GET_PARENT_OPERATOR_IDS_FAILED":           10058,
		"CREATE_ORIGIN_TO_OPERATOR_MAPPING_FAILED": 10059,
		"ORIGIN_ALREADY_EXISTS":                    10060,
		"GENERATE_BUSINESS_ID_FAILED":              10061,
		"CREATE_BUSINESS_FAILED":                   10062,
		"OPERATOR_IDS_NOT_FOUND_BY_ORIGIN":         10063,
		"OPERATOR_IDS_NOT_FOUND_IN_CONTEXT":        10064,
		"ACTUAL_OPERATOR_ID_NOT_GET_IN_CONTEXT":    10065,
		"INVALID_OPERATOR_KEY":                     10066,
		"GET_OPERATOR_BY_KEY_FAILED":               10067,
		"OPERATOR_KEY_ALREADY_EXISTS":              10068,
		"EMPTY_PASSWORD":                           10069,
		"EMAIL_ALREADY_EXISTS":                     10070,
		"INVALID_OPERATOR_STATUS":                  10071,
		"DOMAIN_POOL_EMPTY":                        10072,
		"TARGET_OPERATOR_NOT_BELONG_TO_OPERATOR":   10073,
		"OPERATOR_NOT_FOUND_BY_ID":                 10074,
		"GET_OPERATOR_BY_ID_FAILED":                10075,
		"INVALID_OPERATOR_TRANSITION_ACTION":       10076,
		"UPDATE_OPERATOR_STATUS_FAILED":            10077,
		"INVALID_OPERATOR_CONTEXT":                 10078,
		"INVALID_OPERATOR_ID":                      10079,
		"OPERATOR_CONTEXT_PERMISSION_DENIED":       10080,
		"OPERATOR_CONTEXT_NOT_FOUND_IN_CONTEXT":    10081,
		"GET_ADMIN_USERS_BY_EMAIL_FAILED":          10082,
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
	"\"user/service/v1/error_reason.proto\x12\x13api.user.service.v1\x1a\x13errors/errors.proto*\xd1\x14\n" +
	"\vErrorReason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12#\n" +
	"\x1eUSER_INFO_NOT_FOUND_IN_CONTEXT\x10\x90N\x12&\n" +
	"!REQUEST_INFO_NOT_FOUND_IN_CONTEXT\x10\x91N\x12\x13\n" +
	"\x0eUSER_NOT_FOUND\x10\x92N\x12\x18\n" +
	"\x13USER_AUTH_NOT_FOUND\x10\x93N\x12\x18\n" +
	"\x13USER_ALREADY_EXISTS\x10\x94N\x12\x18\n" +
	"\x13REVOKE_TOKEN_FAILED\x10\x95N\x12%\n" +
	" GENERATE_REFRESH_TOKEN_ID_FAILED\x10\x96N\x12\x1d\n" +
	"\x18GENERATE_TOKEN_ID_FAILED\x10\x97N\x12\x16\n" +
	"\x11SAVE_TOKEN_FAILED\x10\x98N\x12\x19\n" +
	"\x14GET_USER_AUTH_FAILED\x10\x99N\x12\x14\n" +
	"\x0fGET_USER_FAILED\x10\x9aN\x12\x17\n" +
	"\x12UPDATE_USER_FAILED\x10\x9bN\x12\x1a\n" +
	"\x15LOCK_USER_AUTH_FAILED\x10\x9cN\x12\x1e\n" +
	"\x19HASH_USER_PASSWORD_FAILED\x10\x9dN\x12\x1c\n" +
	"\x17GENERATE_USER_ID_FAILED\x10\x9eN\x12\x14\n" +
	"\x0fADD_USER_FAILED\x10\x9fN\x12\x19\n" +
	"\x14ADD_USER_AUTH_FAILED\x10\xa0N\x12\x12\n" +
	"\rUSER_DISABLED\x10\xa1N\x12\x16\n" +
	"\x11USER_LOGIN_BANNED\x10\xa2N\x12\x1a\n" +
	"\x15INVALID_USER_PASSWORD\x10\xa3N\x12(\n" +
	"#GET_TOKEN_WITH_REFRESH_TOKEN_FAILED\x10\xa4N\x12(\n" +
	"#GET_UNEXPIRED_TOKEN_FOR_USER_FAILED\x10\xa5N\x12!\n" +
	"\x1cOAUTH_PROVIDER_NOT_SUPPORTED\x10\xa6N\x12!\n" +
	"\x1cUSERNAME_OR_PASSWORD_INVALID\x10\xa7N\x12\x1e\n" +
	"\x19ADD_USER_TO_WALLET_FAILED\x10\xa8N\x12%\n" +
	" GET_USERS_BY_OPERATOR_IDS_FAILED\x10\xa9N\x12%\n" +
	" GET_OPERATOR_ID_BY_ORIGIN_FAILED\x10\xaaN\x12\x1c\n" +
	"\x17USER_TAG_ALREADY_EXISTS\x10\xabN\x12\x17\n" +
	"\x12USER_TAG_NOT_EXIST\x10\xaeN\x12\x1f\n" +
	"\x1aVERIFY_GOOGLE_TOKEN_FAILED\x10\xafN\x12 \n" +
	"\x1bMARSHAL_REQUEST_INFO_FAILED\x10\xb0N\x12\x1a\n" +
	"\x15ADD_USER_EVENT_FAILED\x10\xb1N\x12\"\n" +
	"\x1dGENERATE_USER_EVENT_ID_FAILED\x10\xb2N\x12\x1a\n" +
	"\x15FOLLOW_PARENT_ENABLED\x10\xb3N\x12!\n" +
	"\x1cOPERATOR_TAGS_ALREADY_EXISTS\x10\xb4N\x12\x1b\n" +
	"\x16OPERATOR_TAG_NOT_FOUND\x10\xb5N\x12\x1e\n" +
	"\x19OPERATOR_PARENT_NOT_FOUND\x10\xb6N\x12)\n" +
	"$NON_FOLLOW_PARENT_OPERATOR_NOT_FOUND\x10\xb7N\x12\x14\n" +
	"\x0fSYSTEM_OPERATOR\x10\xb8N\x12\x16\n" +
	"\x11LIST_USERS_FAILED\x10\xb9N\x12\x13\n" +
	"\x0eROLE_NOT_FOUND\x10\xbaN\x12\x19\n" +
	"\x14FAILED_TO_SEND_EMAIL\x10\xbbN\x12*\n" +
	"%VERIFICATION_CODE_SEND_TOO_FREQUENTLY\x10\xbcN\x12\x1e\n" +
	"\x19EMAIL_VERIFICATION_FAILED\x10\xbdN\x12\x1d\n" +
	"\x18EMAIL_ALREADY_REGISTERED\x10\xbeN\x12\x1f\n" +
	"\x1aGENERATE_COMMENT_ID_FAILED\x10\xbfN\x12\x17\n" +
	"\x12ADD_COMMENT_FAILED\x10\xc0N\x12#\n" +
	"\x1eGET_COMMENTS_BY_USER_ID_FAILED\x10\xc1N\x12&\n" +
	"!USER_TAGS_NOT_MATCH_OPERATOR_TAGS\x10\xc2N\x12\x1a\n" +
	"\x15INVALID_OPERATOR_TYPE\x10\xc3N\x12#\n" +
	"\x1eADD_USER_DAILY_ACTIVITY_FAILED\x10\xc4N\x12$\n" +
	"\x1fOPERATOR_ID_NOT_FOUND_BY_ORIGIN\x10\xc5N\x12%\n" +
	" OPERATOR_ID_NOT_FOUND_IN_CONTEXT\x10\xc6N\x12\x1a\n" +
	"\x15LIST_OPERATORS_FAILED\x10\xc7N\x12 \n" +
	"\x1bGET_OPERATORS_BY_IDS_FAILED\x10\xc8N\x12 \n" +
	"\x1bOPERATOR_HIERARCHY_TOO_DEEP\x10\xc9N\x12#\n" +
	"\x1eGET_PARENT_OPERATOR_IDS_FAILED\x10\xcaN\x12-\n" +
	"(CREATE_ORIGIN_TO_OPERATOR_MAPPING_FAILED\x10\xcbN\x12\x1a\n" +
	"\x15ORIGIN_ALREADY_EXISTS\x10\xccN\x12 \n" +
	"\x1bGENERATE_BUSINESS_ID_FAILED\x10\xcdN\x12\x1b\n" +
	"\x16CREATE_BUSINESS_FAILED\x10\xceN\x12%\n" +
	" OPERATOR_IDS_NOT_FOUND_BY_ORIGIN\x10\xcfN\x12&\n" +
	"!OPERATOR_IDS_NOT_FOUND_IN_CONTEXT\x10\xd0N\x12*\n" +
	"%ACTUAL_OPERATOR_ID_NOT_GET_IN_CONTEXT\x10\xd1N\x12\x19\n" +
	"\x14INVALID_OPERATOR_KEY\x10\xd2N\x12\x1f\n" +
	"\x1aGET_OPERATOR_BY_KEY_FAILED\x10\xd3N\x12 \n" +
	"\x1bOPERATOR_KEY_ALREADY_EXISTS\x10\xd4N\x12\x13\n" +
	"\x0eEMPTY_PASSWORD\x10\xd5N\x12\x19\n" +
	"\x14EMAIL_ALREADY_EXISTS\x10\xd6N\x12\x1c\n" +
	"\x17INVALID_OPERATOR_STATUS\x10\xd7N\x12\x16\n" +
	"\x11DOMAIN_POOL_EMPTY\x10\xd8N\x12+\n" +
	"&TARGET_OPERATOR_NOT_BELONG_TO_OPERATOR\x10\xd9N\x12\x1d\n" +
	"\x18OPERATOR_NOT_FOUND_BY_ID\x10\xdaN\x12\x1e\n" +
	"\x19GET_OPERATOR_BY_ID_FAILED\x10\xdbN\x12'\n" +
	"\"INVALID_OPERATOR_TRANSITION_ACTION\x10\xdcN\x12\"\n" +
	"\x1dUPDATE_OPERATOR_STATUS_FAILED\x10\xddN\x12\x1d\n" +
	"\x18INVALID_OPERATOR_CONTEXT\x10\xdeN\x12\x18\n" +
	"\x13INVALID_OPERATOR_ID\x10\xdfN\x12'\n" +
	"\"OPERATOR_CONTEXT_PERMISSION_DENIED\x10\xe0N\x12*\n" +
	"%OPERATOR_CONTEXT_NOT_FOUND_IN_CONTEXT\x10\xe1N\x12$\n" +
	"\x1fGET_ADMIN_USERS_BY_EMAIL_FAILED\x10\xe2N\x1a\x04\xa0E\xf4\x03BO\n" +
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
