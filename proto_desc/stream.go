package proto_desc

import (
	"unsafe"
)

type Stream struct {
	Data   []byte
	MaxLen uint64
	Pos    uint64
}

func (t *Stream) Init(d []byte) {
	t.Data = d
	t.MaxLen = len(t.Data)
	t.Pos = 0
}

func (t *Stream) Seek(p uint64) {
	if t.Data != nil && p >= 0 && p < t.MaxLen {
		t.Pos = p
	}
}

func (t *Stream) ReadInt8() int64 {
	if t.Pos < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 1
		return int64(t.Data[o])
	}
	return 0
}

func (t *Stream) ReadInt16() int64 {
	if t.Pos+1 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 2
		return (int64(t.Data[o]) << 8) |
			int64(t.Data[o+1])
	}
	return 0
}

func (t *Stream) ReadInt24() int64 {
	if t.Pos+2 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 3
		return (int64(t.Data[o]) << 16) |
			(int64(t.Data[o+1]) << 8) |
			(int64(t.Data[o+2]))
	}
	return 0
}

func (t *Stream) ReadInt32() int64 {
	if t.Pos+3 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 4
		return (int64(t.Data[o]) << 24) |
			(int64(t.Data[o+1]) << 16) |
			(int64(t.Data[o+2]) << 8) |
			(int64(t.Data[o+3]))
	}
	return 0
}

func (t *Stream) ReadInt40() int64 {
	if t.Pos+4 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 5
		return (int64(t.Data[o]) << 32) |
			(int64(t.Data[o+1]) << 24) |
			(int64(t.Data[o+2]) << 16) |
			(int64(t.Data[o+3]) << 8) |
			(int64(t.Data[o+4]))
	}
	return 0
}

func (t *Stream) ReadInt48() int64 {
	if t.Pos+5 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 6
		return (int64(t.Data[o]) << 40) |
			(int64(t.Data[o+1]) << 32) |
			(int64(t.Data[o+2]) << 24) |
			(int64(t.Data[o+3]) << 16) |
			(int64(t.Data[o+4]) << 8) |
			(int64(t.Data[o+5]))
	}
	return 0
}

func (t *Stream) ReadInt56() int64 {
	if t.Pos+6 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 7
		return (int64(t.Data[o]) << 48) |
			(int64(t.Data[o+1]) << 40) |
			(int64(t.Data[o+2]) << 32) |
			(int64(t.Data[o+3]) << 24) |
			(int64(t.Data[o+4]) << 16) |
			(int64(t.Data[o+5]) << 8) |
			(int64(t.Data[o+6]))
	}
	return 0
}

func (t *Stream) ReadInt64() int64 {
	if t.Pos+7 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 8
		return (int64(t.Data[o+0]) << 56) |
			(int64(t.Data[o+1]) << 48) |
			(int64(t.Data[o+2]) << 40) |
			(int64(t.Data[o+3]) << 32) |
			(int64(t.Data[o+4]) << 24) |
			(int64(t.Data[o+5]) << 16) |
			(int64(t.Data[o+6]) << 8) |
			(int64(t.Data[o+7]))
	}
	return 0
}

func (t *Stream) ReadUint8() uint64 {
	if t.Pos < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 1
		return uint64(t.Data[o])
	}
	return 0
}

func (t *Stream) ReadUint16() uint64 {
	if t.Pos+1 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 2
		return (uint64(t.Data[o]) << 8) |
			(uint64(t.Data[o+1]))
	}
	return 0
}

func (t *Stream) ReadUint24() uint64 {
	if t.Pos+2 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 3
		return (uint64(t.Data[o]) << 16) |
			(uint64(t.Data[o+1]) << 8) |
			(uint64(t.Data[o+2]))
	}
	return 0
}

func (t *Stream) ReadUint32() uint64 {
	if t.Pos+3 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 4
		return (uint64(t.Data[o]) << 24) |
			(uint64(t.Data[o+1]) << 16) |
			(uint64(t.Data[o+2]) << 8) |
			(uint64(t.Data[o+3]))
	}
	return 0
}

func (t *Stream) ReadUint40() uint64 {
	if t.Pos+4 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 5
		return (uint64(t.Data[o]) << 32) |
			(uint64(t.Data[o+1]) << 24) |
			(uint64(t.Data[o+2]) << 16) |
			(uint64(t.Data[o+3]) << 8) |
			(uint64(t.Data[o+4]))
	}
	return 0
}

