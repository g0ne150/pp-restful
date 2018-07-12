package commons

import (
	"errors"
	"pp-restful/commons/bytesutils"
	"strconv"
	"strings"
)

const (
	TransactionIDDelimiter string = "^"
	Version                byte   = 0
	versionSize            int    = 1
)

func ParseTransactionID(tranID string) (*TransactionID, error) {
	if tranID == "" {
		return nil, errors.New("tranID string must not be empty")
	}

	res := strings.Split(tranID, TransactionIDDelimiter)

	if len(res) != 3 {
		return nil, errors.New("wrong tranID")
	}

	agentID := res[0]
	agentStartTime, _ := strconv.Atoi(res[1])
	transactionSequence, _ := strconv.Atoi(res[2])

	return &TransactionID{
		agentID:             agentID,
		agentStartTime:      agentStartTime,
		transactionSequence: transactionSequence,
	}, nil
}

func WriteTransactionID(agentID string, agentStartTime, transactionSequence int) []byte {
	agentIDBytes := []byte(agentID)
	agentIDLength := len(agentIDBytes)
	zigZagAgentIDLength := bytesutils.IntToZigZag(agentIDLength)
	agentIDPrefixSize := bytesutils.SizeVarint(uint64(zigZagAgentIDLength))
	agentStartTimeSize := bytesutils.SizeVarint(uint64(agentStartTime))
	transactionIDSequenceSize := bytesutils.SizeVarint(uint64(transactionSequence))

	bufferSize := versionSize + agentIDPrefixSize + len(agentIDBytes) + agentStartTimeSize + transactionIDSequenceSize

	buffer := make([]byte, bufferSize)
	buffer[0] = Version
	offset := versionSize
	offset, _ = bytesutils.WriteVar(zigZagAgentIDLength, buffer, offset)
	if agentIDBytes != nil {
		offset = bytesutils.WriteBytes(buffer, agentIDBytes, offset)
	}
	offset, _ = bytesutils.WriteVar(agentStartTime, buffer, offset)
	bytesutils.WriteVar(transactionSequence, buffer, offset)

	return buffer
}
