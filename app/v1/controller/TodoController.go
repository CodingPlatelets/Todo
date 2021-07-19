package controller

import (
	"Todo/app/helper"
	"Todo/app/v1/model"
	"Todo/constants"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AddTodo(c *gin.Context) {
	todoModel := model.TodoItem{}
	// todoValidte := validate.TodoValidate
	var todoJson model.TodoItem

	todoJson.UserID = int(GetUserIdFromSession(c))

	if todoJson.UserID == -1 {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "未登录，请先登录"))
		return
	}

	if err := c.ShouldBindJSON(&todoJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}

	//todoMap := helper.Struct2Map(todoJson)
	//if res, err := todoValidte.ValidateMap(todoMap, "add"); !res {
	//	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
	//	return
	//}
	todoJson.CreateAt = time.Now()
	todoJson.IsFinished = false
	res := todoModel.AddTodoItem(todoJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg))
	return
}

// GetUsersTodoList create_at: 'datetime'  # 根据创建时间筛选
// keyword: str  # 根据关键词筛选
// todo_group_id: int  # 根据分组筛选
// is_finished: bool  # 根据是否已完成筛选
func GetUsersTodoList(c *gin.Context) {
	todoModel := model.TodoItem{}
	//todoValidte := validate.TodoValidate
	var todoJson model.TodoItem
	log.Println(GetUserIdFromSession(c))
	todoJson.UserID = int(GetUserIdFromSession(c))

	if todoJson.UserID == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "未登录，请先登录"))
		return
	}

	//if err := c.ShouldBindJSON(&todoJson); err != nil {
	//	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定失败", err.Error()))
	//	return
	//}

	//todoMap := helper.Struct2Map(todoJson)
	//if res, err := todoValidte.ValidateMap(todoMap, "get"); !res {
	//	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "输入信息不完整或有误", err.Error()))
	//	return
	//}

	res := todoModel.GetUserTodoItem(todoJson)
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "OK",
		"todo_list": res.Data,
	})
	return
}

func UpdateTodoItem(c *gin.Context) {
	todoModel := model.TodoItem{}
	var todoJson model.TodoItem
	TodoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)

	updateJson, _ := helper.JsonToMap(string(body))
	log.Println(updateJson)

	todoJson.TodoID = TodoID
	res := todoModel.UpdateTodoItem(todoJson, updateJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg))
	return
}

func DeleteTodoItem(c *gin.Context) {
	todoModel := model.TodoItem{}

	TodoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
		return
	}
	//if err := c.ShouldBindUri(&todoJson); err != nil {
	//	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
	//	return
	//}
	log.Println(TodoID)
	res := todoModel.DeleteTodoItemByID(TodoID)

	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, res.Msg))
	return
}
