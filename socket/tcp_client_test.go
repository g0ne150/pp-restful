package socket

import (
	"encoding/hex"
	"fmt"
	"pp-restful/thrift/dto/trace"
	"pp-restful/thrift/io"
	"testing"
)

func TestConnect(t *testing.T) {
	// conn, err := net.Dial("tcp", "10.241.12.209:9996")
	// if err != nil {
	// 	t.Fatal("net.Dial error: ", err.Error())
	// }
	// defer conn.Close()

	// rpc := "/"
	// stringV := "Error: 304 Not Modified"
	// var applicationServiceType int16 = 100

	tSpan := trace.NewTSpan()
	tSpan.AgentId = "node-agent-test"
	tSpan.ApplicationName = "node-agent-test"

	serializer := io.NewSerializer()

	sendData, err := serializer.Serialize(tSpan)
	if err != nil {
		t.Fatal("serialize tSpan failed: ", err.Error())
	}

	fmt.Println(hex.EncodeToString(sendData))

	// go func() {
	// 	_, err = conn.Write(sendData)
	// 	if err != nil {
	// 		t.Fatal("conn.Write data error: ", err.Error())
	// 	}
	// }()

	// res, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	t.Fatal("read from conection error: ", err.Error())
	// }

	// fmt.Println(bytes.NewBuffer(res).String())
}
