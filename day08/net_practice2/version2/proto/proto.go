package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。

// 封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)
// 包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

// 我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度
// 在服务端和客户端分别使用上面定义的proto包的Decode和Encode函数处理数据

// 大端BigEndian 数据的高字节，保存在内存的低地址中，而数据的低字节，保存在内存的高地址中
// 这样的存储模式有点儿类似于把数据当作字符串顺序处理：地址由小向大增加，而数据从高位往低位放；
// 小端LittleEndian 数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中
/// 这种存储模式将地址的高低和数据位权有效地结合起来，高地址部分权值高，低地址部分权值低，和我们的逻辑方法一致。
// 每个地址单元对应一个字节，有一个变量的值要存到几个字节中，必然存在着一个如何将多个字节安排的问题，存储和读取的顺序、方向
// 只要选择一种模式，编码和解码中保持一致即可

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer) // 创建字节类型的缓冲区
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length) 
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
