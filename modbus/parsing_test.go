package modbus

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	fmt.Println(([]byte{0x02, 0x07}))
}
