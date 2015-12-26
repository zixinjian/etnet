package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_CONN_NUM = 1000
)

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	i, err := conn.Read(buf)
	if err != nil {
		println("Error reading:", err.Error())
		//		return
	}
	fmt.Println("客户端发来数据:", i, buf[0:i])
	fmt.Println("客户端字符串:", i, string(buf[0:i]))

//	//send reply
//	writeMsg:=[]byte{0x01, 0x03, 0x00, 0x32, 0x00, 0x01, 0x25, 0x5C}
//	_, err = conn.Write(writeMsg)
//	fmt.Println("向客户端发送数据:", string(buf[0:i]))
//	if err != nil {
//		//println("Error send reply:", err.Error())
//		return
//	}
//	_, err = conn.Read(buf)
//	if err != nil {
//		//println("Error reading:", err.Error())
//		return
//	}
//	fmt.Println("客户端发来数据:", i, string(buf[0:i]))
}

//initial listener and run
func main() {
	port := "9172"
	listener, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("TcpServer Running on :", port)

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		cur_conn_numNow := cur_conn_num
		for _ = range time.Tick(1e8) {
			if cur_conn_num != cur_conn_numNow{
				fmt.Printf("cur conn num: %f\n", cur_conn_num)
				cur_conn_numNow = cur_conn_num
			}
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				ch_conn_change <- 1
				EchoFunc(conn)
				ch_conn_change <- -1
			}
		}()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}
