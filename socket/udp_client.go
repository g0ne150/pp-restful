package socket

import (
	"errors"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Serializer interface {
	Serialize(tbase thrift.TStruct) ([]byte, error)
}

type UDPClient struct {
	host       string
	port       string
	serializer Serializer
}

func NewUDPClient(host, port string) *UDPClient {
	return &UDPClient{
		host: host,
		port: port,
	}
}

func (c *UDPClient) SetHost(host string) {
	c.host = host
}

func (c *UDPClient) SetPort(port string) {
	c.port = port
}

func (c *UDPClient) SetSerializer(serializer Serializer) {
	c.serializer = serializer
}

func (c *UDPClient) send(msg thrift.TStruct) (err error) {
	if c.serializer == nil {
		return errors.New("Set serializer first before call UDPClient.send() please")
	}

	var serializedMsg []byte
	serializedMsg, err = c.serializer.Serialize(msg)
	if err != nil {
		return err
	}

	// TODO implement UDP socket and send serialized data

	return nil
}
