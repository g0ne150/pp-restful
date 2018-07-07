package io

import (
	"encoding/binary"
	"io/ioutil"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Serializer struct {
	locator    *defaultTBaseLocator
	tTransport *thrift.TMemoryBuffer
	protocol   *thrift.TCompactProtocol
}

func NewSerializer() *Serializer {
	sr := &Serializer{
		locator:    newDefaultTBaseLocator(),
		tTransport: thrift.NewTMemoryBuffer(),
	}
	sr.protocol = thrift.NewTCompactProtocol(sr.tTransport)

	return sr
}

func (s Serializer) Serialize(tbase thrift.TStruct) ([]byte, error) {
	s.tTransport.Reset()

	header, err := s.locator.headerLookup(tbase)
	if err != nil {
		return nil, err
	}

	err = tbase.Write(s.protocol)
	if err != nil {
		return nil, err
	}

	// err = s.tTransport.Flush()
	err = s.tTransport.Flush()
	if err != nil {
		return nil, err
	}

	headerbuff := headerSerialize(header)
	bodyBuff, _ := ioutil.ReadAll(s.tTransport)

	return append(headerbuff, bodyBuff...), nil
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
