// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.

package modbus

import (
	"testing"
)

func TestCRC(t *testing.T) {
	var crc crc
	reset(&crc)
	pushBytes(&crc, []byte{0x02, 0x07})
	//crc.reset()
	//crc.pushBytes([]byte{0x02, 0x07})
	if 0x1241 != value(&crc) {
		t.Fatalf("crc expected %v, actual %v", 0x1241, crc.value())
	}
}
