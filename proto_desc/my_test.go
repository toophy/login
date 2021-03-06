package proto_desc

import (
	"fmt"
	. "github.com/toophy/login/help"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

type Mallocer struct {
}

func (this *Mallocer) Alloc(s int) {
}

func (this *Mallocer) Free(s interface{}) {

}

func (this *Mallocer) relloc(s int, a int) {
	if a < s {
		return
	}

	block := make([]byte, a)

	first := uintptr(unsafe.Pointer(&block[0]))

	count := a/s - 1

	for i := 0; i < count; i++ {
		x := first + uintptr(i*s)
		*(*uintptr)(unsafe.Pointer(x)) = x
	}

	fmt.Print(block)
}

func Tcmalloc(t *testing.T) {

	var mc Mallocer
	mc.relloc(10, 100)

	// block_1 := make([]byte, 4096)
	// block_2 := make([]byte, 4096)
	// block_3 := make([]byte, 4096)

	// *(*uintptr)(unsafe.Pointer(&block_1[0])) = uintptr(unsafe.Pointer(&block_2[0]))
	// *(*uintptr)(unsafe.Pointer(&block_2[0])) = uintptr(unsafe.Pointer(&block_3[0]))
	// *(*uintptr)(unsafe.Pointer(&block_3[0])) = uintptr(0)

	// fmt.Print(block_1[:8])
	// fmt.Print("\n")
	// fmt.Print(block_2[:8])
	// fmt.Print("\n")
	// fmt.Print(block_3[:8])
	// fmt.Print("\n")

	// fmt.Println(&block_1[0])
	// fmt.Println(&block_2[0])
	// fmt.Println(&block_3[0])
}

func TestMain(t *testing.T) {
	Tcmalloc(t)
	//MyByte(t)
}

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

func ParseToNewGolangRow(d string) {

	// mesasge -- 规则解释, } 结束符
	// enum    -- 规则解释, } 结束符

	r1 := strings.Replace(d, "\t", " ", -1)
	r2 := strings.Replace(r1, "\r\n", " ", -1)
	r3 := strings.Replace(r2, "\n", " ", -1)

	m := strings.Fields(r3)
	lens := len(m)

	fmt.Println(lens)
	fmt.Println(m)

}

func MyByte(t *testing.T) {
	// 捕捉异常
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				println("TestByte:" + r.(error).Error())
			case string:
				println("TestByte:" + r.(string))
			}
		}
	}()

	rg, err := strconv.ParseInt("3", 10, 32)
	if err != nil {
		println(rg, err.Error())
	} else {
		println(rg)
	}

	tb_1 := make([]byte, 1024)

	var s Stream
	s.Init(tb_1)

	s.WriteInt24(-16)
	s.WriteUint24(16)
	s.WriteInt24(-8388607)

	s.Seek(0)
	println(s.ReadInt24())
	println(s.ReadInt24())
	println(s.ReadInt24())

	ParseToNewGolangRow("       message              Acter        1      呵呵   {")

}
