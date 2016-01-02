package main

import (
	"fmt"
	"net"
)

const (
	addr = "p2x.pub:9171"
//	addr = "localhost:9171"
//	addr = "localhost:9172"
//	addr = "p2x.pub:9172"
)

func main() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
//	bdsClient(conn)
	modbusClient(conn)
}
//var ms = map[string][]byte{
//	"01030000000ec40e":[]byte{01,03,1c,00,00,00,00,00,00,00,00,00000000000000000000000000000000000000005ca5",
//}
func modbusClient(conn net.Conn){
	conn.Write([]byte{0x7E, 0x01, 0x00, 0x7E})
	for {

		buf := make([]byte, 128)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
			return
		}
		fmt.Println("rec: ", buf[0:c], c)
		if err != nil {
			fmt.Println("数据输入异常:", err.Error())
			return
		}
		s:= []byte{0x01,0x03,0x1c,1,2,3,4,5,6,7,1,2,10,1,2, 3, 4, 55, 3, 7, 8, 7, 4, 45, 23}
		fmt.Println("send: ", s)
		conn.Write(s)
	}
}

func bdsClient(conn net.Conn){
	s:= []byte{0x7E, 0x01, 0x00, 0x7E}
	conn.Write(s)
	fmt.Println("Send:", s)
	buf := make([]byte, 128)
	c, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器数据异常:", err.Error())
		return
	}
	fmt.Println("rec: ", buf[0:c], c)


	s = []byte{0x7E, 0x01, 0x02, 0x7E}
	conn.Write(s)
	fmt.Println("Send:", s)
	c, err = conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器数据异常:", err.Error())
		return
	}
	fmt.Println("rec: ", buf[0:c], c)


	s = []byte{0x7E, 0x00, 0x02, 0x7E}
	conn.Write(s)
	fmt.Println("Send:", s)
	c, err = conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器数据异常:", err.Error())
		return
	}
	fmt.Println("rec: ", buf[0:c], c)


	conn.Write([]byte{0x7E, 0x01, 0x02, 0x00, 0x05, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x00, 0xC9, 0x7E})
	c, err = conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器数据异常:", err.Error())
		return
	}
	fmt.Println("rec: ", buf[0:c], c)
}
func Client(conn net.Conn) {
	sms := make([]byte, 128)
	for {
		fmt.Print("请输入要发送的消息:")
		_, err := fmt.Scan(&sms)
		if err != nil {
			fmt.Println("数据输入异常:", err.Error())
		}
		conn.Write(sms)
		buf := make([]byte, 128)
		c, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println("xxxx:", string(buf[0:c]), c)
	}

}