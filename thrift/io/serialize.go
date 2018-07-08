package io

import (
	"encoding/binary"
	"io/ioutil"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var locator *defaultTBaseLocator

func init() {
	locator = newDefaultTBaseLocator()
}

func Serialize(tbase thrift.TStruct) ([]byte, error) {
	transport := thrift.NewTMemoryBuffer()
	protocol := thrift.NewTCompactProtocol(transport)

	header, err := locator.headerLookup(tbase)
	if err != nil {
		return nil, err
	}

	_, err = transport.Write(headerSerialize(header))
	if err != nil {
		return nil, err
	}

	err = tbase.Write(protocol)
	if err != nil {
		return nil, err
	}

	// transport.Flush()

	result, err := ioutil.ReadAll(transport)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func headerSerialize(header *header) []byte {
	buff := make([]byte, 4)

	buff[0] = header.signature
	buff[1] = header.version
	binary.BigEndian.PutUint16(buff[2:], header.hType)

	return buff
}

/*
func getLocator() *DefaultTBaseLocator {
	return pri.locator
}

func getHeader() *header {
	return pri.header
}
*/