func (t *Stream) ReadUint48() uint64 {
	if t.Pos+5 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 6
		return (uint64(t.Data[o]) << 40) |
			(uint64(t.Data[o+1]) << 32) |
			(uint64(t.Data[o+2]) << 24) |
			(uint64(t.Data[o+3]) << 16) |
			(uint64(t.Data[o+4]) << 8) |
			(uint64(t.Data[o+5]))
	}
	return 0
}

func (t *Stream) ReadUint56() uint64 {
	if t.Pos+6 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 7
		return (uint64(t.Data[o]) << 48) |
			(uint64(t.Data[o+1]) << 40) |
			(uint64(t.Data[o+2]) << 32) |
			(uint64(t.Data[o+3]) << 24) |
			(uint64(t.Data[o+4]) << 16) |
			(uint64(t.Data[o+5]) << 8) |
			(uint64(t.Data[o+6]))
	}
	return 0
}

func (t *Stream) ReadUint64() uint64 {
	if t.Pos+7 < t.MaxLen {
		o := t.Pos
		t.Pos = t.Pos + 8
		return (uint64(t.Data[o]) << 56) |
			(uint64(t.Data[o+1]) << 48) |
			(uint64(t.Data[o+2]) << 40) |
			(uint64(t.Data[o+3]) << 32) |
			(uint64(t.Data[o+4]) << 24) |
			(uint64(t.Data[o+5]) << 16) |
			(uint64(t.Data[o+6]) << 8) |
			(uint64(t.Data[o+7]))
	}
	return 0
}

func (t *Stream) ReadBytes(d int) []byte {
	data_len := t.ReadUint16()
	if data_len > 0 && (t.Pos+data_len) < (t.MaxLen+1) {
		o := t.Pos
		t.Pos = t.Pos + data_len
		return t.Data[o : o+data_len]
	}
	return nil
}

func (t *Stream) ReadString() string {
	data_len := t.ReadUint16()
	if data_len > 0 && (t.Pos+data_len) < (t.MaxLen+1) {
		o := t.Pos
		t.Pos = t.Pos + data_len
		return string(t.Data[o : o+data_len])
	}
	return ""
}

func (t *Stream) WriteInt8(d int64) bool {
	if t.Pos+1 < t.MaxLen {
		t.Data[t.Pos] = byte(d)
		t.Pos = t.Pos + 1
		return true
	}

	return false
}

func (t *Stream) WriteInt16(d int64) bool {
	if t.Pos+2 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 8)
		t.Data[t.Pos+1] = byte(d)
		t.Pos = t.Pos + 2
		return true
	}

	return false
}

func (t *Stream) WriteInt24(d int64) bool {
	if t.Pos+3 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 16)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d)
		t.Pos = t.Pos + 3
		return true
	}

	return false
}

func (t *Stream) WriteInt32(d int64) bool {
	if t.Pos+4 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 24)
		t.Data[t.Pos+1] = byte(d >> 16)
		t.Data[t.Pos+2] = byte(d >> 8)
		t.Data[t.Pos+3] = byte(d)
		t.Pos = t.Pos + 4
		return true
	}

	return false
}

func (t *Stream) WriteInt40(d int64) bool {
	if t.Pos+5 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 32)
		t.Data[t.Pos+1] = byte(d >> 24)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 8)
		t.Data[t.Pos+4] = byte(d)
		t.Pos = t.Pos + 5
		return true
	}

	return false
}

func (t *Stream) WriteInt48(d int64) bool {
	if t.Pos+6 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 40)
		t.Data[t.Pos+1] = byte(d >> 32)
		t.Data[t.Pos+2] = byte(d >> 24)
		t.Data[t.Pos+3] = byte(d >> 16)
		t.Data[t.Pos+4] = byte(d >> 8)
		t.Data[t.Pos+5] = byte(d)
		t.Pos = t.Pos + 6
		return true
	}

	return false
}

func (t *Stream) WriteInt56(d int64) bool {
	if t.Pos+7 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 48)
		t.Data[t.Pos+1] = byte(d >> 40)
		t.Data[t.Pos+2] = byte(d >> 32)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Data[t.Pos+4] = byte(d >> 16)
		t.Data[t.Pos+5] = byte(d >> 8)
		t.Data[t.Pos+6] = byte(d)
		t.Pos = t.Pos + 7
		return true
	}

	return false
}

