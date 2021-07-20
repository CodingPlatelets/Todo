package controller

import (
	"Todo/app/helper"
	"Todo/app/v1/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AddTodoGroup(c *gin.Context) {
	todoGroupModel := model.TodoGroup{}
	var todoGroupJson model.TodoGroup

	if err := c.ShouldBindJSON(&todoGroupJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, err.Error()))
		return
	}

	todoGroupModel.AddTodoGroup(todoGroupJson)
	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, "OK"))

}

func GetAllTodoGroup(c *gin.Context) {
	todoGroupModel := model.TodoGroup{}
	var todoGroupJson model.TodoGroup

	todoGroupJson.UserID = GetUserIdFromSession(c)

	res := todoGroupModel.GetAllTodoGroups(todoGroupJson)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "OK",
		"todo_group_list": res.Data,
	})
	return
}

func UpdateTodoGroup(c *gin.Context) {
	var todoGroupJson model.TodoGroup
	TodoGroupID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, err.Error()))
		return
	}

	if err = c.ShouldBindJSON(&todoGroupJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, err.Error()))
		return
	}

	todoGroupJson.TodoGroupID = TodoGroupID
	log.Println(todoGroupJson)

	todoGroupModel := model.TodoGroup{}
	res := todoGroupModel.UpdateTodoGroup(todoGroupJson)

	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, res.Msg))
	return
}

func DeleteTodoGroup(c *gin.Context) {
	var todoGroupJson model.TodoGroup
	todoGroupModel := model.TodoGroup{}

	todoGroupId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, err.Error()))
		return
	}

	//if err = c.ShouldBindJSON(&todoGroupJson); err != nil {
	//	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error()))
	//	return
	//}

	todoGroupJson.TodoGroupID = todoGroupId
	log.Println(todoGroupJson)

	res := todoGroupModel.DeleteTodoGroup(todoGroupJson)

	c.JSON(http.StatusOK, helper.ApiReturn(http.StatusOK, res.Msg))
	return
}
