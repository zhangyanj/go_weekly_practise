package userservice

import (
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/models/userloginsmodel"
	"time"
)

func InsertLoginInfo(userId int, ip string, loginTime time.Time) (id int64, err error) {
	var instanceEntity userloginsmodel.UserLogins
	instanceEntity.UserId = userId
	instanceEntity.Ip = ip
	instanceEntity.LoginTime = loginTime
	id, err = userloginsmodel.GetInstance().InsertUserLogin(instanceEntity)
	return
}
