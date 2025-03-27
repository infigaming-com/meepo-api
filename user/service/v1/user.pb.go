// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: user/service/v1/user.proto

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

type OAuthProvider int32

const (
	OAuthProvider_OAUTH_PROVIDER_UNSPECIFIED OAuthProvider = 0
	OAuthProvider_OAUTH_PROVIDER_GOOGLE      OAuthProvider = 1
	OAuthProvider_OAUTH_PROVIDER_TELEGRAM    OAuthProvider = 2
)

// Enum value maps for OAuthProvider.
var (
	OAuthProvider_name = map[int32]string{
		0: "OAUTH_PROVIDER_UNSPECIFIED",
		1: "OAUTH_PROVIDER_GOOGLE",
		2: "OAUTH_PROVIDER_TELEGRAM",
	}
	OAuthProvider_value = map[string]int32{
		"OAUTH_PROVIDER_UNSPECIFIED": 0,
		"OAUTH_PROVIDER_GOOGLE":      1,
		"OAUTH_PROVIDER_TELEGRAM":    2,
	}
)

func (x OAuthProvider) Enum() *OAuthProvider {
	p := new(OAuthProvider)
	*p = x
	return p
}

func (x OAuthProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OAuthProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_v1_user_proto_enumTypes[0].Descriptor()
}

func (OAuthProvider) Type() protoreflect.EnumType {
	return &file_user_service_v1_user_proto_enumTypes[0]
}

func (x OAuthProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OAuthProvider.Descriptor instead.
func (OAuthProvider) EnumDescriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{0}
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_user_service_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	UserInfo      *UserInfo              `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RegisterResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	UserInfo      *UserInfo              `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type OAuthLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      OAuthProvider          `protobuf:"varint,1,opt,name=provider,proto3,enum=api.user.service.v1.OAuthProvider" json:"provider,omitempty"`
	AuthCode      string                 `protobuf:"bytes,2,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	RedirectUri   string                 `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuthLoginRequest) Reset() {
	*x = OAuthLoginRequest{}
	mi := &file_user_service_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginRequest) ProtoMessage() {}

func (x *OAuthLoginRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OAuthLoginRequest.ProtoReflect.Descriptor instead.
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *OAuthLoginRequest) GetProvider() OAuthProvider {
	if x != nil {
		return x.Provider
	}
	return OAuthProvider_OAUTH_PROVIDER_UNSPECIFIED
}

func (x *OAuthLoginRequest) GetAuthCode() string {
	if x != nil {
		return x.AuthCode
	}
	return ""
}

