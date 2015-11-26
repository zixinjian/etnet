package tcp

import (
	"fmt"
	"net"
	"os"
	"time"
	"etnet/tcp/modbus"
	"wb/om"
	"etnet/models/s"
	"github.com/astaxie/beego"
	"wb/cc"
	"wb/st"
)

const (
	MAX_CONN_NUM = 1000
)

const (
	ctrlAddr		=0
	funcNo			=1
	dataStart 		=3
)


func GetStatus(conn net.Conn) (string, map[string]interface{}){
	vMap := make(map[string]interface{})
	modbus.GetRegisters(vMap, conn, 40001, 14, 2)
	modbus.GetRegister(vMap, conn, 40023)
	modbus.GetRegisters(vMap, conn, 40051, 6, 2)
	modbus.GetRegisters(vMap, conn, 40062, 2, 2)
	modbus.GetRegister(vMap, conn, 40065)
	modbus.GetRegisters(vMap, conn, 40067, 2, 1)
	modbus.GetRegisters(vMap, conn, 40071, 2, 2)
	modbus.GetRegisters(vMap, conn, 40079, 2, 1)
	modbus.GetRegister(vMap, conn, 43001)
	modbus.GetRegister(vMap, conn, 43004)
	modbus.GetRegister(vMap, conn, 43005)
	modbus.GetRegister(vMap, conn, 43010)
	modbus.GetRegister(vMap, conn, 43012)
//	fmt.Println(len(mapValue), mapValue)
	return st.Success, vMap
}
func SendCmd(cmd int) {
	fmt.Println("Cmd: ", cmd)
}
//echo server Goroutine
func EchoFunc(mConn MCon) {
	conn := mConn.Conn
	cmdChan := mConn.CmdCh
	defer conn.Close()
	for {
		select {
		case cmd:=<-cmdChan:
			SendCmd(cmd)
		case <- time.After(20 * time.Second):
		}
		status, vMap := GetStatus(conn)
		if status == st.Success{
			vMap[cc.Sn] = mConn.Id
			om.AddValueMap(s.Status, vMap)
		}
	}
}

func ReadId(conn net.Conn) string{
	buf := make([]byte, 1024)
	i, err := conn.Read(buf)
	if err != nil {
		println("Error receive Id:", err.Error())
		return ""
	}
	fmt.Println("Receive Id:", i, string(buf[0:i]))
	return string(buf[0:i])
}

type MCon struct {
	Id   string
	Conn net.Conn
	CmdCh chan int
}

func (this *MCon) SendCmd(cmd int){
	this.CmdCh <- cmd
}

func ServerRun() {
	port := beego.AppConfig.String("tcpport")
	listener, err := net.Listen("tcp", "0.0.0.0:" + port)

	if err != nil {
		beego.Error("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	beego.Info("TcpServer Running on :", port)

	var cur_conn_num int = 0
	conn_chan := make(chan MCon)
	ch_conn_change := make(chan int)
	mapConn := make(map[string] MCon)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		cur_conn_numNow := cur_conn_num
		for _ = range time.Tick(1e8) {
			if cur_conn_num != cur_conn_numNow{
				fmt.Printf("cur conn num: %d\n", cur_conn_num)
				cur_conn_numNow = cur_conn_num
			}
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for mConn := range conn_chan {
				ch_conn_change <- 1
				EchoFunc(mConn)
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
		id := ReadId(conn)
		mConn := MCon{Id:id, Conn:conn, CmdCh:make(chan int)}
		mapConn[id]=mConn
		conn_chan <- mConn
	}
}
