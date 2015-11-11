package app

import (
	"io"
	"net"
)

// 消息节点(list节点)
type Msg_node struct {
	Len   uint32 // 包长度
	Token uint32 // 包令牌
	Count uint32 // 包内消息数
	Data  []byte // 数据
}

// 读取网络消息
func ReadConnData(conn *net.TCPConn) (msg Msg_node, ret error) {
	const ConnHeaderSize = 4
	var header [ConnHeaderSize]byte
	var length int
	length, ret = io.ReadFull(conn, header[:])

	if length != ConnHeaderSize {
		GetApp().LogWarn("Net packet header : %d != %d", length, ConnHeaderSize)
		return
	}
	if ret != nil {
		return
	}

	msg.Len = (uint32(header[0])) | (uint32(header[1]) << 8)
	msg.Token = uint32(header[2])
	msg.Count = uint32(header[3])

	GetApp().LogInfo("ReadConnData : len =%d, token=%d, count=%d", msg.Len, msg.Token, msg.Count)

	// 根据 msg.Len 分配一个 缓冲, 并读取 body
	buf := make([]byte, msg.Len)
	length, ret = io.ReadFull(conn, buf[:])
	if length != ConnHeaderSize {
		GetApp().LogWarn("Net packet body : %d != %d", length, msg.Len)
		return
	}
	if ret != nil {
		return
	}

	msg.Data = buf

	return
}
