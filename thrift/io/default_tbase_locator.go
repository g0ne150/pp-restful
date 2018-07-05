package io

import (
	"errors"
	"pp-restful/thrift/dto/pinpoint"
	"pp-restful/thrift/dto/trace"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type defaultTBaseLocator struct {
	SPAN        uint16
	SPAN_HEADER *header

	AGENT_INFO        uint16
	AGENT_INFO_HEADER *header

	AGENT_STAT        uint16
	AGENT_STAT_HEADER *header

	SPANCHUNK        uint16
	SPANCHUNK_HEADER *header

	SPANEVENT        uint16
	SPANEVENT_HEADER *header

	SQLMETADATA        uint16
	SQLMETADATA_HEADER *header

	APIMETADATA        uint16
	APIMETADATA_HEADER *header

	RESULT        uint16
	RESULT_HEADER *header

	STRINGMETADATA        uint16
	STRINGMETADATA_HEADER *header

	CHUNK        uint16
	CHUNK_HEADER *header
}

func newDefaultTBaseLocator() *defaultTBaseLocator {
	newL := &defaultTBaseLocator{
		SPAN:           40,
		AGENT_INFO:     50,
		AGENT_STAT:     55,
		SPANCHUNK:      70,
		SPANEVENT:      80,
		SQLMETADATA:    300,
		APIMETADATA:    310,
		RESULT:         320,
		STRINGMETADATA: 330,
		// CHUNK:          400,
	}
	newL.SPAN_HEADER = createHeader(newL.SPAN)
	newL.AGENT_INFO_HEADER = createHeader(newL.AGENT_INFO)
	newL.AGENT_STAT_HEADER = createHeader(newL.AGENT_STAT)
	newL.SPANCHUNK_HEADER = createHeader(newL.SPANCHUNK)
	newL.SPANEVENT_HEADER = createHeader(newL.SPANEVENT)
	newL.SQLMETADATA_HEADER = createHeader(newL.SQLMETADATA)
	newL.APIMETADATA_HEADER = createHeader(newL.APIMETADATA)
	newL.RESULT_HEADER = createHeader(newL.RESULT)
	newL.STRINGMETADATA_HEADER = createHeader(newL.STRINGMETADATA)
	// newL.CHUNK_HEADER = createHeader(newL.CHUNK)
	return newL
}

func createHeader(hType uint16) *header {
	newHeader := newHeader()
	newHeader.hType = hType
	return newHeader
}

func (l defaultTBaseLocator) headerLookup(tbase thrift.TStruct) (*header, error) {

	if _, ok := tbase.(*trace.TSpan); ok {
		return l.SPAN_HEADER, nil
	}

	if _, ok := tbase.(*pinpoint.TAgentInfo); ok {
		return l.AGENT_INFO_HEADER, nil
	}

	if _, ok := tbase.(*trace.TSpanEvent); ok {
		return l.SPANEVENT_HEADER, nil
	}

	if _, ok := tbase.(*trace.TSpanChunk); ok {
		return l.SPANCHUNK_HEADER, nil
	}

	if _, ok := tbase.(*pinpoint.TAgentStat); ok {
		return l.AGENT_STAT_HEADER, nil
	}

	if _, ok := tbase.(*trace.TSqlMetaData); ok {
		return l.SQLMETADATA_HEADER, nil
	}

	if _, ok := tbase.(*trace.TApiMetaData); ok {
		return l.APIMETADATA_HEADER, nil
	}

	if _, ok := tbase.(*trace.TResult_); ok {
		return l.RESULT_HEADER, nil
	}

	if _, ok := tbase.(*trace.TStringMetaData); ok {
		return l.STRINGMETADATA_HEADER, nil
	}

	return nil, errors.New("invalid tbase")

}
