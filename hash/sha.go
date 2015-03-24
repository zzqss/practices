package main

import (
	"encoding/binary"
)

const (
	H0 uint32 = 0x67452301
	H1 uint32 = 0xEFCDAB89
	H2 uint32 = 0x98BADCFE
	H3 uint32 = 0x10325476
	H4 uint32 = 0xC3D2E1F0
)

// make up message
// append the bit '1' to the message
// append k bits '0', where k is the minimum number >= 0 such that the resulting message
//    length (in bits) is congruent to 448(mod 512)
// append length of message (before pre-processing), in bits, as 64-bit big-endian integer
func makeUpMessage(oldMessage []byte) []byte {

}

// Sha1 encode message
func Sha1(message []byte) []byte {

}

func main() {

}
