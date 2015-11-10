package app

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	MaxDataLen     = 5080
	MaxSendDataLen = 4000
	MaxHeader      = 4
)

// 消息节点(list节点)
type Msg_node struct {
	Len   int32 // 包长度
	Token int32 // 包令牌
	Count int32 // 包内消息数
	Data  *byte // 数据
}

// 网络包
type Net_packet struct {
	Header [MaxHeader]byte // 包头
	Len    uint32          // 包长度
	Token  uint32          // 包令牌
	Count  uint32          // 包内消息数
	Data   *byte           // 数据
}

func (t *Net_packet) Clear() {
	t.Len = 0
	t.Token = 0
	t.Count = 0
}

// 读取网络消息
func (t *Net_packet) ReadHeader(conn *net.TCPConn) error {
	t.Clear()
	length, err := io.ReadFull(conn, t.Header[:])
	if length != MaxHeader {
		GetApp().LogWarn("Net packet header : %d != %d", length, MaxHeader)
		return err
	}
	if err != nil {
		return err
	}

	t.Len = (uint32(t.Header[0])) | (uint32(t.Header[1]) << 8)
	t.Token = uint32(t.Header[2])
	t.Count = uint32(t.Header[3])

	// 根据 t.Len 分配一个 缓冲, 并读取 body
}

func (this *Net_packet) ReadBody(conn *net.TCPConn) error {

}
