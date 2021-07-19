package validate

import (
	"Todo/app/helper"
)

var TodoValidate helper.Validator
//
//TodoID   int       `gorm:"todo_id"  json:"todo_id"`
//Content  string    `gorm:"content"  json:"content"`
//Deadline time.Time `gorm:"deadline" json:"deadline"`
//Tag      string    `gorm:"tag" json:"tag"`
//Location string    `gorm:"location" json:"location"`
//Pin2Top  bool      `gorm:"pin2top"  json:"pin2top"`
//CreateAt time.Time `gorm:"create_at" json:"create_at"`
//UserID 	 int 	   `gorm:"user_id" json:"user_id"`

func init() {
	rules := map[string]string{
		"todo_id": "required",
		"content": "required",
		"deadline": "required",
		"tag": "required",
		"location": "required",
		"pin2top": "required",
		"user_id": "required",
		"create_at": "required",
	}

	scenes := map[string][]string{
		"add": {"content", "deadline"},
		"get": {"user_id"},
	}
	TodoValidate.Rules = rules
	TodoValidate.Scenes = scenes
}
