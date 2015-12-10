package device
import (
	"wb/st"
	"etnet/tcp/modbus"
)


func GetDevices()[]Device{
	retList := make([]Device, 0)
	retList = append(retList, Device{Sn:"123456"})
	return retList
}


//func GetParams(sn string)map[string]interface{}{
//	v := ut.Round(440, 2)
//	return map[string]interface{}{s.V:v, s.A:55, s.P:550, "param1":248, "param2":888, "param3":2012, "param4":2344, s.LocX:118.234, s.LocY:44.555}
//}

func SendCmd(sn int64, cmd string)string{
	modbus.SendCmd(sn, cmd)
	return st.Success
}