package usermodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
	"time"
)

var (
	once sync.Once
	m    *UserModel
)

//实体
type User struct {
	Id        int       `xorm:"'id' int(10) pk autoincr"`
	Name      string    `xorm:"'name' varchar(255)"`
	Mobile    string    `xorm:"'mobile' varchar(255)"`
	Email     string    `xorm:"'email' varchar(255)"`
	Status    int       `xorm:"'status' TINYINT"`
	CreatedAt time.Time `xorm:"'created_at' timestamp"`
	UpdatedAt time.Time `xorm:"'updated_at' timestamp"`
}

//表名
func (m *User) TableName() string {
	return "user"
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

func (m *UserModel) SaveUser(user User) (affectedRows int64, err error) {
	return m.GetDb().Table(new(User)).InsertOne(user)
}

func (m *UserModel) GetListByName(name string) (*User, error) {
	var model *User
	_, err := m.GetDb().Table(new(User)).
		Where("name = ?", name).
		Get(model)
	return model, err
}

func (m *UserModel) UpListById(id string, user User) (affectedRows int64, err error) {
	return m.GetDb().Table(new(User)).Where("id = ?", id).Update(&User{})
}
