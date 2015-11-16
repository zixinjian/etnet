package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_CONN_NUM = 5
)

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		i, err := conn.Read(buf)
		if err != nil {
			//println("Error reading:", err.Error())
			return
		}
		fmt.Println("客户端发来数据:", i, string(buf[0:i]))
		//send reply
		_, err = conn.Write(buf[0:i])
		fmt.Println("向客户端发送数据:", string(buf[0:i]))
		if err != nil {
			//println("Error send reply:", err.Error())
			return
		}
	}
}

//initial listener and run
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("running ...\n")

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