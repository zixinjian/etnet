package userMgr
import (
	"wb/om"
	"wb/cc"
	"wb/st"
	"etnet/models/s"
)

//type User struct {
//	Sn          string
//	UserName	string
//	Name 		string
//	Role		string
//	Company		string
//	Department	string
//	Flag        string
//}

func Query(queryParams om.Params) (string, om.ValueMap) {
	if s, v := om.Table(s.User).Query(queryParams);s == st.Success{
		delete(v, cc.Password)
		return s, v
	}
	return om.Table(s.User).Query(queryParams)
}
func GetValidUser(queryParams om.Params)(string, om.ValueMap){
	if stat, valueMap := Query(queryParams); stat == st.Success{
		if valueMap.GetString(cc.Flag) == cc.Flag_Avaliable{
			return st.Success, valueMap
		}else{
			return st.ItemNotFound, nil
		}
	}else{
		return stat, nil
	}
}



//func GetUserEnum(queryParams t.Params, orderParams t.Params, limitParams t.LimitParams) []itemDef.EnumValue {
//	if len(orderParams) <= 0{
//		orderParams[s.Name]= s.Asc
//	}
//	if code, userMaps := svc.GetItems(s.User, queryParams, orderParams, limitParams); strings.EqualFold(code, stat.Success) {
//		EnumList := make([]itemDef.EnumValue, len(userMaps))
//		for idx, user := range userMaps {
//			v, _ := user[s.Sn]
//			u, _ := user[s.UserName]
//			l, _ := user[s.Name]
//			EnumList[idx] = itemDef.EnumValue{v.(string), u.(string), l.(string)}
//		}
//		return EnumList
//	} else {
//		return make([]itemDef.EnumValue, 0)
//	}
//}