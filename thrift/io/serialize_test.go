package io

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"pp-restful/thrift/dto/trace"

	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func TestSerialize(t *testing.T) {
	tSpan := trace.NewTSpan()
	tSpan.ApplicationName = "application name"
	tSpan.AgentId = "agent id"

	buf, err := Serialize(tSpan)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("--------------->", hex.EncodeToString(buf))

	if len(buf) == 0 {
		t.Errorf("buf length is zero")
	}

	// deserialize
	deTSpan := trace.NewTSpan()
	deserializer := thrift.NewTDeserializer()
	transport := thrift.NewTMemoryBuffer()
	transport.Buffer = bytes.NewBuffer(buf[4:])

	deserializer.Protocol = thrift.NewTCompactProtocol(transport)

	buffer := []byte{}
	err = deserializer.Read(deTSpan, buffer)
	if err != nil {
		t.Errorf("deserialize failed: %s", err.Error())
	}

	fmt.Println(deTSpan)

	if deTSpan.ApplicationName != tSpan.ApplicationName || deTSpan.AgentId != tSpan.AgentId {
		t.Errorf("deserialize error")
	}

}
