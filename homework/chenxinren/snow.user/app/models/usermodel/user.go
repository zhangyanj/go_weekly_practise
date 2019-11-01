package usermodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
)

var (
	once sync.Once
	m    *UserModel
)

//实体
type User struct {
	Id            int64     `xorm:"'id' bigint(20) pk autoincr"`
	Mobile        string    `xorm:"'mobile' varchar(11)"`
}

//表名
func (m *User) TableName() string {
	return "users"
}

//私有化，防止被外部new
type UserModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *UserModel {
	once.Do(func() {
		m = new(UserModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}

func (m *UserModel) Insert(user *User) (id int64, err error) {
	return m.GetDb().Insert(user)
}
