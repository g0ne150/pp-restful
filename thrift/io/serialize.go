package io

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Serializer struct {
	locator *defaultTBaseLocator
	// header           *Header
	tTransport *thrift.StreamTransport
	protocol   *thrift.TCompactProtocol
}

func NewSerializer() *Serializer {
	sr := &Serializer{
		locator: newDefaultTBaseLocator(),
		// tTransport: thrift.NewTMemoryBuffer(),
	}
	// sr.tFramedTransport = thrift.NewTFramedTransport(sr.tTransport)
	// sr.tTransport = thrift.NewStreamTransportFactory().GetTransport()
	// sr.protocol = thrift.NewTCompactProtocol(sr.tTransport)

	return sr
}

func (s Serializer) Serialize(tbase thrift.TStruct) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	s.tTransport = thrift.NewStreamTransport(buf, buf)
	s.protocol = thrift.NewTCompactProtocol(s.tTransport)

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
