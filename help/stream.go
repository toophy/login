package help

import (
	"errors"
	"math"
)

type Stream struct {
	Data   []byte
	MaxLen uint64
	Pos    uint64
}

func (t *Stream) Init(d []byte) {
	t.Data = d
	t.MaxLen = uint64(len(t.Data))
	t.Pos = 0
}

func (t *Stream) Seek(p uint64) {
	if t.Data != nil && p >= 0 && p < t.MaxLen {
		t.Pos = p
	}
}

func (t *Stream) GetPos() uint64 {
	return t.Pos
}

func (t *Stream) GetMaxLen() uint64 {
	return t.MaxLen
}

func (t *Stream) GetData() []byte {
	return t.Data
}

func (t *Stream) ReadBool() bool {
	if t.Pos < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 1
		return t.Data[o] != 0
	}
	panic(errors.New("Stream:ReadBool no long"))
}

func (t *Stream) ReadInt8() int8 {
	if t.Pos < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 1
		return int8(t.Data[o])
	}
	panic(errors.New("Stream:ReadInt8 no long"))
}

func (t *Stream) ReadInt16() int16 {
	if t.Pos+1 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 2

		return (int16(t.Data[o])) |
			(int16(t.Data[o+1]) << 8)
	}
	panic(errors.New("Stream:ReadInt16 no long"))
}

func (t *Stream) ReadInt24() int32 {
	if t.Pos+2 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 3
		n := (uint32(t.Data[o])) |
			(uint32(t.Data[o+1]) << 8) |
			(uint32(t.Data[o+2]) << 16)
		if t.Data[o+2]&0x80 == 0x80 {
			n |= 0xff000000
		}
		return int32(n)
	}
	panic(errors.New("Stream:ReadInt24 no long"))
}

func (t *Stream) ReadInt32() int32 {
	if t.Pos+3 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 4
		return (int32(t.Data[o])) |
			(int32(t.Data[o+1]) << 8) |
			(int32(t.Data[o+2]) << 16) |
			(int32(t.Data[o+3]) << 24)
	}
	panic(errors.New("Stream:ReadInt32 no long"))
}

func (t *Stream) ReadInt40() int64 {
	if t.Pos+4 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 5
		n := (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32)
		if t.Data[o+4]&0x80 == 0x80 {
			n |= 0xffffff0000000000
		}
		return int64(n)
	}
	panic(errors.New("Stream:ReadInt40 no long"))
}

func (t *Stream) ReadInt48() int64 {
	if t.Pos+5 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 6
		n := (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40)
		if t.Data[o+5]&0x80 == 0x80 {
			n |= 0xffff000000000000
		}
	}
	panic(errors.New("Stream:ReadInt48 no long"))
}

func (t *Stream) ReadInt56() int64 {
	if t.Pos+6 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 7
		n := (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40) |
			(uint64(t.Data[o+6]) << 48)
		if t.Data[o+6]&0x80 == 0x80 {
			n |= 0xff00000000000000
		}
	}
	panic(errors.New("Stream:ReadInt56 no long"))
}

func (t *Stream) ReadInt64() int64 {
	if t.Pos+7 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 8
		return (int64(t.Data[o])) |
			(int64(t.Data[o+1]) << 8) |
			(int64(t.Data[o+2]) << 16) |
			(int64(t.Data[o+3]) << 24) |
			(int64(t.Data[o+4]) << 32) |
			(int64(t.Data[o+5]) << 40) |
			(int64(t.Data[o+6]) << 48) |
			(int64(t.Data[o+7]) << 56)
	}
	panic(errors.New("Stream:ReadInt64 no long"))
}

func (t *Stream) ReadUint8() uint8 {
	if t.Pos < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 1
		return uint8(t.Data[o])
	}
	panic(errors.New("Stream:ReadUint8 no long"))
}

func (t *Stream) ReadUint16() uint16 {
	if t.Pos+1 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 2
		return (uint16(t.Data[o])) |
			(uint16(t.Data[o+1]) << 8)
	}
	panic(errors.New("Stream:ReadUint16 no long"))
}

