package validate

import (
	"Todo/app/helper"
)

var UserValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id":        "required",
		"username":       "required|maxLen:25",
		"password":       "required|minLen:6|maxLen:16",
		"old_password":   "required|minLen:6|maxLen:16",
		"password_check": "required|minLen:6|maxLen:16",
		"check":          "required",
		"status":         "required|int:-1,0",
		"verify_code":    "required",
	}

	scenes := map[string][]string{
		"register":         {"username", "password", "password_check"},
		"login": 			{"username", "password"},
	}
	UserValidate.Rules = rules
	UserValidate.Scenes = scenes
}
