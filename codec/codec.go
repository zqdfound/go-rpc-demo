package codec

import "io"

// RPC 请求头
type Header struct {
	ServiceMethod string //服务名和方法名，通常与 Go 语言中的结构体和方法相映射
	Seq           uint64 //请求ID
	Error         string
}

type Codec interface {
	io.Closer
	ReadHeader(header *Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

// Codec 的构造函数,客户端和服务端可以通过 Codec 的 Type 得到构造函数，从而创建 Codec 实例
type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
