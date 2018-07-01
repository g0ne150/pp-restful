package io

import (
	"bytes"
	"encoding/binary"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Serializer struct {
	locator *defaultTBaseLocator
	// header           *Header
	tTransport       *thrift.TMemoryBuffer
	tFramedTransport *thrift.TFramedTransport
	protocol         *thrift.TCompactProtocol
}

func NewSerializer() *Serializer {
	sr := &Serializer{
		locator:    newDefaultTBaseLocator(),
		tTransport: thrift.NewTMemoryBuffer(),
	}
	sr.tFramedTransport = thrift.NewTFramedTransport(sr.tTransport)
	sr.protocol = thrift.NewTCompactProtocol(sr.tFramedTransport)

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
	err = s.tFramedTransport.Flush()
	// err = s.tTransport.Flush()

	if err != nil {
		return nil, err
	}

	headerbuff := headerSerialize(header)
	_, err = headerbuff.Write(s.tTransport.Buffer.Bytes())
	if err != nil {
		return nil, err
	}

	return headerbuff.Bytes(), nil
}

func headerSerialize(header *header) *bytes.Buffer {
	buff := make([]byte, 4)

	buff[0] = byte(header.signature)
	buff[1] = byte(header.version)
	binary.BigEndian.PutUint16(buff[2:], header.hType)

	return bytes.NewBuffer(buff)
}

/*
func getLocator() *DefaultTBaseLocator {
	return pri.locator
}

func getHeader() *header {
	return pri.header
}
*/