func (t *Stream) WriteInt64(d int64) bool {
	if t.Pos+8 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 56)
		t.Data[t.Pos+1] = byte(d >> 48)
		t.Data[t.Pos+2] = byte(d >> 40)
		t.Data[t.Pos+3] = byte(d >> 32)
		t.Data[t.Pos+4] = byte(d >> 24)
		t.Data[t.Pos+5] = byte(d >> 16)
		t.Data[t.Pos+6] = byte(d >> 8)
		t.Data[t.Pos+7] = byte(d)
		t.Pos = t.Pos + 8
		return true
	}

	return false
}

func (t *Stream) WriteUint8(d uint64) bool {
	if t.Pos < t.MaxLen {
		t.Data[t.Pos] = byte(d)
		t.Pos = t.Pos + 1
		return true
	}

	return false
}

func (t *Stream) WriteUint16(d uint64) bool {
	if t.Pos+1 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 8)
		t.Data[t.Pos+1] = byte(d)
		t.Pos = t.Pos + 2
		return true
	}

	return false
}

func (t *Stream) WriteUint24(d uint64) bool {
	if t.Pos+2 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 16)
		t.Data[t.Pos+1] = byte(d >> 8)
		t.Data[t.Pos+2] = byte(d)
		t.Pos = t.Pos + 3
		return true
	}

	return false
}

func (t *Stream) WriteUint32(d uint64) bool {
	if t.Pos+3 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 24)
		t.Data[t.Pos+1] = byte(d >> 16)
		t.Data[t.Pos+2] = byte(d >> 8)
		t.Data[t.Pos+3] = byte(d)
		t.Pos = t.Pos + 4
		return true
	}

	return false
}

func (t *Stream) WriteUint40(d uint64) bool {
	if t.Pos+4 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 32)
		t.Data[t.Pos+1] = byte(d >> 24)
		t.Data[t.Pos+2] = byte(d >> 16)
		t.Data[t.Pos+3] = byte(d >> 8)
		t.Data[t.Pos+4] = byte(d)
		t.Pos = t.Pos + 5
		return true
	}

	return false
}

func (t *Stream) WriteUint48(d uint64) bool {
	if t.Pos+5 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 40)
		t.Data[t.Pos+1] = byte(d >> 32)
		t.Data[t.Pos+2] = byte(d >> 24)
		t.Data[t.Pos+3] = byte(d >> 16)
		t.Data[t.Pos+4] = byte(d >> 8)
		t.Data[t.Pos+5] = byte(d)
		t.Pos = t.Pos + 6
		return true
	}

	return false
}

func (t *Stream) WriteUint56(d uint64) bool {
	if t.Pos+6 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 48)
		t.Data[t.Pos+1] = byte(d >> 40)
		t.Data[t.Pos+2] = byte(d >> 32)
		t.Data[t.Pos+3] = byte(d >> 24)
		t.Data[t.Pos+4] = byte(d >> 16)
		t.Data[t.Pos+5] = byte(d >> 8)
		t.Data[t.Pos+6] = byte(d)
		t.Pos = t.Pos + 7
		return true
	}

	return false
}

func (t *Stream) WriteUint64(d uint64) bool {
	if t.Pos+7 < t.MaxLen {
		t.Data[t.Pos+0] = byte(d >> 56)
		t.Data[t.Pos+1] = byte(d >> 48)
		t.Data[t.Pos+2] = byte(d >> 40)
		t.Data[t.Pos+3] = byte(d >> 32)
		t.Data[t.Pos+4] = byte(d >> 24)
		t.Data[t.Pos+5] = byte(d >> 16)
		t.Data[t.Pos+6] = byte(d >> 8)
		t.Data[t.Pos+7] = byte(d)
		t.Pos = t.Pos + 8
		return true
	}

	return false
}

func (t *Stream) WriteString(d *string) bool {
	d_len := uint64(len(*d))

	if t.Pos+d_len+1 < t.MaxLen {
		if t.WriteUint16(d_len) {
			ds := (*d)[:]
			dx := t.Data[t.Pos : t.Pos+d_len]
			copy(dx, ds)
			t.Pos = t.Pos + d_len
			return true
		}
	}
	return false
}
