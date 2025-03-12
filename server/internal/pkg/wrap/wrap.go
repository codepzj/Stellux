package wrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Code int    `json:"code,omitempty"`
	Data T      `json:"data,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func respond[T any](code int, data T, msg string) *Response[T] {
	return &Response[T]{Code: code, Data: data, Msg: msg}
}

func Success[T any](data T, msg string) *Response[T] {
	return respond(http.StatusOK, data, msg)
}

func Fail[T any](code int, data T, msg string) *Response[T] {
	return respond(code, data, msg)
}

func Wrap[T any](fn func(ctx *gin.Context) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := fn(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, data)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}

func WrapWithBody[R any, T any](fn func(ctx *gin.Context, req R) (T, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req R
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, Fail[any](http.StatusBadRequest, nil, err.Error()))
			return
		}
		data, err := fn(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, data)
			return
		}
		ctx.JSON(http.StatusOK, data)
	}
}
