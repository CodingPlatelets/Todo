package controller

import (
	"Todo/app/helper"
	"Todo/app/v1/model"
	"Todo/app/validate"
	"Todo/constants"
	"Todo/db_server"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RequestEmailValidation(c *gin.Context) {
	var userValidate = validate.UserValidate
	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "request"); err != nil && !res {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, err.Error()))
		return
	}

	res, err := helper.SendMail(userJson.Email)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, res.Msg))
	}
	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, res.Msg))
}

func Register(c *gin.Context) {
	var userModel = model.User{}
	var userValidate = validate.UserValidate

	var userJson struct {
		model.User
		VerifyCode    string `json:"verify_code" form:"verify_code"`
		PasswordCheck string `json:"password_check" form:"password_check"`
	}
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "register"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, err.Error()))
		return
	}

	KeyValue := "VerifyCode" + userJson.User.Email

	VerifyCodeCorrect, err := redis.String(db_server.GetFromRedis(KeyValue))
	db_server.DeleteFromRedis(KeyValue)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, "验证码已过期，请重新发送验证码"))
		return
	}
	log.Print(VerifyCodeCorrect)
	if userJson.VerifyCode != VerifyCodeCorrect {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, "验证码输入错误"))
		return
	}

	if userJson.Password != userJson.PasswordCheck {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "两次密码输入不一致"))
	}

	userJson.Password = helper.GetMd5(userJson.Password)

	res := userModel.AddUser(userJson.User)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg))
	return
}

func DoLogin(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("user_id") != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "已登陆"))
		return
	}

	var userModel = model.User{}
	var userValidate = validate.UserValidate

	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "login"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}

	userJson.Password = helper.GetMd5(userJson.Password)

	res := userModel.LoginCheck(userJson)

	if res.Status == constants.CodeSuccess {
		userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
		returnData := map[string]interface{}{
			"userId":   userInfo.UserID,
			"username": userInfo.Username,
		}
		jsonData, _ := json.Marshal(returnData)
		session.Set("user_id", userInfo.UserID)
		session.Set("username", userInfo.Username)
		session.Set("data", string(jsonData))
		session.Save()
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg))
	return
}

func DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "注销成功"))
}
