// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: public_proto/service_plminibooking.proto

package public_proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_public_proto_service_plminibooking_proto protoreflect.FileDescriptor

var file_public_proto_service_plminibooking_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x6d, 0x69, 0x6e, 0x69, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xe6, 0x02, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x67, 0x75,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x6f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6b, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70,
	0x62, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x32, 0x5a, 0x30, 0x70, 0x6c, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var file_public_proto_service_plminibooking_proto_goTypes = []any{
	(*CreateUserRequest)(nil),   // 0: public_pb.CreateUserRequest
	(*LoginUserRequest)(nil),    // 1: public_pb.LoginUserRequest
	(*VerifyEmailRequest)(nil),  // 2: public_pb.VerifyEmailRequest
	(*CreateUserResponse)(nil),  // 3: public_pb.CreateUserResponse
	(*LoginUserResponse)(nil),   // 4: public_pb.LoginUserResponse
	(*VerifyEmailResponse)(nil), // 5: public_pb.VerifyEmailResponse
}
var file_public_proto_service_plminibooking_proto_depIdxs = []int32{
	0, // 0: public_pb.PremierLeagueBooking.CreateUser:input_type -> public_pb.CreateUserRequest
	1, // 1: public_pb.PremierLeagueBooking.LoginUser:input_type -> public_pb.LoginUserRequest
	2, // 2: public_pb.PremierLeagueBooking.VerifyEmail:input_type -> public_pb.VerifyEmailRequest
	3, // 3: public_pb.PremierLeagueBooking.CreateUser:output_type -> public_pb.CreateUserResponse
	4, // 4: public_pb.PremierLeagueBooking.LoginUser:output_type -> public_pb.LoginUserResponse
	5, // 5: public_pb.PremierLeagueBooking.VerifyEmail:output_type -> public_pb.VerifyEmailResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_proto_service_plminibooking_proto_init() }
func file_public_proto_service_plminibooking_proto_init() {
	if File_public_proto_service_plminibooking_proto != nil {
		return
	}
	file_public_proto_rpc_create_user_proto_init()
	file_public_proto_rpc_login_user_proto_init()
	file_public_proto_rpc_verify_email_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_public_proto_service_plminibooking_proto_rawDesc), len(file_public_proto_service_plminibooking_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_proto_service_plminibooking_proto_goTypes,
		DependencyIndexes: file_public_proto_service_plminibooking_proto_depIdxs,
	}.Build()
	File_public_proto_service_plminibooking_proto = out.File
	file_public_proto_service_plminibooking_proto_goTypes = nil
	file_public_proto_service_plminibooking_proto_depIdxs = nil
}
