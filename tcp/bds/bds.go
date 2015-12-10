package bds
import (
	"etnet/tcp/tcpserver"
	"github.com/astaxie/beego"
	"net"
	"fmt"
)

const (
	msgFlag = 0x7E

)

const (
	mtRegister = iota
	mtAuth
	mtHeartBeat
	mtPosUpload
)

func CheckRegister(conn net.Conn, buf []byte) bool{
	if buf[1] !=0x01 || buf[2] != 0x00{
		return false
	}
	beego.Info(fmt.Sprintf("TERMINAL[%s] REGISTER :", "test"))
	response := []byte{0x7E, 0x81, 0x00, 0x00, 0x13, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x01, 0x00, 0x54, 0x52, 0x32, 0x30, 0x31, 0x33, 0x30, 0x31, 0x31, 0x37, 0x31, 0x36, 0x34, 0x37, 0x31, 0x34, 0xDF, 0x7E}
	beego.Info("SEND TERMINAL:",response)
	conn.Write(response)
	return true
}

func CheckAuth(conn net.Conn, buf[]byte)bool{
	if buf[1] !=0x01 || buf[2] != 0x02{
		return false
	}
	beego.Info(fmt.Sprintf("TERMINAL[%s] AUTH :", "test"))
	response := []byte{0x7E, 0x80, 0x01, 0x00, 0x05, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x00, 0xC9, 0x7E}
	beego.Info("SEND TERMINAL:", response)
	conn.Write(response)
	return true
}
func CheckHeartBeat(conn net.Conn, buf []byte)bool{
	if buf[1] !=0x00 || buf[2] != 0x02{
		return false
	}
	beego.Info(fmt.Sprintf("TERMINAL[%s] HEARTBEAT :", "test"))
	response := []byte{0x7E, 0x80, 0x01, 0x00, 0x05, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x00, 0xC9, 0x7E}
	beego.Info("SEND TERMINAL:", response)
	conn.Write(response)
	return true
}
func CheckPos(conn net.Conn, buf[]byte)bool{
	if buf[1] !=0x02 || buf[2] != 0x00{
		return false
	}
	beego.Info(fmt.Sprintf("TERMINAL[%s] POSUPLOAD :", "test"))
	response := []byte{0x7E, 0x80, 0x01, 0x00, 0x05, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x00, 0xC9, 0x7E}
	beego.Info("SEND TERMINAL:", response)
	conn.Write(response)
	return true
}

func EchoFunc(conn net.Conn){
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		i, err := conn.Read(buf)
		if err != nil {
			beego.Error("Error receive re:", err.Error())
			//		panic("network broken")
			return
		}
		beego.Info(fmt.Sprintf("RECEIVE TERMINAL[%s]:", "test"), buf[0:i])
		if CheckHeartBeat(conn, buf[0:i]){
			continue
		}
		if CheckRegister(conn, buf[0:i]){
			continue
		}
		if CheckAuth(conn, buf[0:i]){
			continue
		}
		if CheckPos(conn, buf[0:i]){
			continue
		}
		beego.Info(fmt.Sprintf("TERMINAL[%s] ERROR :", "test"))
		response := []byte{0x7E, 0x80, 0x01, 0x00, 0x05, 0x01, 0x39, 0x09, 0x25, 0x50, 0x08, 0x00, 0x00, 0x00, 0x02, 0x01, 0x02, 0x00, 0xC9, 0x7E}
		beego.Info("SEND TERMINAL:", response)
		conn.Write(response)
	}
}

func ServerRun() {
	port := beego.AppConfig.String("bdsport")
	server.Run(port, EchoFunc)
}