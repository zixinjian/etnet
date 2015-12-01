package statusMgr
import (
	"wb/cc"
	"github.com/astaxie/beego"
	"wb/om"
	"etnet/models/s"
)


var statusCache = make(map[int64]map[string]interface{})

func AddStatus(statusMap map[string]interface{}){

	if sn, ok:= statusMap[cc.Sn];ok{
		statusCache[sn.(int64)] = statusMap
		om.Table(s.Status).AddMap(statusMap)
	}else{
		beego.Error("Status sn not found!")
	}
}

func GetStatus(sn int64)map[string]interface{}{
	if statusMap, ok := statusCache[sn];ok{
		return statusMap
	}
	return make(map[string]interface{})
}