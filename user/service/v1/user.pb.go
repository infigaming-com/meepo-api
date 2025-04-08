// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.6
// source: user/service/v1/user.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AuthType int32

const (
	AuthType_AUTH_TYPE_UNSPECIFIED AuthType = 0
	AuthType_AUTH_TYPE_PASSWORD    AuthType = 1
	AuthType_AUTH_TYPE_GOOGLE      AuthType = 2
	AuthType_AUTH_TYPE_FACEBOOK    AuthType = 3
	AuthType_AUTH_TYPE_TELEGRAM    AuthType = 4
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "AUTH_TYPE_UNSPECIFIED",
		1: "AUTH_TYPE_PASSWORD",
		2: "AUTH_TYPE_GOOGLE",
		3: "AUTH_TYPE_FACEBOOK",
		4: "AUTH_TYPE_TELEGRAM",
	}
	AuthType_value = map[string]int32{
		"AUTH_TYPE_UNSPECIFIED": 0,
		"AUTH_TYPE_PASSWORD":    1,
		"AUTH_TYPE_GOOGLE":      2,
		"AUTH_TYPE_FACEBOOK":    3,
		"AUTH_TYPE_TELEGRAM":    4,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_v1_user_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_user_service_v1_user_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{0}
}

type IdentifierType int32

const (
	IdentifierType_IDENTIFIER_TYPE_UNSPECIFIED IdentifierType = 0
	IdentifierType_IDENTIFIER_TYPE_USERNAME    IdentifierType = 1
	IdentifierType_IDENTIFIER_TYPE_EMAIL       IdentifierType = 2
	IdentifierType_IDENTIFIER_TYPE_MOBILE      IdentifierType = 3
)

// Enum value maps for IdentifierType.
var (
	IdentifierType_name = map[int32]string{
		0: "IDENTIFIER_TYPE_UNSPECIFIED",
		1: "IDENTIFIER_TYPE_USERNAME",
		2: "IDENTIFIER_TYPE_EMAIL",
		3: "IDENTIFIER_TYPE_MOBILE",
	}
	IdentifierType_value = map[string]int32{
		"IDENTIFIER_TYPE_UNSPECIFIED": 0,
		"IDENTIFIER_TYPE_USERNAME":    1,
		"IDENTIFIER_TYPE_EMAIL":       2,
		"IDENTIFIER_TYPE_MOBILE":      3,
	}
)

func (x IdentifierType) Enum() *IdentifierType {
	p := new(IdentifierType)
	*p = x
	return p
}

func (x IdentifierType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentifierType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_v1_user_proto_enumTypes[1].Descriptor()
}

func (IdentifierType) Type() protoreflect.EnumType {
	return &file_user_service_v1_user_proto_enumTypes[1]
}