func (x *OAuthLoginRequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type OAuthLoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	UserInfo      *UserInfo              `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuthLoginResponse) Reset() {
	*x = OAuthLoginResponse{}
	mi := &file_user_service_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginResponse) ProtoMessage() {}

func (x *OAuthLoginResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OAuthLoginResponse.ProtoReflect.Descriptor instead.
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *OAuthLoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *OAuthLoginResponse) GetUserInfo() *UserInfo {
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
	mi := &file_user_service_v1_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{9}
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
	mi := &file_user_service_v1_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserResponse) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_user_service_v1_user_proto protoreflect.FileDescriptor

const file_user_service_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x1auser/service/v1/user.proto\x12\x13api.user.service.v1\x1a\x1cgoogle/api/annotations.proto\"\x14\n" +
	"\x12HealthCheckRequest\"-\n" +
	"\x13HealthCheckResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"#\n" +
	"\bUserInfo\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"C\n" +
	"\x0fRegisterRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"q\n" +
	"\x10RegisterResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12:\n" +
	"\tuser_info\x18\x02 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"n\n" +
	"\rLoginResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12:\n" +
	"\tuser_info\x18\x02 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo\"\x93\x01\n" +
	"\x11OAuthLoginRequest\x12>\n" +
	"\bprovider\x18\x01 \x01(\x0e2\".api.user.service.v1.OAuthProviderR\bprovider\x12\x1b\n" +
	"\tauth_code\x18\x02 \x01(\tR\bauthCode\x12!\n" +
	"\fredirect_uri\x18\x03 \x01(\tR\vredirectUri\"s\n" +
	"\x12OAuthLoginResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12:\n" +
	"\tuser_info\x18\x02 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo\" \n" +
	"\x0eGetUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"M\n" +
	"\x0fGetUserResponse\x12:\n" +
	"\tuser_info\x18\x01 \x01(\v2\x1d.api.user.service.v1.UserInfoR\buserInfo*g\n" +
	"\rOAuthProvider\x12\x1e\n" +
	"\x1aOAUTH_PROVIDER_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15OAUTH_PROVIDER_GOOGLE\x10\x01\x12\x1b\n" +
	"\x17OAUTH_PROVIDER_TELEGRAM\x10\x022\xd0\x04\n" +
	"\x04User\x12y\n" +
	"\vHealthCheck\x12'.api.user.service.v1.HealthCheckRequest\x1a(.api.user.service.v1.HealthCheckResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/v1/user/health\x12u\n" +
	"\bRegister\x12$.api.user.service.v1.RegisterRequest\x1a%.api.user.service.v1.RegisterResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v1/user/register\x12i\n" +
	"\x05Login\x12!.api.user.service.v1.LoginRequest\x1a\".api.user.service.v1.LoginResponse\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\"\x0e/v1/user/login\x12~\n" +
	"\n" +
	"OAuthLogin\x12&.api.user.service.v1.OAuthLoginRequest\x1a'.api.user.service.v1.OAuthLoginResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v1/user/oauth/login\x12k\n" +
	"\aGetUser\x12#.api.user.service.v1.GetUserRequest\x1a$.api.user.service.v1.GetUserResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/v1/user/{id}BO\n" +
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

var file_user_service_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_service_v1_user_proto_goTypes = []any{
	(OAuthProvider)(0),          // 0: api.user.service.v1.OAuthProvider
	(*HealthCheckRequest)(nil),  // 1: api.user.service.v1.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 2: api.user.service.v1.HealthCheckResponse
	(*UserInfo)(nil),            // 3: api.user.service.v1.UserInfo
	(*RegisterRequest)(nil),     // 4: api.user.service.v1.RegisterRequest
	(*RegisterResponse)(nil),    // 5: api.user.service.v1.RegisterResponse
	(*LoginRequest)(nil),        // 6: api.user.service.v1.LoginRequest
	(*LoginResponse)(nil),       // 7: api.user.service.v1.LoginResponse
	(*OAuthLoginRequest)(nil),   // 8: api.user.service.v1.OAuthLoginRequest
	(*OAuthLoginResponse)(nil),  // 9: api.user.service.v1.OAuthLoginResponse
	(*GetUserRequest)(nil),      // 10: api.user.service.v1.GetUserRequest
	(*GetUserResponse)(nil),     // 11: api.user.service.v1.GetUserResponse
}
var file_user_service_v1_user_proto_depIdxs = []int32{
	3,  // 0: api.user.service.v1.RegisterResponse.user_info:type_name -> api.user.service.v1.UserInfo
	3,  // 1: api.user.service.v1.LoginResponse.user_info:type_name -> api.user.service.v1.UserInfo
	0,  // 2: api.user.service.v1.OAuthLoginRequest.provider:type_name -> api.user.service.v1.OAuthProvider
	3,  // 3: api.user.service.v1.OAuthLoginResponse.user_info:type_name -> api.user.service.v1.UserInfo
	3,  // 4: api.user.service.v1.GetUserResponse.user_info:type_name -> api.user.service.v1.UserInfo
	1,  // 5: api.user.service.v1.User.HealthCheck:input_type -> api.user.service.v1.HealthCheckRequest
	4,  // 6: api.user.service.v1.User.Register:input_type -> api.user.service.v1.RegisterRequest
	6,  // 7: api.user.service.v1.User.Login:input_type -> api.user.service.v1.LoginRequest
	8,  // 8: api.user.service.v1.User.OAuthLogin:input_type -> api.user.service.v1.OAuthLoginRequest
	10, // 9: api.user.service.v1.User.GetUser:input_type -> api.user.service.v1.GetUserRequest
	2,  // 10: api.user.service.v1.User.HealthCheck:output_type -> api.user.service.v1.HealthCheckResponse
	5,  // 11: api.user.service.v1.User.Register:output_type -> api.user.service.v1.RegisterResponse
	7,  // 12: api.user.service.v1.User.Login:output_type -> api.user.service.v1.LoginResponse
	9,  // 13: api.user.service.v1.User.OAuthLogin:output_type -> api.user.service.v1.OAuthLoginResponse
	11, // 14: api.user.service.v1.User.GetUser:output_type -> api.user.service.v1.GetUserResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
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
			NumEnums:      1,
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
