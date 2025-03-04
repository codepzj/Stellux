package wrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

func respond(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(code, Response{Code: code, Data: data, Msg: msg})
}

func Success(c *gin.Context) {
	respond(c, http.StatusOK, nil, "操作成功")
}

func SuccessWithMsg(c *gin.Context, msg string) {
	respond(c, http.StatusOK, nil, msg)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	respond(c, http.StatusOK, data, "操作成功")
}

func SuccessWithDetail(c *gin.Context, data interface{}, msg string) {
	respond(c, http.StatusOK, data, msg)
}

func Fail(c *gin.Context, code int) {
	respond(c, code, nil, "操作失败")
}

func FailWithMsg(c *gin.Context, code int, msg string) {
	respond(c, code, nil, msg)
}

func FailWithData(c *gin.Context, code int, data interface{}) {
	respond(c, code, data, "Failure")
}

func FailWithDetail(c *gin.Context, code int, data interface{}, msg string) {
	respond(c, code, data, msg)
}