func (t *Stream) ReadUint24() uint32 {
	if t.Pos+2 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 3

		return (uint32(t.Data[o])) |
			(uint32(t.Data[o+1]) << 8) |
			(uint32(t.Data[o+2]&0x7F) << 16) |
			(uint32(t.Data[o+2]&0x80) << 24)
	}
	panic(errors.New("Stream:ReadUint24 no long"))
}

func (t *Stream) ReadUint32() uint32 {
	if t.Pos+3 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 4
		return (uint32(t.Data[o])) |
			(uint32(t.Data[o+1]) << 8) |
			(uint32(t.Data[o+2]) << 16) |
			(uint32(t.Data[o+3]) << 24)
	}
	panic(errors.New("Stream:ReadUint32 no long"))
}

func (t *Stream) ReadUint40() uint64 {
	if t.Pos+4 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 5
		return (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32)
	}
	panic(errors.New("Stream:ReadUint40 no long"))
}

func (t *Stream) ReadUint48() uint64 {
	if t.Pos+5 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 6
		return (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40)
	}
	panic(errors.New("Stream:ReadUint48 no long"))
}

func (t *Stream) ReadUint56() uint64 {
	if t.Pos+6 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 7
		return (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40) |
			(uint64(t.Data[o+6]) << 48)
	}
	panic(errors.New("Stream:ReadUint56 no long"))
}

func (t *Stream) ReadUint64() uint64 {
	if t.Pos+7 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 8
		return (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40) |
			(uint64(t.Data[o+6]) << 48) |
			(uint64(t.Data[o+7]) << 56)
	}
	panic(errors.New("Stream:ReadUint64 no long"))
}

func (t *Stream) ReadFloat32() float32 {
	if t.Pos+3 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 4
		n := (uint32(t.Data[o])) |
			(uint32(t.Data[o+1]) << 8) |
			(uint32(t.Data[o+2]) << 16) |
			(uint32(t.Data[o+3]) << 24)
		return math.Float32frombits(n)
	}
	panic(errors.New("Stream:ReadFloat32 no long"))
}

func (t *Stream) ReadFloat64() float64 {
	if t.Pos+7 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 8
		n := (uint64(t.Data[o])) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 32) |
			(uint64(t.Data[o+5]) << 40) |
			(uint64(t.Data[o+6]) << 48) |
			(uint64(t.Data[o+7]) << 56)
		return math.Float64frombits(n)
	}
	panic(errors.New("Stream:ReadFloat64 no long"))
}

func (t *Stream) ReadBytes() []byte {
	data_len := uint64(t.ReadUint16())
	if data_len > 0 && (t.Pos+data_len) < (t.MaxLen+1) {
		o := t.Pos
		t.Pos = t.Pos + data_len
		return t.Data[o : o+data_len]
	}
	panic(errors.New("Stream:ReadBytes no long"))
}

func (t *Stream) ReadString() string {
	data_len := uint64(t.ReadUint16())
	if data_len > 0 && (t.Pos+data_len) < (t.MaxLen+1) {
		o := t.Pos
		t.Pos = t.Pos + data_len
		return string(t.Data[o : o+data_len])
	}
	panic(errors.New("Stream:ReadString no long"))
}

func (t *Stream) WriteBool(d bool) {
	if t.Pos+1 < t.MaxLen {
		if d {
			t.Data[t.Pos] = 1
		} else {
			t.Data[t.Pos] = 0
		}

		t.Pos = t.Pos + 1
		return
	}

	panic(errors.New("Stream:WriteBool no long"))
}

func (t *Stream) WriteInt8(d int8) {
	if t.Pos+1 < t.MaxLen {
		t.Data[t.Pos] = byte(d)
		t.Pos = t.Pos + 1
		return
	}

	panic(errors.New("Stream:WriteInt8 no long"))
}

func (t *Stream) WriteInt16(d int16) {
	if t.Pos+2 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Pos = t.Pos + 2
		return
	}

	panic(errors.New("Stream:WriteInt16 no long"))
}

