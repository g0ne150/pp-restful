package io_test

import (
	"fmt"
	"pp-restful/thrift/dto/trace"
	"pp-restful/thrift/io"
	"testing"
)

func TestSerialize(t *testing.T) {
	tSpan := trace.NewTSpan()
	tSpan.ApplicationName = "application name"
	tSpan.AgentId = "agent id"

	buf, err := io.Serialize(tSpan)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("--------------->", buf)

}

func TestBinary(t *testing.T) {
	var s int8 = io.SIGNATURE
	var v int8 = 16
	var hType uint16 = 40

	fmt.Println(byte(s))
	fmt.Println(byte(v))
	fmt.Println(byte(hType))
}
