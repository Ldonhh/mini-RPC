package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	SeqNumber     uint64 // sequence number called by client
	Error         string // save error
}

// define Codec interface
// in order to implement different encoder/decoder instance

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc construct function
type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GodType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodeFuncMap map[Type]NewCodecFunc

// similar to factory, but return a function instead of an instance

func init() {
	NewCodeFuncMap = make(map[Type]NewCodecFunc)
	NewCodeFuncMap[GodType] = NewGobCodec
}
