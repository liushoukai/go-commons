package jsonutil

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoToJSON 将 protobuf 消息转换为 JSON
func ProtoToJSON(message proto.Message) (string, error) {
	// 创建 marshaler
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   false, // false 使用 json 标签名称（如 accountId），true 使用原始字段名（如 account_id）
		EmitUnpopulated: true,  // 是否包含零值字段
	}
	// 转换为 JSON
	b, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// JsonToProto 将 JSON 转换回 protobuf 消息
func JsonToProto(data string, message proto.Message) error {
	// 创建 unmarshaler
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true, // 是否忽略未知字段
	}
	// 转换回 protobuf
	return unmarshaler.Unmarshal([]byte(data), message)
}
