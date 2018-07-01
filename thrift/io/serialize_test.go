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

	// FIXME body buffer 的序列看起对的，与 nodejs 实现有差异，猜测是 go 的零值导致的，感觉无碍，待查
	// [0 0 0 39 24 8 97 103 101 110 116 32 105 100 24 16 97 112 112 108 105 99 97 116 105 111 110 32 110 97 109 101 22 0 24 0 54 0 38 0 52 0 0]
	// nodejs: [ 24 8 97 103 101 110 116 32 105 100 24 16 97 112 112 108 105 99 97 116 105 111 110 32 110 97 109 101 102 1 37 0 100 0 0 ]
	fmt.Println("--------------->", buf)

}
