package es

import (
	"fmt"
	"testing"
)

func TestEs(t *testing.T) {
	//sql := "select count(*) from hub_phone where app_id = '600002' and addDate = '2020-08-22' limit 10"
	//sli := strings.Split(sql, " ")
	//for _, v := range sli {
	//	v = strings.ToUpper(v)
	//	if v == `UPDATE` || v == "DELETE" {
	//		return
	//	}
	//}
	returnDatas := Select("select app_id from hub_phone where app_id = '600002' and addDate = '2020-08-22' limit 1")
	fmt.Println(returnDatas)
}
