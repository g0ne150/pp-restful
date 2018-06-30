package io

import (
	"bytes"
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

	headerbuff.Write(tTransport.Buffer.Bytes())

	return headerbuff.Bytes(), nil
}

func HeaderSerialize(header *Header) *bytes.Buffer {
	buff := bytes.NewBuffer(make([]byte, 4, 4))

	// FIXME 二进制拼接还是有问题
	buff.WriteByte(byte(header.signature))
	buff.WriteByte(byte(header.version))
	buff.WriteByte(byte(header.hType))

	return buff
}

func getLocator() *DefaultTBaseLocator {
	return pri.locator
}

func getHeader() *Header {
	return pri.header
}
