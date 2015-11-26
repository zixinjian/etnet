package modbus
import (
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/csv"
	"wb/u"
)

const(
	keyCol = iota
	registerCol
	comObjCol
	nameCol
	dimCol
	typeCol
	retLenCol
	decCol
	minCol
	maxCol
	groupCol
	lenCol

)

func init(){
	fileName := "conf/modbus.csv"
	cntb, err := ioutil.ReadFile(fileName)
//	fmt.Println(fmt.Sprintf("start parse uiList File:%s", fileName))
	if err != nil {
		panic(err)
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	rows, _ := r2.ReadAll()
//	fmt.Println(rows)
	for i, row := range rows{
		fmt.Println(i, row)
		register, err:=u.StrTo(strings.TrimSpace(row[registerCol])).Int32()
		if err != nil{
//			fmt.Println("row:", i, " to int32 error")
			continue
		}
		retLen, err := u.StrTo(strings.TrimSpace(row[retLenCol])).Int16()
		if err != nil{
			continue
		}
		len, err := u.StrTo(strings.TrimSpace(row[lenCol])).Int16()
		if err != nil{
			continue
		}
		dec, err := u.StrTo(strings.TrimSpace(row[decCol])).Int16()
		if err != nil{
			continue
		}
		name := strings.TrimSpace(row[nameCol])
		dim := strings.TrimSpace(row[dimCol])
		key := strings.TrimSpace(row[keyCol])

		var mType MType
		switch u.StrTo(strings.TrimSpace(row[typeCol])) {
		case "Integer":
			mType = Integer
		case "Unsigned":
			mType = Unsigned
		case "String0":
			mType = String0
		case "Table#1":
			mType = Table1
		case "Table#2":
			mType = Table2
		case "Binary#1":
			mType = Binary1
		case "Binary#2":
			mType = Binary2
		default:
			continue
		}
		MoMap[register]=Mo{Register:register, Name:name, Dim:dim, Type:mType, Len:len, RetLen:retLen, Dec:dec, Key:key}
	}
	fmt.Println("Read ", len(MoMap), "Mos")
}