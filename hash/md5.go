package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

/*
*	md5(Message-Digest Algorithm 5)
*	md5 is little-endian
**/

// some init number
const (
	H1 uint32 = 0x67452301
	H2 uint32 = 0xEFCDAB89
	H3 uint32 = 0x98BADCFE
	H4 uint32 = 0x10325476
)

type MD5 struct {
	Bytes []byte
}

var r [64]uint32 = [64]uint32{7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21}

//for i:=0;i<64;i++{
//    k[i] := floor(abs(sin(i + 1)) × 2^32)
//}
var k [64]uint32 = [64]uint32{0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee, 0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501, 0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be, 0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821, 0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa, 0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8, 0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed, 0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a, 0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c, 0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70, 0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05, 0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665, 0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039, 0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1, 0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1, 0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391}

// first is make up message
// for bit level,there are three step
// *	append "1" bit to message
// *	append "0" bits until message length in bits ≡ 448 (mod 512)
// *	append bit length of message as 64-bit little-endian integer to message
// because byte is the lowest level in go lang,so we can translate it to byte level:
// *	append byte(0x80) to the end of []byte,no matter little-endian or big-endian
// *	append byte(0x00) to []byte entil len([]byte) % 64 == 56
// *	append bit length as int64 in little-endian to message
func MakeUpBytes(oldBytes []byte) []byte {
	newBytes := make([]byte, len(oldBytes))
	copy(newBytes, oldBytes)
	var endSign byte = 0x80
	var messageLength uint64 = uint64(len(newBytes)) * 8
	littleEndianBuf := new(bytes.Buffer)
	err := binary.Write(littleEndianBuf, binary.LittleEndian, messageLength)
	if err != nil {
		log.Println(err)
		return nil
	}
	newBytes = append(newBytes, endSign)
	var nowLength uint = uint(len(newBytes))
	modLength := nowLength % 64
	var fixLength uint
	if modLength <= 56 {
		fixLength = 56 - modLength
	} else {
		fixLength = 56 + 64 - modLength
	}
	var i uint = 0
	for ; i < fixLength; i++ {
		newBytes = append(newBytes, 0x00)
	}
	newBytes = append(newBytes, littleEndianBuf.Bytes()...)
	return newBytes
}

//  transport big-endian number to litter-endian number
func TransportEndian(old []byte) uint32 {
	buf := bytes.NewBuffer(old)
	var newInt uint32
	binary.Read(buf, binary.LittleEndian, &newInt)
	return newInt
}

// transport []byte to litter-endian []uint32
func TransportEndianBatch(oldBytes []byte) []uint32 {
	byteLength := len(oldBytes)
	intSlice := make([]uint32, byteLength/4)
	for i := 0; i < byteLength; i += 4 {
		old := oldBytes[i : i+4]
		newInt := TransportEndian(old)
		intSlice[i/4] = newInt
	}
	return intSlice
}

// change uint32 to []byte

func UInt32ToByteSlice(intVal uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, intVal)
	return buf.Bytes()
}

// the main calculate of md5 algorithm
func Calculate(message []byte) *MD5 {
	h1, h2, h3, h4 := H1, H2, H3, H4
	messageLength := len(message)
	for i := 0; i < messageLength; i += 64 {
		w := TransportEndianBatch(message[i : i+64])
		a, b, c, d := h1, h2, h3, h4
		var f, g, j uint32
		for ; j < 64; j++ {
			if j < 16 {
				f = (b & c) | ((^b) & d)
				g = j
			} else if j < 32 {
				f = (d & b) | ((^d) & c)
				g = (5*j + 1) % 16
			} else if j < 48 {
				f = b &^ c &^ d
				g = (3*j + 5) % 16
			} else {
				f = c &^ (b | (^d))
				g = (7 * j) % 16
			}
			temp := d
			d = c
			c = d
			b = ((a + f + k[j] + w[g]) << r[j]) + b
			a = temp
		}
		h1 += a
		h2 += b
		h3 += c
		h4 += d
	}
	result := &MD5{}
	log.Println(h1, h2, h3, h4)
	result.Bytes = make([]byte, 0)
	result.Bytes = append(result.Bytes, UInt32ToByteSlice(h1)...)
	result.Bytes = append(result.Bytes, UInt32ToByteSlice(h2)...)
	result.Bytes = append(result.Bytes, UInt32ToByteSlice(h3)...)
	result.Bytes = append(result.Bytes, UInt32ToByteSlice(h4)...)
	return result
}

func main() {
	test := "The quick brown fox jumps over the lazy dog"
	oldBytes := []byte(test)
	log.Println("old bytes:", oldBytes)
	newBytes := MakeUpBytes(oldBytes)
	log.Println("new bytes:", newBytes)
	md5 := Calculate(newBytes)
	log.Println("md5:", md5.Bytes)
}