func (x IdentifierType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentifierType.Descriptor instead.
func (IdentifierType) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{1}
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_user_service_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RegisterRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	IdentifierType IdentifierType         `protobuf:"varint,1,opt,name=identifier_type,json=identifierType,proto3,enum=api.user.service.v1.IdentifierType" json:"identifier_type,omitempty"`
	Identifier     string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Password       string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetIdentifierType() IdentifierType {
	if x != nil {
		return x.IdentifierType
	}
	return IdentifierType_IDENTIFIER_TYPE_UNSPECIFIED
}

func (x *RegisterRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	IdentifierType IdentifierType         `protobuf:"varint,1,opt,name=identifier_type,json=identifierType,proto3,enum=api.user.service.v1.IdentifierType" json:"identifier_type,omitempty"`
	Identifier     string                 `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Password       string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetIdentifierType() IdentifierType {
	if x != nil {
		return x.IdentifierType
	}
	return IdentifierType_IDENTIFIER_TYPE_UNSPECIFIED
}

func (x *LoginRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type OAuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthType      AuthType               `protobuf:"varint,1,opt,name=auth_type,json=authType,proto3,enum=api.user.service.v1.AuthType" json:"auth_type,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuthRequest) Reset() {
	*x = OAuthRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthRequest) ProtoMessage() {}

func (x *OAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthRequest.ProtoReflect.Descriptor instead.
func (*OAuthRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *OAuthRequest) GetAuthType() AuthType {
	if x != nil {
		return x.AuthType
	}
	return AuthType_AUTH_TYPE_UNSPECIFIED
}

func (x *OAuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TelegramAuthRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	PhotoUrl      string                 `protobuf:"bytes,5,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	AuthDate      int64                  `protobuf:"varint,6,opt,name=auth_date,json=authDate,proto3" json:"auth_date,omitempty"`
	Hash          string                 `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TelegramAuthRequest) Reset() {
	*x = TelegramAuthRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TelegramAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelegramAuthRequest) ProtoMessage() {}

func (x *TelegramAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelegramAuthRequest.ProtoReflect.Descriptor instead.
func (*TelegramAuthRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *TelegramAuthRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TelegramAuthRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *TelegramAuthRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *TelegramAuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TelegramAuthRequest) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *TelegramAuthRequest) GetAuthDate() int64 {
	if x != nil {
		return x.AuthDate
	}
	return 0
}

func (x *TelegramAuthRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn     int64                  `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	UserInfo      *UserInfo              `protobuf:"bytes,4,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *AuthResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *AuthResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserInfo      *UserInfo              `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{9}
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_v1_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{10}
}

var File_user_service_v1_user_proto protoreflect.FileDescriptor

const file_user_service_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x1auser/service/v1/user.proto\x12\x13api.user.service.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17validate/validate.proto\"#\n" +
	"\bUserInfo\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"\x9b\x01\n" +
	"\x0fRegisterRequest\x12L\n" +
	"\x0fidentifier_type\x18\x01 \x01(\x0e2#.api.user.service.v1.IdentifierTypeR\x0eidentifierType\x12\x1e\n" +
	"\n" +
	"identifier\x18\x02 \x01(\tR\n" +
	"identifier\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"\x98\x01\n" +
	"\fLoginRequest\x12L\n" +
	"\x0fidentifier_type\x18\x01 \x01(\x0e2#.api.user.service.v1.IdentifierTypeR\x0eidentifierType\x12\x1e\n" +
	"\n" +
	"identifier\x18\x02 \x01(\tR\n" +
	"identifier\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"`\n" +
	"\fOAuthRequest\x12:\n" +
	"\tauth_type\x18\x01 \x01(\x0e2\x1d.api.user.service.v1.AuthTypeR\bauthType\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"\xcb\x01\n" +
	"\x13TelegramAuthRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x1a\n" +
	"\busername\x18\x04 \x01(\tR\busername\x12\x1b\n" +
	"\tphoto_url\x18\x05 \x01(\tR\bphotoUrl\x12\x1b\n" +
	"\tauth_date\x18\x06 \x01(\x03R\bauthDate\x12\x12\n" +
	"\x04hash\x18\a \x01(\tR\x04hash\":\n" +
	"\x13RefreshTokenRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"\xb1\x01\n" +
	"\fAuthResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x12\x1d\n" +
	"\n" +
	"expires_in\x18\x03 \x01(\x03R\texpiresIn\x12:\n" +
	"\tuser_info\x18\x04 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"M\n" +
	"\x0fGetUserResponse\x12:\n" +
	"\tuser_info\x18\x01 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo\"\x0f\n" +
	"\rLogoutRequest\"\x10\n" +
	"\x0eLogoutResponse*\x83\x01\n" +
	"\bAuthType\x12\x19\n" +
	"\x15AUTH_TYPE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12AUTH_TYPE_PASSWORD\x10\x01\x12\x14\n" +
	"\x10AUTH_TYPE_GOOGLE\x10\x02\x12\x16\n" +
	"\x12AUTH_TYPE_FACEBOOK\x10\x03\x12\x16\n" +
	"\x12AUTH_TYPE_TELEGRAM\x10\x04*\x86\x01\n" +
	"\x0eIdentifierType\x12\x1f\n" +
	"\x1bIDENTIFIER_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18IDENTIFIER_TYPE_USERNAME\x10\x01\x12\x19\n" +
	"\x15IDENTIFIER_TYPE_EMAIL\x10\x02\x12\x1a\n" +
	"\x16IDENTIFIER_TYPE_MOBILE\x10\x032\xfa\x06\n" +
	"\x04User\x12\x85\x01\n" +
	"\x17RegisterWithCredentials\x12$.api.user.service.v1.RegisterRequest\x1a!.api.user.service.v1.AuthResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v1/user/auth/register\x12|\n" +
	"\x14LoginWithCredentials\x12!.api.user.service.v1.LoginRequest\x1a!.api.user.service.v1.AuthResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v1/user/auth/login\x12\x80\x01\n" +
	"\x18RegisterOrLoginWithOAuth\x12!.api.user.service.v1.OAuthRequest\x1a!.api.user.service.v1.AuthResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v1/user/auth/oauth\x12\x8d\x01\n" +
	"\x1bRegisterOrLoginWithTelegram\x12(.api.user.service.v1.TelegramAuthRequest\x1a!.api.user.service.v1.AuthResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v1/user/auth/telegram\x12}\n" +
	"\fRefreshToken\x12(.api.user.service.v1.RefreshTokenRequest\x1a!.api.user.service.v1.AuthResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/v1/user/auth/refresh\x12k\n" +
	"\aGetUser\x12#.api.user.service.v1.GetUserRequest\x1a$.api.user.service.v1.GetUserResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/user/{id}\x12m\n" +
	"\x06Logout\x12\".api.user.service.v1.LogoutRequest\x1a#.api.user.service.v1.LogoutResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/v1/user/logoutBO\n" +
	"\x13api.user.service.v1P\x01Z6github.com/infigaming-com/meepo-api/user/service/v1;v1b\x06proto3"

var (
	file_user_service_v1_user_proto_rawDescOnce sync.Once
	file_user_service_v1_user_proto_rawDescData []byte
)

func file_user_service_v1_user_proto_rawDescGZIP() []byte {
	file_user_service_v1_user_proto_rawDescOnce.Do(func() {
		file_user_service_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_service_v1_user_proto_rawDesc), len(file_user_service_v1_user_proto_rawDesc)))
	})
	return file_user_service_v1_user_proto_rawDescData
}

var file_user_service_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_service_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_service_v1_user_proto_goTypes = []any{
	(AuthType)(0),               // 0: api.user.service.v1.AuthType
	(IdentifierType)(0),         // 1: api.user.service.v1.IdentifierType
	(*UserInfo)(nil),            // 2: api.user.service.v1.UserInfo
	(*RegisterRequest)(nil),     // 3: api.user.service.v1.RegisterRequest
	(*LoginRequest)(nil),        // 4: api.user.service.v1.LoginRequest
	(*OAuthRequest)(nil),        // 5: api.user.service.v1.OAuthRequest
	(*TelegramAuthRequest)(nil), // 6: api.user.service.v1.TelegramAuthRequest
	(*RefreshTokenRequest)(nil), // 7: api.user.service.v1.RefreshTokenRequest
	(*AuthResponse)(nil),        // 8: api.user.service.v1.AuthResponse
	(*GetUserRequest)(nil),      // 9: api.user.service.v1.GetUserRequest
	(*GetUserResponse)(nil),     // 10: api.user.service.v1.GetUserResponse
	(*LogoutRequest)(nil),       // 11: api.user.service.v1.LogoutRequest
	(*LogoutResponse)(nil),      // 12: api.user.service.v1.LogoutResponse
}
var file_user_service_v1_user_proto_depIdxs = []int32{
	1,  // 0: api.user.service.v1.RegisterRequest.identifier_type:type_name -> api.user.service.v1.IdentifierType
	1,  // 1: api.user.service.v1.LoginRequest.identifier_type:type_name -> api.user.service.v1.IdentifierType
	0,  // 2: api.user.service.v1.OAuthRequest.auth_type:type_name -> api.user.service.v1.AuthType
	2,  // 3: api.user.service.v1.AuthResponse.user_info:type_name -> api.user.service.v1.UserInfo
	2,  // 4: api.user.service.v1.GetUserResponse.user_info:type_name -> api.user.service.v1.UserInfo
	3,  // 5: api.user.service.v1.User.RegisterWithCredentials:input_type -> api.user.service.v1.RegisterRequest
	4,  // 6: api.user.service.v1.User.LoginWithCredentials:input_type -> api.user.service.v1.LoginRequest
	5,  // 7: api.user.service.v1.User.RegisterOrLoginWithOAuth:input_type -> api.user.service.v1.OAuthRequest
	6,  // 8: api.user.service.v1.User.RegisterOrLoginWithTelegram:input_type -> api.user.service.v1.TelegramAuthRequest
	7,  // 9: api.user.service.v1.User.RefreshToken:input_type -> api.user.service.v1.RefreshTokenRequest
	9,  // 10: api.user.service.v1.User.GetUser:input_type -> api.user.service.v1.GetUserRequest
	11, // 11: api.user.service.v1.User.Logout:input_type -> api.user.service.v1.LogoutRequest
	8,  // 12: api.user.service.v1.User.RegisterWithCredentials:output_type -> api.user.service.v1.AuthResponse
	8,  // 13: api.user.service.v1.User.LoginWithCredentials:output_type -> api.user.service.v1.AuthResponse
	8,  // 14: api.user.service.v1.User.RegisterOrLoginWithOAuth:output_type -> api.user.service.v1.AuthResponse
	8,  // 15: api.user.service.v1.User.RegisterOrLoginWithTelegram:output_type -> api.user.service.v1.AuthResponse
	8,  // 16: api.user.service.v1.User.RefreshToken:output_type -> api.user.service.v1.AuthResponse
	10, // 17: api.user.service.v1.User.GetUser:output_type -> api.user.service.v1.GetUserResponse
	12, // 18: api.user.service.v1.User.Logout:output_type -> api.user.service.v1.LogoutResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_user_service_v1_user_proto_init() }
func file_user_service_v1_user_proto_init() {
	if File_user_service_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_service_v1_user_proto_rawDesc), len(file_user_service_v1_user_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_v1_user_proto_goTypes,
		DependencyIndexes: file_user_service_v1_user_proto_depIdxs,
		EnumInfos:         file_user_service_v1_user_proto_enumTypes,
		MessageInfos:      file_user_service_v1_user_proto_msgTypes,
	}.Build()
	File_user_service_v1_user_proto = out.File
	file_user_service_v1_user_proto_goTypes = nil
	file_user_service_v1_user_proto_depIdxs = nil
}
