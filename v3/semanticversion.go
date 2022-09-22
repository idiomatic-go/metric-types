package v3

import "google.golang.org/protobuf/runtime/protoimpl"

// Envoy uses SemVer (https://semver.org/). Major/minor versions indicate
// expected behaviors and APIs, the patch version field is used only
// for security fixes and can be generally ignored.
type SemanticVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MajorNumber uint32 `protobuf:"varint,1,opt,name=major_number,json=majorNumber,proto3" json:"major_number,omitempty"`
	MinorNumber uint32 `protobuf:"varint,2,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Patch       uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}
