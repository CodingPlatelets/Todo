package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type ReturnType struct {
	Status int
	Msg    string
	Data   interface{}
}

// 结构体转换为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		mapValue := v.Field(i).Interface()
		// 递归获取数据
		if reflect.TypeOf(mapValue).Kind() == reflect.Struct {
			if v.Field(i).Type().String() != "time.Time" {
				innerMap := Struct2Map(mapValue)
				for key, value := range innerMap {
					data[key] = value
				}
			}
		}
		// 转换驼峰为下划线
		upperField := t.Field(i).Name
		field := ""
		index := 0
		for j := 0; j < len(upperField)-1; j++ {
			if (upperField[j] >= 'a' && upperField[j] <= 'z') &&
				(upperField[j+1] >= 'A' && upperField[j+1] <= 'Z') {
				field += upperField[index:j+1] + "_"
				index = j + 1
			}
		}
		field += upperField[index:]
		data[strings.ToLower(field)] = v.Field(i).Interface()
	}
	return data
}

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		//fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	//
	//for k, v := range m {
	//	fmt.Printf("%v: %v\n", k, v)
	//}

	return m, nil
}

// 模块内统一返回格式
func ReturnRes(status int, msg string, data interface{}) ReturnType {
	returnType := ReturnType{status, msg, data}
	return returnType
}

func ApiReturn(status int, msg string) gin.H {
	return gin.H{
		"code": status,
		"msg":  msg,
	}
}
