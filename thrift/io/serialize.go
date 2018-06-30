package io

import (
	"bytes"
	"encoding/binary"
	"fmt"

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

	fmt.Println(headerbuff.Bytes())
	fmt.Println(tTransport.Buffer.Bytes())

	// FIXME 两个 buffer 拼接还有问题
	headerbuff.Write(tTransport.Buffer.Bytes())

	return headerbuff.Bytes(), nil
}

func HeaderSerialize(header *Header) *bytes.Buffer {
	buff := make([]byte, 4, 4)

	// FIXME 二进制拼接还是有问题
	binary.PutVarint(buff[:1], int64(header.signature))
	binary.PutVarint(buff[1:2], int64(header.version))
	binary.BigEndian.PutUint16(buff[2:], header.hType)

	return bytes.NewBuffer(buff)
}

func getLocator() *DefaultTBaseLocator {
	return pri.locator
}

func getHeader() *Header {
	return pri.header
}
