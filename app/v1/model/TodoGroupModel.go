package model

import (
	"Todo/app/helper"
	"Todo/constants"
	"log"
)

type TodoGroup struct {
	TodoGroupID   int    `gorm:"primaryKey;todo_group_id" json:"todo_group_id"`
	TodoGroupName string `gorm:"todo_group_name" json:"todo_group_name"`
	UserID        int    `json:"user_id" gorm:"user_id"`
	ItemCount     int    `json:"item_count" gorm:"item_count"`
	User User `gorm:"foreignKey:UserID"`
}

func (TodoGroup) TableName() string {
	return "todo_group"
}

func GetReturnData(todoGroups []TodoGroup) []map[string]interface{} {
	var todoGroupsInfo []map[string]interface{}
	for _, todoGroup := range todoGroups {
		todoGroupsInfo = append(todoGroupsInfo, map[string]interface{}{
			"todo_group_id": todoGroup.TodoGroupID,
			"todo_group_name": todoGroup.TodoGroupName,
			"item_count": todoGroup.ItemCount,
		})
	}
	return todoGroupsInfo
}

func (model *TodoGroup) AddTodoGroup(todoGroup TodoGroup) helper.ReturnType {
	err := db.Create(&todoGroup).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: ""}
}

func (model *TodoGroup) GetAllTodoGroups(todoGroup TodoGroup) helper.ReturnType {
	var todoGroups []TodoGroup
	err := db.
		Model(&TodoGroup{}).
		Where("user_id = ?", todoGroup.UserID).
		Select([]string{"todo_group_name", "todo_group_id", "item_count"}).
		Find(&todoGroups).
		Error
	log.Println()
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}
	returnData := GetReturnData(todoGroups)
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: returnData}
}

func (model *TodoGroup) UpdateTodoGroup(todoGroup TodoGroup) helper.ReturnType {
	err := db.
		Model(&TodoGroup{}).
		Select([]string{"todo_group_id", "todo_group_name", "item_count"}).
		Where("todo_group_id", todoGroup.TodoGroupID).
		Updates(todoGroup).
		Error
	if err != nil {
		return helper.ReturnType{
			Status: constants.CodeError,
			Msg: err.Error(),
			Data: -1,
		}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "修改成功", Data: 0}
}

func (model *TodoGroup) DeleteTodoGroup(todoGroup TodoGroup) helper.ReturnType {
	err := db.
		Model(&TodoGroup{}).
		Where("todo_group_id = ?", todoGroup.TodoGroupID).
		Delete(&todoGroup).
		Error
	if err != nil {
		return helper.ReturnType{
			Status: constants.CodeError,
			Msg: err.Error(),
			Data: -1,
		}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "修改成功", Data: 0}

}

func (model TodoGroup) GetTodoGroupByID(todoGroupID int) TodoGroup {
	var todoGroup TodoGroup
	err := db.
		Model(&TodoGroup{}).
		Where("todo_group_id = ?", todoGroupID).
		Find(&todoGroup).
		Error
	if err != nil {
		return TodoGroup{}
	}
	return todoGroup
}