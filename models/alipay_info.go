package models

import (
	"fmt"
	"time"
	"xorm.io/xorm"
)

type Users struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	Imei         string    `xorm:"not null default '' comment('设备id') VARCHAR(100)"`
	Oaid         string    `xorm:"not null default '' comment('设备id2') VARCHAR(100)"`
	Idfa         string    `xorm:"not null default '' comment('设备id2') VARCHAR(100)"`
	Mobile       string    `xorm:"not null default '' comment('手机号') VARCHAR(64)"`
	LastLogin    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后登陆时间') TIMESTAMP"`
	CreatedAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	Balance      string    `xorm:"not null default 0.00 comment('余额') DECIMAL(10,2)"`
	TotalBalance string    `xorm:"not null default 0.00 comment('总共获取余额') DECIMAL(10,2)"`
	Gold         int       `xorm:"not null comment('金币') INT(11)"`
	RegIp        string    `xorm:"not null comment('注册IP') VARCHAR(64)"`
	TotalGold    int       `xorm:"not null comment('总金币') INT(11)"`
	LastType     string    `xorm:"not null comment('最后的设备类型') VARCHAR(64)"`
	DeviceId     string    `xorm:"not null comment('设备ID') VARCHAR(64)"`
	ParentId     int       `xorm:"not null comment('推荐人Id') INT(11)"`
}


func (U *Users) Insert(dbSession *xorm.Session)(int64,error){
	if dbSession  != nil{
		return dbSession.Insert(U)
	}
	return getDb().Insert(U)
}
func (U *Users) GetByMobile(mobile string)(bool,error){
	return getDb().Where(`mobile=?`,mobile).Get(U)
}

func(U *Users) GetByDeviceId(deviceId string)(bool,error){
	return getDb().Where(`device_id=?`,deviceId).Get(U)
}

func (U *Users) UpdateById(dbSession *xorm.Session,id int,updateMap map[string]interface{}){
	if dbSession != nil {
		_,err := dbSession.ID(id).Update(updateMap)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	_,err := getDb().ID(id).Update(updateMap)
	if err != nil {
		fmt.Println(err)
	}
	return

}