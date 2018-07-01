package io

// type HeaderTBaseSerializer struct {
// 	protocol *thrift.TProtocol
// 	locator  *thrift.TBaseLocator
// }

const (
	sIGNATURE   int8 = -17
	hEADER_SIZE      = 4
)

type header struct {
	signature int8
	version   int8
	hType     uint16
}

func newHeader() *header {
	return &header{
		hType:     0,
		signature: sIGNATURE,
		version:   16,
	}
}
