package model

import (
	"Todo/app/helper"
	"Todo/constants"
)

type User struct {
	UserID   int    `gorm:"primaryKey;int;user_id" json:"user_id"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func (User) TableName() string {
	return "users"
}

// AddUser 添加用户
func (model *User) AddUser(data User) helper.ReturnType {
	user := User{}
	// 判断昵称是否已存在
	err := db.
		Select("username").
		Where("username = ?", data.Username).
		First(&user).
		Error
	if err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: user.Username}
	}
	// 创建记录
	err = db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: 1}
	}
}

func (model *User) LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.Where("username = ? AND password = ?", data.Username, data.Password).First(&user).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		resp := make(map[string]interface{})
		resp["userInfo"] = user
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "登录验证成功", Data: resp}
	}
}
