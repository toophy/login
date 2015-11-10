// easybuff
// 不要修改本文件, 每次消息有变动, 请手动生成本文件
// easybuff -s 描述文件目录 -o 目标文件目录 -l 语言(go,cpp)

package proto

import (
	. "github.com/toophy/login/help"
)

// 背包格上限
const BagGridMax = 128

// 创建角色
const C2m_create_role_Id = 102

type C2m_create_role struct {
	Account string // 帐号
	Name    string // 角色名
	Job     string // 职业
	Age     int64  // 年龄
}

func (t *C2m_create_role) Read(s *Stream) {
	t.Account = s.ReadString()
	t.Name = s.ReadString()
	t.Job = s.ReadString()
	t.Age = s.ReadInt40()
}

func (t *C2m_create_role) Write(s *Stream) {
	s.WriteString(&t.Account)
	s.WriteString(&t.Name)
	s.WriteString(&t.Job)
	s.WriteInt40(t.Age)
}

// 帐号登录
const C2m_login_Id = 100

type C2m_login struct {
	Account string // 帐号
	Token   string // 令牌
	Time    string // 时间戳
	Skey    string // 小区ID
}

func (t *C2m_login) Read(s *Stream) {
	t.Account = s.ReadString()
	t.Token = s.ReadString()
	t.Time = s.ReadString()
	t.Skey = s.ReadString()
}

func (t *C2m_login) Write(s *Stream) {
	s.WriteString(&t.Account)
	s.WriteString(&t.Token)
	s.WriteString(&t.Time)
	s.WriteString(&t.Skey)
}

// 返回创建角色
const M2c_create_role_ret_Id = 103

type M2c_create_role_ret struct {
	Ret  int8   // 返回值
	Msg  string // 返回消息
	Name string // 角色名
	Job  string // 职业
	Age  int64  // 年龄
}

func (t *M2c_create_role_ret) Read(s *Stream) {
	t.Ret = s.ReadInt8()
	t.Msg = s.ReadString()
	t.Name = s.ReadString()
	t.Job = s.ReadString()
	t.Age = s.ReadInt40()
}

func (t *M2c_create_role_ret) Write(s *Stream) {
	s.WriteInt8(t.Ret)
	s.WriteString(&t.Msg)
	s.WriteString(&t.Name)
	s.WriteString(&t.Job)
	s.WriteInt40(t.Age)
}

// 返回登录结果
const M2c_login_ret_Id = 101

type M2c_login_ret struct {
	Ret   int8   // 返回结果
	Msg   string // 返回消息
	Skey  string // 小区ID
	Token string // 令牌
}

func (t *M2c_login_ret) Read(s *Stream) {
	t.Ret = s.ReadInt8()
	t.Msg = s.ReadString()
	t.Skey = s.ReadString()
	t.Token = s.ReadString()
}

func (t *M2c_login_ret) Write(s *Stream) {
	s.WriteInt8(t.Ret)
	s.WriteString(&t.Msg)
	s.WriteString(&t.Skey)
	s.WriteString(&t.Token)
}
