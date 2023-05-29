// Configuration proto for Kubernetes provider.
//
// Example provider config:
// {
//   pods {}
// }
//
// In probe config:
// probe {
//   targets{
//     rds_targets {
//       resource_path: "k8s://pods"
//       filter {
//         key: "namespace"
//         value: "default"
//       }
//       filter {
//         key: "name"
//         value: "cloudprober.*"
//       }
//     }
//   }
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/rds/file/proto/config.proto

package proto

import (
	proto "github.com/cloudprober/cloudprober/rds/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProviderConfig_Format int32

const (
	ProviderConfig_UNSPECIFIED ProviderConfig_Format = 0 // Determine format using file extension/
	ProviderConfig_TEXTPB      ProviderConfig_Format = 1 // Text proto format (.textpb).
	ProviderConfig_JSON        ProviderConfig_Format = 2 // JSON proto format (.json).
)

// Enum value maps for ProviderConfig_Format.
var (
	ProviderConfig_Format_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "TEXTPB",
		2: "JSON",
	}
	ProviderConfig_Format_value = map[string]int32{
		"UNSPECIFIED": 0,
		"TEXTPB":      1,
		"JSON":        2,
	}
)

func (x ProviderConfig_Format) Enum() *ProviderConfig_Format {
	p := new(ProviderConfig_Format)
	*p = x
	return p
}

func (x ProviderConfig_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderConfig_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_enumTypes[0].Descriptor()
}

func (ProviderConfig_Format) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_enumTypes[0]
}

func (x ProviderConfig_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProviderConfig_Format) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProviderConfig_Format(num)
	return nil
}

// Deprecated: Use ProviderConfig_Format.Descriptor instead.
func (ProviderConfig_Format) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescGZIP(), []int{0, 0}
}

// File provider config.
type ProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// File that contains resources in either textproto or json format. File can
	// be local, on GCS, on S3, or any HTTP(S) URL.
	// e.g.:
	//   - /tmp/resources.textpb
	//   - gs://my-bucket/resources.json
	//   - s3://my-bucket/resources.json
	//   - https://my-public-bucket.s3.amazonaws.com/resources.json
	//
	// Example in textproto format:
	//
	//	resource {
	//	  name: "switch-xx-01"
	//	  ip: "10.11.112.3"
	//	  port: 8080
	//	  labels {
	//	    key: "device_type"
	//	    value: "switch"
	//	  }
	//	}
	//
	//	resource {
	//	  name: "switch-yy-01"
	//	  ip: "10.16.110.12"
	//	  port: 8080
	//	}
	FilePath []string               `protobuf:"bytes,1,rep,name=file_path,json=filePath" json:"file_path,omitempty"`
	Format   *ProviderConfig_Format `protobuf:"varint,2,opt,name=format,enum=cloudprober.rds.file.ProviderConfig_Format" json:"format,omitempty"`
	// If specified, file will be re-read at the given interval.
	ReEvalSec *int32 `protobuf:"varint,3,opt,name=re_eval_sec,json=reEvalSec" json:"re_eval_sec,omitempty"`
	// Whenever possible, we reload a file only if it has been modified since the
	// last load. If following option is set, mod time check is disabled.
	// Note that mod-time check doesn't work for GCS.
	DisableModifiedTimeCheck *bool `protobuf:"varint,4,opt,name=disable_modified_time_check,json=disableModifiedTimeCheck" json:"disable_modified_time_check,omitempty"`
}

func (x *ProviderConfig) Reset() {
	*x = ProviderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderConfig) ProtoMessage() {}

func (x *ProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderConfig.ProtoReflect.Descriptor instead.
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderConfig) GetFilePath() []string {
	if x != nil {
		return x.FilePath
	}
	return nil
}

func (x *ProviderConfig) GetFormat() ProviderConfig_Format {
	if x != nil && x.Format != nil {
		return *x.Format
	}
	return ProviderConfig_UNSPECIFIED
}

func (x *ProviderConfig) GetReEvalSec() int32 {
	if x != nil && x.ReEvalSec != nil {
		return *x.ReEvalSec
	}
	return 0
}

func (x *ProviderConfig) GetDisableModifiedTimeCheck() bool {
	if x != nil && x.DisableModifiedTimeCheck != nil {
		return *x.DisableModifiedTimeCheck
	}
	return false
}

type FileResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource []*proto.Resource `protobuf:"bytes,1,rep,name=resource" json:"resource,omitempty"`
}

func (x *FileResources) Reset() {
	*x = FileResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResources) ProtoMessage() {}

func (x *FileResources) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResources.ProtoReflect.Descriptor instead.
func (*FileResources) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *FileResources) GetResource() []*proto.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_github_com_cloudprober_cloudprober_rds_file_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82,
	0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x43,
	0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x72, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x45, 0x76, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x1b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x22, 0x2f, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x45, 0x58, 0x54, 0x50, 0x42, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f,
	0x4e, 0x10, 0x02, 0x22, 0x46, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x72, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2f, 0x72, 0x64, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_goTypes = []interface{}{
	(ProviderConfig_Format)(0), // 0: cloudprober.rds.file.ProviderConfig.Format
	(*ProviderConfig)(nil),     // 1: cloudprober.rds.file.ProviderConfig
	(*FileResources)(nil),      // 2: cloudprober.rds.file.FileResources
	(*proto.Resource)(nil),     // 3: cloudprober.rds.Resource
}
var file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.rds.file.ProviderConfig.format:type_name -> cloudprober.rds.file.ProviderConfig.Format
	3, // 1: cloudprober.rds.file.FileResources.resource:type_name -> cloudprober.rds.Resource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_rds_file_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_rds_file_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_rds_file_proto_config_proto_depIdxs = nil
}
