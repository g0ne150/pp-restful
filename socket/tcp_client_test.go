package socket

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"pp-restful/commons"
	"pp-restful/thrift/dto/trace"
	"pp-restful/thrift/io"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	conn, err := net.Dial("tcp", "10.241.12.209:9996")
	if err != nil {
		t.Fatal("net.Dial error: ", err.Error())
	}
	defer conn.Close()

	rpc := "/"
	endPoint := "localhost:8080"
	remoteAddr := "127.0.0.1"
	var applicationServiceType int16 = 1210
	var apiId int32 = 3

	tranID, _ := commons.ParseTransactionID("shark-test^1534496788335^14")

	tSpan := trace.NewTSpan()
	tSpan.AgentId = "shark-test"
	tSpan.ApplicationName = "shark-test"
	tSpan.AgentStartTime = 1534496788335
	tSpan.TransactionId = commons.WriteTransactionID(tranID.GetAgentID(), tranID.GetAgentStartTime(), tranID.GetTransactionSequence())
	tSpan.SpanId = -9076112736953707450
	tSpan.ParentSpanId = -1
	tSpan.StartTime = time.Now().Unix()
	tSpan.Elapsed = 9
	tSpan.RPC = &rpc
	tSpan.ServiceType = 1010
	tSpan.EndPoint = &endPoint
	tSpan.RemoteAddr = &remoteAddr
	tSpan.ApplicationServiceType = &applicationServiceType
	tSpan.ApiId = &apiId

	sendData, err := io.Serialize(tSpan)
	if err != nil {
		t.Fatal("serialize tSpan failed: ", err.Error())
	}

	fmt.Println(hex.EncodeToString(sendData))

	// h, _ := hex.DecodeString("000100000104")

	h := make([]byte, 2+4)
	binary.BigEndian.PutUint16(h[:2], 1)
	binary.BigEndian.PutUint32(h[2:], uint32(len(sendData)))
	sendData = append(h, sendData...)
	_, err = conn.Write(sendData)
	if err != nil {
		t.Fatal("conn.Write data error: ", err.Error())
	}

	// res, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	t.Fatal("read from conection error: ", err.Error())
	// }

	// fmt.Println(bytes.NewBuffer(res).String())
}
