package bytesutils

import (
	"errors"
	"fmt"
)

func IntToZigZag(n int) int {
	return (n << 1) ^ (n >> 31)
}

func SizeVarint(x uint64) int {
	switch {
	case x < 1<<7:
		return 1
	case x < 1<<14:
		return 2
	case x < 1<<21:
		return 3
	case x < 1<<28:
		return 4
	case x < 1<<35:
		return 5
	case x < 1<<42:
		return 6
	case x < 1<<49:
		return 7
	case x < 1<<56:
		return 8
	case x < 1<<63:
		return 9
	}
	return 10
}

var errBufNil = errors.New("buf must not be null")

func WriteVar(value int, buf []byte, offset int) (int, error) {
	if buf == nil {
		return 0, errBufNil
	}
	err := checkBound(len(buf), offset)
	if err != nil {
		return 0, err
	}
	for {
		if (value & ^0x7F) == 0 {
			buf[offset] = byte(value)
			offset++
			return offset, nil
		}
		// buf[offset] = (value & 0x7F) | 0x80
		buf[offset] = 0x80 | uint8(value&0x7F)
		offset++
		value = value >> 7
	}
}

const maxVarintBytes = 10 // maximum length of a varint
// EncodeVarint returns the varint encoding of x.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
// protocol buffer types.
// Not used by the package itself, but helpful to clients
// wishing to use the same encoding.
func EncodeVarint(x uint64) []byte {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

func checkBound(bufferLength, offset int) error {
	if offset < 0 {
		return fmt.Errorf("negative offset: %d", offset)
	}

	if offset >= bufferLength {
		return fmt.Errorf("invalid offset: %d, bufferLength: %d", offset, bufferLength)
	}

	return nil
}

// final byte[] buffer, int bufferOffset, final byte[] srcBytes
func WriteBytes(buffer, srcBytes []byte, bufferOffset int) int {
	for i := 0; i < len(srcBytes); i++ {
		buffer[bufferOffset+i] = srcBytes[i]
	}
	return len(srcBytes) + bufferOffset
}
