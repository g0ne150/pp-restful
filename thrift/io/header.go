package io

// type HeaderTBaseSerializer struct {
// 	protocol *thrift.TProtocol
// 	locator  *thrift.TBaseLocator
// }

const (
	sIGNATURE   byte = 0xef
	hEADER_SIZE      = 4
)

type header struct {
	signature byte
	version   byte
	hType     uint16
}

func newHeader() *header {
	return &header{
		hType:     0,
		signature: sIGNATURE,
		version:   0x10,
	}
}
