package io

import (
	"errors"
	"pp-restful/thrift/dto/trace"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func Serialize(tbase trace.TSpan) (err error) {
	tansportFactory := thrift.NewTTransportFactory()

	var transport thrift.TTransport

	transport, err = tansportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	TFtransport := thrift.NewTFramedTransport(transport)
	TCprotocol := thrift.NewTCompactProtocol(transport)
	tbase.Write(TCprotocol)

	// return bytes.NewBuffer()

	return errors.New("// TODO")
}