// [-8388608,8388607]
func (t *Stream) WriteInt24(d int32) {
	if t.Pos+3 < t.MaxLen {
		// -2^23 ~ 2^23-1
		if d >= -8388608 && d <= 8388607 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Pos = t.Pos + 3
			return
		}
		panic(errors.New("Stream:WriteInt24 out of [-8388608,8388607]"))
	}

	panic(errors.New("Stream:WriteInt24 no long"))
}

func (t *Stream) WriteInt32(d int32) {
	if t.Pos+4 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Pos = t.Pos + 4
		return
	}

	panic(errors.New("Stream:WriteInt32 no long"))
}

// [-549755813888,549755813887]
func (t *Stream) WriteInt40(d int64) {
	if t.Pos+5 < t.MaxLen {
		// -2^39 ~ 2^39-1
		if d >= -549755813888 && d <= 549755813887 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Pos = t.Pos + 5
			return
		}
		panic(errors.New("Stream:WriteInt40 out of [-549755813888,549755813887]"))
	}

	panic(errors.New("Stream:WriteInt40 no long"))
}

// [-140737488355328,140737488355327]
func (t *Stream) WriteInt48(d int64) {
	if t.Pos+6 < t.MaxLen {
		// -2^47 ~ 2^47-1
		if d >= -140737488355328 && d <= 140737488355327 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Data[t.Pos+5] = byte(d >> 40)
			t.Pos = t.Pos + 6
			return
		}
		panic(errors.New("Stream:WriteInt48 out of [-140737488355328,140737488355327]"))
	}

	panic(errors.New("Stream:WriteInt48 no long"))
}

// [-36028797018963968,36028797018963967]
func (t *Stream) WriteInt56(d int64) {
	if t.Pos+7 < t.MaxLen {
		// -2^55 ~ 2^55-1
		if d >= -36028797018963968 && d <= 36028797018963967 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Data[t.Pos+5] = byte(d >> 40)
			t.Data[t.Pos+6] = byte(d >> 48)
			t.Pos = t.Pos + 7
			return
		}
		panic(errors.New("Stream:WriteInt56 out of [-36028797018963968,36028797018963967]"))
	}

	panic(errors.New("Stream:WriteInt56 no long"))
}

func (t *Stream) WriteInt64(d int64) {
	if t.Pos+8 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Data[t.Pos+4] = byte(d >> 32)
		t.Data[t.Pos+5] = byte(d >> 40)
		t.Data[t.Pos+6] = byte(d >> 48)
		t.Data[t.Pos+7] = byte(d >> 56)
		t.Pos = t.Pos + 8
		return
	}

	panic(errors.New("Stream:WriteInt64 no long"))
}

func (t *Stream) WriteUint8(d uint8) {
	if t.Pos < t.MaxLen {
		t.Data[t.Pos] = byte(d)
		t.Pos = t.Pos + 1
		return
	}

	panic(errors.New("Stream:WriteUint8 no long"))
}

func (t *Stream) WriteUint16(d uint16) {
	if t.Pos+1 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Pos = t.Pos + 2
		return
	}

	panic(errors.New("Stream:WriteUint16 no long"))
}

// [0,16777215]
func (t *Stream) WriteUint24(d uint32) {
	if t.Pos+2 < t.MaxLen {
		// 0 ~ 2^24-1
		if d <= 16777215 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Pos = t.Pos + 3
			return
		}
		panic(errors.New("Stream:WriteUint24 out of [0,16777215]"))

	}

	panic(errors.New("Stream:WriteUint24 no long"))
}

func (t *Stream) WriteUint32(d uint32) {
	if t.Pos+3 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Pos = t.Pos + 4
		return
	}

	panic(errors.New("Stream:WriteUint32 no long"))
}

// [0,1099511627775]
func (t *Stream) WriteUint40(d uint64) {
	if t.Pos+4 < t.MaxLen {
		// 0 ~ 2^40-1
		if d <= 1099511627775 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Pos = t.Pos + 5
			return
		}
		panic(errors.New("Stream:WriteUint40 out of [0,1099511627775]"))
	}

	panic(errors.New("Stream:WriteUint40 no long"))
}

