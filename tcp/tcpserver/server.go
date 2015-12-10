package server
import (
	"net"
	"github.com/astaxie/beego"
	"time"
	"os"
)


const (
	MAX_CONN_NUM = 1000
)
func Run(port string, echoFunc func(conn net.Conn)) {
	listener, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		beego.Error("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	beego.Info("TcpServer Running on :", port)

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
				beego.Info("cur conn num: %d\n", cur_conn_num)
				cur_conn_numNow = cur_conn_num
			}
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				ch_conn_change <- 1
				echoFunc(conn)
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