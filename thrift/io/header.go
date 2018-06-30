package io

// type HeaderTBaseSerializer struct {
// 	protocol *thrift.TProtocol
// 	locator  *thrift.TBaseLocator
// }

const (
	SIGNATURE   int8 = -17
	HEADER_SIZE      = 4
)

type Header struct {
	signature int8
	version   int8
	hType     uint16
}

func NewHeader() *Header {
	return &Header{
		hType:     0,
		signature: SIGNATURE,
		version:   16,
	}
}
