package user

import "snow.user/app/models/usermodel"

func Create(mobile string) (id int64, err error) {
	um := usermodel.GetInstance()
	user := &usermodel.User{}
	user.Mobile = mobile
	id, err = um.Insert(user)
	return
}