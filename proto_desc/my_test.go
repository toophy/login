package proto_desc

import (
	"testing"
	"unsafe"
)

// func GetInt32(buf []byte) (n int32) {
// 	n = int32(buf[0]) | (int32(buf[1]) << 8) | (int32(buf[2]) << 16) | (int32(buf[3]) << 24)
// 	return
// }

// func WriteInt16(n int16, buf []byte) {
// 	buf[0] = byte(n & 0xff)
// 	buf[1] = byte((n & 0x7f00) >> 8)
// }

func WriteInt64(n int64, buf []byte) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			println("\nhere:\n" + r.(error).Error())
		}
	}()
	*(*int64)(unsafe.Pointer(&buf[0])) = n
}

func GetInt64(buf []byte) (n int64) {
	n = *(*int64)(unsafe.Pointer(&buf[0]))
	return
}

func WriteInt32(n int32, buf []byte) {
	*(*int32)(unsafe.Pointer(&buf[0])) = n
}

func GetInt32(buf []byte) (n int32) {
	n = *(*int32)(unsafe.Pointer(&buf[0]))
	return
}

func WriteInt16(n int16, buf []byte) {
	*(*int16)(unsafe.Pointer(&buf[0])) = n
}

func GetInt16(buf []byte) (n int16) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			println("\n2here:\n" + r.(error).Error())
		}
	}()
	n = *(*int16)(unsafe.Pointer(&buf[0]))
	return
}

func TestByte(t *testing.T) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			println("\nhere:\n" + r.(error).Error())
		}
	}()

	tb_1 := []byte{0, 0, 0}
	WriteInt64(0x0102, tb_1)
	println(GetInt16(tb_1))

	var s Stream
	println(unsafe.Pointer(&tb_1[0]))
	s.Init(tb_1)

}
