package proto_desc

import (
	"fmt"
	"github.com/toophy/login/help"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func GetReadFunc(s string) string {
	func_name := ""
	switch s {
	case "int8":
		func_name = "ReadInt8"
	case "int16":
		func_name = "ReadInt16"
	case "int32":
		func_name = "ReadInt32"
	case "int64":
		func_name = "ReadInt64"
	case "uint8":
		func_name = "ReadUint8"
	case "uint16":
		func_name = "ReadUint16"
	case "uint32":
		func_name = "ReadUint32"
	case "uint64":
		func_name = "ReadUint64"
	case "string":
		func_name = "ReadString"
	}
	return func_name
}

func GetWriteFunc(s string) string {
	func_name := ""
	switch s {
	case "int8":
		func_name = "WriteInt8"
	case "int16":
		func_name = "WriteInt16"
	case "int32":
		func_name = "WriteInt32"
	case "int64":
		func_name = "WriteInt64"
	case "uint8":
		func_name = "WriteUint8"
	case "uint16":
		func_name = "WriteUint16"
	case "uint32":
		func_name = "WriteUint32"
	case "uint64":
		func_name = "WriteUint64"
	case "string":
		func_name = "WriteString"
	}
	return func_name
}

func TypeToClang(s string) string {
	type_name := ""
	switch s {
	case "int8":
		type_name = "char"
	case "int16":
		type_name = "short"
	case "int32":
		type_name = "int"
	case "int64":
		type_name = "long long"
	case "uint8":
		type_name = "unsigned char"
	case "uint16":
		type_name = "unsigned short"
	case "uint32":
		type_name = "unsigned int"
	case "uint64":
		type_name = "unsigned long long"
	case "string":
		type_name = "std::string"
	}
	return type_name
}

func ParseToGolang(d string, fd string, f string) {

	r1 := strings.Replace(d, "\t", " ", -1)
	r2 := strings.Replace(r1, "\r\n", " ", -1)
	r3 := strings.Replace(r2, "\n", " ", -1)
	r4 := strings.Replace(r3, "{", " ", -1)
	r5 := strings.Replace(r4, "}", " ", -1)

	m := strings.Fields(r5)

	lens := len(m)

	data_struct := make([]string, 10)
	data_read := make([]string, 10)
	data_write := make([]string, 10)

	data_count := 0

	flag := "key"
	member := ""

	for i := 0; i < lens; i++ {
		switch m[i] {
		case "message":
			switch flag {
			case "key":
			default:
				if flag == "member" {
					data_count++
				} else {
					fmt.Printf("flag=%s,解析失败[%s]\n", flag, data_struct[data_count])
					data_struct[data_count] = ""
					data_read[data_count] = ""
					data_write[data_count] = ""
				}
			}
			flag = "name"
		default:
			switch flag {
			case "name":
				data_struct[data_count] += "type " + m[i] + " struct {\n"
				data_read[data_count] += "func (this *" + m[i] + ") Read(s *Stream) {\n"
				data_write[data_count] += "func (this *" + m[i] + ") Write(s *Stream) {\n"
				flag = "member"
			case "member":
				member = m[i]
				flag = "type"
			case "type":
				if len(member) > 0 {
					data_struct[data_count] += "\t" + member + "\t" + m[i] + "\n"
					func_read := GetReadFunc(m[i])
					if func_read == "" {
						data_read[data_count] += "\t" + "this." + member + ".Read(s)\n"
					} else {
						data_read[data_count] += "\t" + "s." + GetReadFunc(m[i]) + "(this." + member + ")\n"
					}

					func_write := GetWriteFunc(m[i])
					if func_write == "" {
						data_write[data_count] += "\t" + "this." + member + ".Write(s)\n"
					} else {
						data_write[data_count] += "\t" + "s." + GetWriteFunc(m[i]) + "(this." + member + ")\n"
					}
				} else {
					fmt.Printf("flag=%s,解析失败[%s]\n", flag, data_struct[data_count])
				}

				flag = "member"
			}
		}
	}

	for i := 0; i <= data_count; i++ {
		data_struct[i] += "}\n\n"
		data_read[i] += "}\n\n"
		data_write[i] += "}\n\n"
	}

	// 检查log目录
	if !help.IsExist(fd) {
		os.MkdirAll(fd, os.ModeDir)
	}

	if !help.IsExist(fd + f) {
		os.Create(fd + f)
	}
	file, err := os.OpenFile(fd+f, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i <= data_count; i++ {
		file.WriteString(data_struct[i])
		file.WriteString(data_read[i])
		file.WriteString(data_write[i])
	}

	file.Close()

}

func ParseToCpplang(d string, fd string, f string) {

	r1 := strings.Replace(d, "\t", " ", -1)
	r2 := strings.Replace(r1, "\r\n", " ", -1)
	r3 := strings.Replace(r2, "\n", " ", -1)
	r4 := strings.Replace(r3, "{", " ", -1)
	r5 := strings.Replace(r4, "}", " ", -1)

	m := strings.Fields(r5)

	lens := len(m)

	data_struct := make([]string, 10)
	data_read := make([]string, 10)
	data_write := make([]string, 10)

	data_count := 0

	flag := "key"
	member := ""

	for i := 0; i < lens; i++ {
		switch m[i] {
		case "message":
			switch flag {
			case "key":
			default:
				if flag == "member" {
					data_count++
				} else {
					fmt.Printf("flag=%s,解析失败[%s]\n", flag, data_struct[data_count])
					data_struct[data_count] = ""
					data_read[data_count] = ""
					data_write[data_count] = ""
				}
			}
			flag = "name"
		default:
			switch flag {
			case "name":
				data_struct[data_count] += "class " + m[i] + " {\n"
				data_read[data_count] += "void " + m[i] + "::Read(s *Stream) {\n"
				data_write[data_count] += "void " + m[i] + "::Write(s *Stream) {\n"
				flag = "member"
			case "member":
				member = m[i]
				flag = "type"
			case "type":
				if len(member) > 0 {
					type_name := TypeToClang(m[i])
					if type_name == "" {
						data_struct[data_count] += "\t" + m[i] + "\t" + member + ";\n"
					} else {
						data_struct[data_count] += "\t" + TypeToClang(m[i]) + "\t" + member + ";\n"
					}

					func_read := GetReadFunc(m[i])
					if func_read == "" {
						data_read[data_count] += "\t" + "this." + member + ".Read(s);\n"
					} else {
						data_read[data_count] += "\t" + "s." + GetReadFunc(m[i]) + "(this." + member + ");\n"
					}

					func_write := GetWriteFunc(m[i])
					if func_write == "" {
						data_write[data_count] += "\t" + "this." + member + ".Write(s);\n"
					} else {
						data_write[data_count] += "\t" + "s." + GetWriteFunc(m[i]) + "(this." + member + ");\n"
					}
				} else {
					fmt.Printf("flag=%s,解析失败[%s]\n", flag, data_struct[data_count])
				}

				flag = "member"
			}
		}
	}

	for i := 0; i <= data_count; i++ {
		data_struct[i] += "};\n\n"
		data_read[i] += "}\n\n"
		data_write[i] += "}\n\n"
	}
	// 检查log目录
	if !help.IsExist(fd) {
		os.MkdirAll(fd, os.ModeDir)
	}

	if !help.IsExist(fd + f) {
		os.Create(fd + f)
	}
	file, err := os.OpenFile(fd+f, os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i <= data_count; i++ {
		file.WriteString(data_struct[i])
		file.WriteString(data_read[i])
		file.WriteString(data_write[i])
	}

	file.Close()
}

func TestToophyBuffers(t *testing.T) {

	desc := "actor.txt"
	if help.IsExist(desc) {

		fi, err := os.Open(desc)
		if err != nil {
			fmt.Println("读文件失败: %s", err.Error())
			return
		}
		fd, err := ioutil.ReadAll(fi)
		fi.Close()

		ParseToGolang(string(fd), "../proto/", "actor.go")
		ParseToCpplang(string(fd), "../proto/", "actor.cpp")
	} else {
		fmt.Println("找不到描述文件 : " + desc)
	}

}
