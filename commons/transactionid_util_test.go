package commons

import (
	"fmt"
	"testing"
)

func TestWriteTransactionID(t *testing.T) {
	// now := time.Now().Unix()
	now := 1531375438037
	tansactionID := WriteTransactionID("test", int(now), 2)

	fmt.Println(tansactionID)

	for i, v := range []byte{0, 8, 116, 101, 115, 116, 213, 145, 219, 232, 200, 44, 2} {
		if v != tansactionID[i] {
			t.Fatal("write transaction failed")
		}
	}
}
