// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.10.0
// source: pkg/protocol/message.proto

package protocol

import (
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar         string `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          // 头像
	SenderUsername string `protobuf:"bytes,2,opt,name=sender_username,json=senderUsername,proto3" json:"sender_username,omitempty"`    // 发送消息用户的用户名
	SenderUserId   int64  `protobuf:"varint,3,opt,name=sender_user_id,json=senderUserId,proto3" json:"sender_user_id,omitempty"`       // 发送消息用户uuid
	ReceiverUserId int64  `protobuf:"varint,4,opt,name=receiver_user_id,json=receiverUserId,proto3" json:"receiver_user_id,omitempty"` // 发送给对端用户的uuid
	Content        string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`                                        // 文本消息内容
	ContentType    int64  `protobuf:"varint,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`            // 消息内容类型：1.文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天
	Type           string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`                                              // 消息传输类型：如果是心跳消息，该内容为heat beat,在线视频或者音频为webrtc
	MessageType    int64  `protobuf:"varint,8,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`            // 消息类型，1.单聊 2.群聊
	Url            string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`                                                // 图片，视频，语音的路径
	FileSuffix     string `protobuf:"bytes,10,opt,name=file_suffix,json=fileSuffix,proto3" json:"file_suffix,omitempty"`               // 文件后缀，如果通过二进制头不能解析文件后缀，使用该后缀
	File           []byte `protobuf:"bytes,11,opt,name=file,proto3" json:"file,omitempty"`                                             // 如果是图片，文件，视频等的二进制
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protocol_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protocol_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_protocol_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *Message) GetSenderUsername() string {
	if x != nil {
		return x.SenderUsername
	}
	return ""
}

func (x *Message) GetSenderUserId() int64 {
	if x != nil {
		return x.SenderUserId
	}
	return 0
}

func (x *Message) GetReceiverUserId() int64 {
	if x != nil {
		return x.ReceiverUserId
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetContentType() int64 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetMessageType() int64 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *Message) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Message) GetFileSuffix() string {
	if x != nil {
		return x.FileSuffix
	}
	return ""
}

func (x *Message) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_pkg_protocol_message_proto protoreflect.FileDescriptor

var file_pkg_protocol_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xd5, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protocol_message_proto_rawDescOnce sync.Once
	file_pkg_protocol_message_proto_rawDescData = file_pkg_protocol_message_proto_rawDesc
)

func file_pkg_protocol_message_proto_rawDescGZIP() []byte {
	file_pkg_protocol_message_proto_rawDescOnce.Do(func() {
		file_pkg_protocol_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protocol_message_proto_rawDescData)
	})
	return file_pkg_protocol_message_proto_rawDescData
}

var file_pkg_protocol_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_protocol_message_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: protocol.Message
}
var file_pkg_protocol_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_protocol_message_proto_init() }
func file_pkg_protocol_message_proto_init() {
	if File_pkg_protocol_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protocol_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_pkg_protocol_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protocol_message_proto_goTypes,
		DependencyIndexes: file_pkg_protocol_message_proto_depIdxs,
		MessageInfos:      file_pkg_protocol_message_proto_msgTypes,
	}.Build()
	File_pkg_protocol_message_proto = out.File
	file_pkg_protocol_message_proto_rawDesc = nil
	file_pkg_protocol_message_proto_goTypes = nil
	file_pkg_protocol_message_proto_depIdxs = nil
}
