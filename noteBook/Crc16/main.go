package main

import (
	"fmt"
	"github.com/howeyc/crc16"
)

/*
crc16  算法有16bit,可容纳2^16=65535种可能超过该数字存在重复可能
*/

func main() {
	data := []byte("中华人民A&002")
	checksum := crc16.Checksum(data, crc16.IBMTable)

	checksumIBM := crc16.ChecksumIBM(data)
	bus := crc16.ChecksumMBus(data)

	ccitt := crc16.ChecksumCCITT(data)

	fmt.Println(checksum, "\n", "\n", checksumIBM, bus, "\n", ccitt)
}
