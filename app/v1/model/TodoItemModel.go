package model

import (
	"Todo/app/helper"
	"Todo/constants"
	"log"
	"time"
)

type TodoItem struct {
	TodoID      int       `gorm:"primaryKey;todo_id" json:"todo_id" uri:"todo_id"`
	UserID      int       `gorm:"user_id" json:"user_id"`
	TodoGroupID int       `json:"todo_group_id" gorm:"todo_group_id"`
	TodoTitle   string    `json:"todo_title" gorm:"todo_title;omitempty"`
	TodoContent string    `json:"todo_content" gorm:"todo_content;omitempty"`
	CreateAt    time.Time `json:"create_at" gorm:"create_at;autoCreateTime"`
	IsFinished  bool      `json:"is_finished" gorm:"is_finished;omitempty"`
	User        User      `gorm:"foreignKey:UserID"`
	TodoGroup   TodoGroup `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TodoGroupID"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func (model *TodoItem) AddTodoItem(item TodoItem) helper.ReturnType {
	err := db.
		Create(&item).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	todoGroupModel := TodoGroup{}
	todoGroup := todoGroupModel.GetTodoGroupByID(item.TodoGroupID)
	todoGroup.ItemCount += 1
	res := todoGroup.UpdateTodoGroup(todoGroup)
	if res.Status != constants.CodeSuccess {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: res.Data}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: ""}
}

func (model *TodoItem) GetTodoItemByID(TodoID int) helper.ReturnType {
	// err := db.Select([]string{"content", ""})
	todo := TodoItem{}
	err := db.
		Where("todo_id = ?", TodoID).
		First(&todo).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: ""}
}

func (model *TodoItem) GetUserTodoItem(item TodoItem) helper.ReturnType {
	// err := db.Select([]string{"content", ""})
	type result struct {
		TodoID      int       `gorm:"primaryKey;todo_id" json:"todo_id" uri:"todo_id"`
		UserID      int       `gorm:"user_id" json:"user_id"`
		TodoGroupID int       `json:"todo_group_id" gorm:"todo_group_id"`
		TodoTitle   string    `json:"todo_title" gorm:"todo_title"`
		TodoContent string    `json:"todo_content" gorm:"todo_content"`
		CreateAt    time.Time `json:"create_at" gorm:"create_at"`
		IsFinished  bool      `json:"is_finished" gorm:"is_finished"`
	}
	var todoList []result
	err := db.
		Model(&TodoItem{}).
		Where("user_id = ?", item.UserID).
		Scan(&todoList).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: todoList}
}

func (model *TodoItem) UpdateTodoItem(todoItem TodoItem) helper.ReturnType {
	err := db.
		Model(&TodoItem{}).
		Select([]string{"todo_id", "todo_content", "todo_group_id", "todo_title", "is_finished"}).
		Where("todo_id = ?", todoItem.TodoID).
		Updates(todoItem).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: ""}

}

func (model *TodoItem) DeleteTodoItemByID(TodoID int) helper.ReturnType {
	var todoItem TodoItem
	err := db.Where("todo_id = ?", TodoID).Find(&todoItem).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}
	err = db.Delete(&todoItem).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: err.Error()}
	}
	todoGroupModel := TodoGroup{}
	todoGroup := todoGroupModel.GetTodoGroupByID(todoItem.TodoGroupID)
	todoGroup.ItemCount -= 1
	log.Println(todoGroup)
	res := todoGroup.UpdateTodoGroup(todoGroup)
	if res.Status != constants.CodeSuccess {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: res.Data}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: ""}
}
