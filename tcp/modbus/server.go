package modbus

import (
	"fmt"
	"net"
	"time"
	"github.com/astaxie/beego"
	"wb/cc"
	"wb/st"
	"encoding/binary"
	"wb/ut"
	"etnet/models/statusMgr"
	"sync"
	"etnet/tcp/tcpserver"
)




//func DoSendCmd(mConn *modbus.MConn){
//	beego.Debug("WriteStartCmd ")
//	mConn.WriteStartCmd()
//}

func GetStatus(mConn *MConn) (string, map[string]interface{}){
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
	if len(statusMap) < len(MoMap){
		return st.Failed, nil
	}
	return st.Success, statusMap
}
func SendCmd(sn int64, cmd string) {
	if mConn, ok := mapConn[sn];ok{
		fmt.Println("sendCmd [", cmd, "] to sn: ", sn)
		mConn.CmdCh <- cmd
	}
}

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	id := ReadId(conn)
	mConn := MConn{Id:id, Conn:conn, CmdCh:make(chan string), IsConnected:true}
	AddMConn(mConn)
	cmdChan := mConn.CmdCh
	for {
		select {
		case cmd:=<-cmdChan:
			fmt.Println("Send cmd: ", cmd)
			mConn.SendCmd(cmd)
			time.Sleep(10 * time.Second)
		case <- time.After(20 * time.Second):
		}
		status, vMap := GetStatus(&mConn)
		if status == st.Success{
			vMap[cc.Sn] = mConn.Id
			vMap[cc.CreateTime] = ut.GetCreateTime()
			//			fmt.Println(vMap)
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
	buf := make([]byte, 1024)
	i, err := conn.Read(buf)
	if err != nil {
		println("Error register:", err.Error())
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


var mapConn = make(map[int64] MConn)

var lMapConn = sync.Mutex{}
func AddMConn(mConn MConn){
	lMapConn.Lock()
	defer lMapConn.Unlock()
	mapConn[mConn.Id] = mConn
}

func ServerRun() {
	port := beego.AppConfig.String("modbusport")
	server.Run(port, EchoFunc)
}
