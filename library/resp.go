package library

import (
	"github.com/gin-gonic/gin"
)

type Redirect struct {
	Location string `json:"location""`
	Token string `json:"token"'`
}

// 返回信息
func RespMessage(code int, message string, data interface{}) interface{} {
	return gin.H{
		"code" : code,
		"message" : message,
		"data" : data,
	}
}
