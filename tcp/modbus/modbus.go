package modbus
import (
	"fmt"
	"net"
	"wb/ut"
	"encoding/binary"
	"bytes"
	"errors"
)

const (
	Unsigned 	= iota
	Integer
	String0
	Table1
	Table2
	Binary1
	Binary2
)



type MType int32

type Mo struct {
	Register    int32
	Key 		string
	Name        string
	Dim         string
	Type		MType
	Dec         int16
	Len         int16
	RetLen	    int16
}


const (
	ctrlAddr		=0x1
	funcRead		=0x3
	ctrlAddrNo		=0
	funcNo			=1
	dataLenNo       =2
	dataStart 		=3
)

var MoMap =make(map[int32]Mo)
func (this *Mo)GetDecValue(v float64)float64{
	if this.Dec > 0{
		return float64(v)/float64(this.Dec * 10)
	}else{
		return float64(v)
	}
}
func (this *Mo)GetValue(bs []byte)interface{}{
	order := binary.BigEndian
	switch this.Type {
	case Unsigned:
		switch this.RetLen{
		case 1:
			v := bs[0]
			return this.GetDecValue(float64(v))
		case 2:
			v :=order.Uint16(bs)
			return this.GetDecValue(float64(v))
		case 4:
			v :=order.Uint32(bs)
			return this.GetDecValue(float64(v))
		default:
			fmt.Println("Error Mo.Len")
			return 0
		}
	case Integer:
		switch this.RetLen{
		case 1:
			v := int8(bs[0])
			return this.GetDecValue(float64(v))
		case 2:
			v :=int16(order.Uint16(bs))
			return this.GetDecValue(float64(v))
		case 4:
			v :=int32(order.Uint32(bs))
			return this.GetDecValue(float64(v))
		default:
			fmt.Println("Error Mo.Len")
			return 0
		}
	case Table1:
		v :=int16(order.Uint16(bs))
		return this.GetDecValue(float64(v))
	case Binary1, Binary2:
		v :=order.Uint16(bs)
		return fmt.Sprintf("%b", v)
	}
	return 0
}
//func (this *Mo)GetMoValue(mConn MConn)(interface{}, error){
//	bs, err := mConn.GetRegisterBytes(this.Register, this.Len)
//	if err != nil{
//		return "", err
//	}
//	return this.GetValue(bs[dataStart:dataStart+this.RetLen]), nil
//}

type MConn struct {
	Conn net.Conn
	Id   int64
	CmdCh chan string
	IsConnected bool
}

const (
	StartCmd = "start"
 	StopCmd = "stop"
)
func (this *MConn) SendCmd(cmd string)([]byte, error){
	switch cmd {
	case StartCmd:
		req := []byte{0x01, 0x10, 0x18, 0xD6, 0x00, 0x03, 0x06, 0x01, 0xFE, 0x00, 0x00, 0x00, 0x01}
		return this.WriteReq(req)
	case StopCmd:
		req := []byte{0x01, 0x10, 0x18, 0xD6, 0x00, 0x03, 0x06, 0x02, 0xFD, 0x00, 0x00, 0x00, 0x01}
		return this.WriteReq(req)
	}
	return make([]byte, 0), errors.New("No cmd name : " + cmd)
}

func (this *MConn) WriteReq(req []byte)([]byte, error){
	req = AppendCrc(req)
	_, err := this.Conn.Write(req)
	if err != nil {
		fmt.Println("Error send cmd:", err.Error())
		this.IsConnected = false
		//		panic("network broken")
		return []byte{}, err
	}
	buf := make([]byte, 1024)
	i, err := this.Conn.Read(buf)
	if err != nil {
		fmt.Println("Error receive re:", err.Error())
		this.IsConnected = false
		//		panic("network broken")
		return []byte{}, err
	}
//	fmt.Println("WriteCmd Receive : ", ut.BytesToHex(buf[0:i]))
	return buf[0:i], nil
}

const register0 = 40001
func BuildReadRequest(start int32, len int16)([]byte){
	startR := start -  register0
	bsStart := ut.Int16ToBytes(int16(startR))
	bsLen := ut.Int16ToBytes(len)
	req := []byte{0x01, 0x03, bsStart[0], bsStart[1], bsLen[0], bsLen[1]}
	req = AppendCrc(req)
//	fmt.Println("Request is :", ut.BytesToHex(req))
	return req
}

//func BuildWriteRequest(start int32, len int16)

func (mConn *MConn)GetRegisterBytes(start int32, len int16)([]byte, error){
	req := BuildReadRequest(start, len)
	_, err := mConn.Conn.Write(req)
	if err != nil {
		fmt.Println("Error send cmd:", err.Error())
		mConn.IsConnected = false
//		panic("network broken")
		return []byte{}, err
	}

	buf := make([]byte, 1024)
	i, err := mConn.Conn.Read(buf)
	if err != nil {
		fmt.Println("Error receive re:", err.Error())
		mConn.IsConnected = false
//		panic("network broken")
		return []byte{}, err
	}
//	fmt.Println("Receive", ut.BytesToHex(buf[0:i]))
	return buf[0:i], nil
}
func (mConn *MConn)GetRegister(mapValue map[string]interface{}, start int32){
	if mo, ok := MoMap[start];ok{
		bs, err:= mConn.GetRegisterBytes(start, mo.Len)
		if err!= nil{
			return
		}
		vBytes := bs[3:3+mo.RetLen]
//		fmt.Println("Read reg:", start, " ", ut.BytesToHex(vBytes))
		mapValue[mo.Key] = mo.GetValue(vBytes)
	}
}

// 只适合连续的寄存器，并且结果位数相同
func (mConn *MConn)GetRegisters(mapValue map[string]interface{}, start int32, len int16, retLen int16){
	bs , err:= mConn.GetRegisterBytes(start, len)
	if err!= nil{
		return
	}
	for i:=int16(0);i<len;i++{
		r := start + int32(i)
		if mo, ok := MoMap[int32(r)];ok{
			vBytes := bs[3+i*retLen:3+i*retLen+retLen]
//			fmt.Println("Read reg:", r, " ", ut.BytesToHex(vBytes))
			mapValue[mo.Key] = mo.GetValue(vBytes)
		}else {
			continue
		}
	}
}

func Crc16(bs []byte) []byte{
	var crc uint16= 0xFFFF
	for _, b := range bs{
		crc ^= uint16(b)
		for i:= 0; i< 8;i++{
			if(crc & 1 != 0){
				crc >>=1
				crc ^= 0xA001;
			}else{
				crc >>=1
			}
		}
	}
	buff :=  bytes .NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, crc)
	//	fmt.Println(fmt.Sprintf("%0x, %0x", buff.Bytes()[0], buff.Bytes()[1]))
	return buff.Bytes()
}
func AppendCrc(req []byte)[]byte{
	csc := Crc16(req)
	req = append(req, csc[1])
	req = append(req, csc[0])
	return req
}