// [0,1099511627775]
func (t *Stream) WriteUint48(d uint64) {
	if t.Pos+5 < t.MaxLen {
		// 0 ~ 2^48-1
		if d <= 1099511627775 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Data[t.Pos+5] = byte(d >> 40)
			t.Pos = t.Pos + 6
			return
		}
		panic(errors.New("Stream:WriteUint48 out of [0,1099511627775]"))
	}

	panic(errors.New("Stream:WriteUint48 no long"))
}

// [0,72057594037927935]
func (t *Stream) WriteUint56(d uint64) {
	if t.Pos+6 < t.MaxLen {
		// 0 ~ 2^56-1
		if d <= 72057594037927935 {
			t.Data[t.Pos+0] = byte(d)
			t.Data[t.Pos+1] = byte(d >> 8)
			t.Data[t.Pos+2] = byte(d >> 16)
			t.Data[t.Pos+3] = byte(d >> 24)
			t.Data[t.Pos+4] = byte(d >> 32)
			t.Data[t.Pos+5] = byte(d >> 40)
			t.Data[t.Pos+6] = byte(d >> 48)
			t.Pos = t.Pos + 7
			return
		}
		panic(errors.New("Stream:WriteUint56 out of [0,72057594037927935]"))
	}

	panic(errors.New("Stream:WriteUint56 no long"))
}

func (t *Stream) WriteUint64(d uint64) {
	if t.Pos+7 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Data[t.Pos+4] = byte(d >> 32)
		t.Data[t.Pos+5] = byte(d >> 40)
		t.Data[t.Pos+6] = byte(d >> 48)
		t.Data[t.Pos+7] = byte(d >> 56)
		t.Pos = t.Pos + 8
		return
	}

	panic(errors.New("Stream:WriteUint64 no long"))
}

func (t *Stream) WriteFloat32(d float32) {
	c := math.Float32bits(d)

	if t.Pos+3 < t.MaxLen {
		t.Data[t.Pos+0] = byte(c)
		t.Data[t.Pos+1] = byte(c >> 8)
		t.Data[t.Pos+2] = byte(c >> 16)
		t.Data[t.Pos+3] = byte(c >> 24)
		t.Pos = t.Pos + 4
		return
	}

	panic(errors.New("Stream:WriteFloat32 no long"))
}

func (t *Stream) WriteFloat64(d float64) {
	c := math.Float64bits(d)

	if t.Pos+7 < t.MaxLen {
		t.Data[t.Pos+0] = byte(c)
		t.Data[t.Pos+1] = byte(c >> 8)
		t.Data[t.Pos+2] = byte(c >> 16)
		t.Data[t.Pos+3] = byte(c >> 24)
		t.Data[t.Pos+4] = byte(c >> 32)
		t.Data[t.Pos+5] = byte(c >> 40)
		t.Data[t.Pos+6] = byte(c >> 48)
		t.Data[t.Pos+7] = byte(c >> 56)
		t.Pos = t.Pos + 8
		return
	}

	panic(errors.New("Stream:WriteFloat64 no long"))
}

func (t *Stream) WriteBytes(d []byte) {
	d_len := uint64(len(d))

	if t.Pos+d_len+1 < t.MaxLen {
		t.WriteUint16(uint16(d_len))
		copy(t.Data[t.Pos:], d[:])
		t.Pos = t.Pos + d_len
		return
	}

	panic(errors.New("Stream:WriteBytes no long"))
}

func (t *Stream) WriteString(d *string) {
	d_len := uint64(len(*d))

	if t.Pos+d_len+1 < t.MaxLen {
		t.WriteUint16(uint16(d_len))
		copy(t.Data[t.Pos:t.Pos+d_len], (*d)[:])
		t.Pos = t.Pos + d_len
		return
	}
	panic(errors.New("Stream:WriteString no long"))
}
