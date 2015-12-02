package tcp

import (
	"fmt"
	"net"
	"os"
	"time"
	"etnet/tcp/modbus"
	"github.com/astaxie/beego"
	"wb/cc"
	"wb/st"
	"encoding/binary"
	"wb/ut"
	"etnet/models/statusMgr"
)

const (
	MAX_CONN_NUM = 1000
)

const (
	ctrlAddr		=0
	funcNo			=1
	dataStart 		=3
)


func GetStatus(mConn *modbus.MConn) (string, map[string]interface{}){
	statusMap := make(map[string]interface{})
	mConn.GetRegisters(statusMap, 40001, 14, 2)
	mConn.GetRegister(statusMap, 40023)
	mConn.GetRegisters(statusMap, 40051, 6, 2)
	mConn.GetRegisters(statusMap, 40062, 2, 2)
	mConn.GetRegister(statusMap, 40065)
	mConn.GetRegisters(statusMap, 40067, 2, 1)
	mConn.GetRegisters(statusMap, 40071, 2, 2)
	mConn.GetRegisters(statusMap, 40079, 2, 1)
	mConn.GetRegister(statusMap, 43001)
	mConn.GetRegister(statusMap, 43004)
	mConn.GetRegister(statusMap, 43005)
	mConn.GetRegister(statusMap, 43010)
	mConn.GetRegister(statusMap, 43012)
	if len(statusMap) < len(modbus.MoMap){
		return st.Failed, nil
	}
	return st.Success, statusMap
}
func SendCmd(cmd int) {
	fmt.Println("Cmd: ", cmd)
}
//echo server Goroutine
func EchoFunc(mConn modbus.MConn) {
	conn := mConn.Conn
	cmdChan := mConn.CmdCh
	defer conn.Close()
	for {
		select {
		case cmd:=<-cmdChan:
			SendCmd(cmd)
		case <- time.After(20 * time.Second):
		}
		status, vMap := GetStatus(&mConn)
		if status == st.Success{
			vMap[cc.Sn] = mConn.Id
			vMap[cc.CreateTime] = ut.GetCreateTime()
			fmt.Println(vMap)
			statusMgr.AddStatus(vMap)
		}
		if mConn.IsConnected == false{
			break
		}
	}
	statusMgr.RemoveStatus(mConn.Id)
	fmt.Println("Close connect for id: ", mConn.Id)
}

func ReadId(conn net.Conn) int64{
	buf := make([]byte, 128)
	i, err := conn.Read(buf)
	if err != nil {
		println("Error receive Id:", err.Error())
		return 0
	}
	ids := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if i > 8{
		i = 8
	}
	for idx, v := range buf[0:i]{
		ids[idx + 8 -i ] = v
	}
	id := int64(binary.BigEndian.Uint64(ids))
	fmt.Println("Receive Id:", i, id)
	return id
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
	conn_chan := make(chan modbus.MConn)
	ch_conn_change := make(chan int)
	mapConn := make(map[int64] modbus.MConn)

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
		mConn := modbus.MConn{Id:id, Conn:conn, CmdCh:make(chan int), IsConnected:true}
		mapConn[id]=mConn
		conn_chan <- mConn
	}
}
