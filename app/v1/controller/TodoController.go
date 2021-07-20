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

func getQueryString(c *gin.Context, UserID int) string {
	UserID = int(GetUserIdFromSession(c))
	StringUserID := strconv.FormatInt(int64(UserID), 10)
	QueryString := "user_id = " + StringUserID
	Keyword, _ := c.GetQuery("keyword")
	if Keyword != "" {
		QueryString += " AND todo_title like '%" + Keyword + "%'"
	}
	TodoGroupID, _ := c.GetQuery("todo_group_id")
	if TodoGroupID != "" {
		QueryString += " AND todo_group_id = " + TodoGroupID
	}
	IsFinished, _ := c.GetQuery("is_finished")
	if IsFinished != "" {
		if IsFinished == "true" {
			QueryString += " AND is_finished = 1"
		} else {
			QueryString += " AND is_finished = 0"
		}
	}
	return QueryString
}

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
	todoJson.UserID = int(GetUserIdFromSession(c))
	log.Println(GetUserIdFromSession(c))
	QueryString := getQueryString(c, todoJson.UserID)
	log.Println(QueryString)

	var QueryTime = struct {
		CreateAt time.Time `form:"create_at"`
	}{}

	if t, _ := c.GetQuery("create_at"); t != "" {
		if err := c.ShouldBind(&QueryTime); err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(http.StatusBadRequest, err.Error()))
			return
		}
	} else {
		QueryTime.CreateAt = time.Now()
	}

	res := todoModel.GetUserTodoItem(todoJson, QueryString, QueryTime.CreateAt)
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
