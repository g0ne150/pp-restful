package io

import (
	"encoding/hex"
	"fmt"
	"pp-restful/thrift/dto/trace"

	"testing"
)

func TestSerialize(t *testing.T) {
	tSpan := trace.NewTSpan()
	tSpan.ApplicationName = "application name"
	tSpan.AgentId = "agent id"

	serializer := NewSerializer()
	buf, err := serializer.Serialize(tSpan)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("--------------->", hex.EncodeToString(buf))

	if len(buf) == 0 {
		t.Errorf("buf length is zero")
	}

	// if buf[0] != 239 || buf[1] != 16 || buf[3] != 40 {
	// 	t.Errorf("header buffer error")
	// }

	// deTSpan := trace.NewTSpan()
	// deserializer := thrift.NewTDeserializer()
	// tTransport := thrift.NewTMemoryBuffer()
	// tTransport.Buffer = bytes.NewBuffer(buf[4:])
	// transport := thrift.NewTFramedTransport(tTransport)

	// deserializer.Protocol = thrift.NewTCompactProtocol(transport)

	// buffer := []byte{}
	// err = deserializer.Read(deTSpan, buffer)
	// if err != nil {
	// 	t.Errorf("deserialize failed: %s", err.Error())
	// }

	// fmt.Println(deTSpan)

	// if deTSpan.ApplicationName != tSpan.ApplicationName || deTSpan.AgentId != tSpan.AgentId {
	// 	t.Errorf("deserialize error")
	// }

}
