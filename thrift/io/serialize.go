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

	// FIXME body buffer 的序列看起对的，比 nodejs 实现前面多了 0 0 0 39 猜测是 go 的零值导致的，感觉无碍，待查
	// [0 0 0 39 24 8 97 103 101 110 116 32 105 100 24 16 97 112 112 108 105 99 97 116 105 111 110 32 110 97 109 101 22 0 24 0 54 0 38 0 52 0 0]
	// nodejs: [ 24 8 97 103 101 110 116 32 105 100 24 16 97 112 112 108 105 99 97 116 105 111 110 32 110 97 109 101 102 1 37 0 100 0 0 ]
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
