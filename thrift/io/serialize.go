package io

import (
	"bytes"
	"encoding/binary"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type serializer struct {
	locator *DefaultTBaseLocator
	header  *Header
}

var pri serializer

func Serialize(tbase thrift.TStruct) ([]byte, error) {
	var err error
	pri.locator = NewDefaultTBaseLocator()
	pri.header, err = pri.locator.HeaderLookup(tbase)
	if err != nil {
		return nil, err
	}

	tTransport := thrift.NewTMemoryBuffer()

	transport := thrift.NewTFramedTransport(tTransport)
	protocol := thrift.NewTCompactProtocol(transport)

	tbase.Write(protocol)
	transport.Flush()

	headerbuff := HeaderSerialize(pri.header)

	headerbuff.Write(tTransport.Buffer.Bytes())

	return headerbuff.Bytes(), nil
}

func HeaderSerialize(header *Header) *bytes.Buffer {
	buff := make([]byte, 4)

	buff[0] = byte(header.signature)
	buff[1] = byte(header.version)
	binary.BigEndian.PutUint16(buff[2:], header.hType)

	return bytes.NewBuffer(buff)
}

func getLocator() *DefaultTBaseLocator {
	return pri.locator
}

func getHeader() *Header {
	return pri.header
}
