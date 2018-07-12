package commons

import (
	"errors"
)

type TransactionID struct {
	agentID             string
	agentStartTime      int
	transactionSequence int
}

func NewTransactionID(agentID string, agentStartTime, transactionSequence int) (*TransactionID, error) {
	if agentID == "" {
		return nil, errors.New("agentID must not be null")
	}

	return &TransactionID{
		agentID:             agentID,
		agentStartTime:      agentStartTime,
		transactionSequence: transactionSequence,
	}, nil
}

func (t *TransactionID) GetAgentID() string {
	return t.agentID
}

func (t *TransactionID) GetAgentStartTime() int {
	return t.agentStartTime
}

func (t *TransactionID) GetTransactionSequence() int {
	return t.transactionSequence
}
