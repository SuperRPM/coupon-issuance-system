// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/coupon/v1/coupon.proto

package couponv1

import (
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

type IssueCouponRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CampaignId    int32                  `protobuf:"varint,1,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueCouponRequest) Reset() {
	*x = IssueCouponRequest{}
	mi := &file_proto_coupon_v1_coupon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCouponRequest) ProtoMessage() {}

func (x *IssueCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_coupon_v1_coupon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCouponRequest.ProtoReflect.Descriptor instead.
func (*IssueCouponRequest) Descriptor() ([]byte, []int) {
	return file_proto_coupon_v1_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *IssueCouponRequest) GetCampaignId() int32 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

type IssueCouponResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CouponId      int32                  `protobuf:"varint,1,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id,omitempty"`
	CouponCode    string                 `protobuf:"bytes,2,opt,name=coupon_code,json=couponCode,proto3" json:"coupon_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueCouponResponse) Reset() {
	*x = IssueCouponResponse{}
	mi := &file_proto_coupon_v1_coupon_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCouponResponse) ProtoMessage() {}

func (x *IssueCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_coupon_v1_coupon_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCouponResponse.ProtoReflect.Descriptor instead.
func (*IssueCouponResponse) Descriptor() ([]byte, []int) {
	return file_proto_coupon_v1_coupon_proto_rawDescGZIP(), []int{1}
}

func (x *IssueCouponResponse) GetCouponId() int32 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *IssueCouponResponse) GetCouponCode() string {
	if x != nil {
		return x.CouponCode
	}
	return ""
}

var File_proto_coupon_v1_coupon_proto protoreflect.FileDescriptor

const file_proto_coupon_v1_coupon_proto_rawDesc = "" +
	"\n" +
	"\x1cproto/coupon/v1/coupon.proto\x12\tcoupon.v1\"5\n" +
	"\x12IssueCouponRequest\x12\x1f\n" +
	"\vcampaign_id\x18\x01 \x01(\x05R\n" +
	"campaignId\"S\n" +
	"\x13IssueCouponResponse\x12\x1b\n" +
	"\tcoupon_id\x18\x01 \x01(\x05R\bcouponId\x12\x1f\n" +
	"\vcoupon_code\x18\x02 \x01(\tR\n" +
	"couponCode2_\n" +
	"\rCouponService\x12N\n" +
	"\vIssueCoupon\x12\x1d.coupon.v1.IssueCouponRequest\x1a\x1e.coupon.v1.IssueCouponResponse\"\x00B\xaa\x01\n" +
	"\rcom.coupon.v1B\vCouponProtoP\x01ZGgithub.com/SuperRPM/coupon-issuance-system/gen/proto/coupon/v1;couponv1\xa2\x02\x03CXX\xaa\x02\tCoupon.V1\xca\x02\tCoupon\\V1\xe2\x02\x15Coupon\\V1\\GPBMetadata\xea\x02\n" +
	"Coupon::V1b\x06proto3"

var (
	file_proto_coupon_v1_coupon_proto_rawDescOnce sync.Once
	file_proto_coupon_v1_coupon_proto_rawDescData []byte
)

func file_proto_coupon_v1_coupon_proto_rawDescGZIP() []byte {
	file_proto_coupon_v1_coupon_proto_rawDescOnce.Do(func() {
		file_proto_coupon_v1_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_coupon_v1_coupon_proto_rawDesc), len(file_proto_coupon_v1_coupon_proto_rawDesc)))
	})
	return file_proto_coupon_v1_coupon_proto_rawDescData
}

var file_proto_coupon_v1_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_coupon_v1_coupon_proto_goTypes = []any{
	(*IssueCouponRequest)(nil),  // 0: coupon.v1.IssueCouponRequest
	(*IssueCouponResponse)(nil), // 1: coupon.v1.IssueCouponResponse
}
var file_proto_coupon_v1_coupon_proto_depIdxs = []int32{
	0, // 0: coupon.v1.CouponService.IssueCoupon:input_type -> coupon.v1.IssueCouponRequest
	1, // 1: coupon.v1.CouponService.IssueCoupon:output_type -> coupon.v1.IssueCouponResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_coupon_v1_coupon_proto_init() }
func file_proto_coupon_v1_coupon_proto_init() {
	if File_proto_coupon_v1_coupon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_coupon_v1_coupon_proto_rawDesc), len(file_proto_coupon_v1_coupon_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_coupon_v1_coupon_proto_goTypes,
		DependencyIndexes: file_proto_coupon_v1_coupon_proto_depIdxs,
		MessageInfos:      file_proto_coupon_v1_coupon_proto_msgTypes,
	}.Build()
	File_proto_coupon_v1_coupon_proto = out.File
	file_proto_coupon_v1_coupon_proto_goTypes = nil
	file_proto_coupon_v1_coupon_proto_depIdxs = nil
}
