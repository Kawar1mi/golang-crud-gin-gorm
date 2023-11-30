package handler

import (
	"net/http"

	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/data/response"
	"github.com/gin-gonic/gin"
)

func WrapError(ctx *gin.Context, err error) {
	WrapErrorWithStatus(ctx, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(ctx *gin.Context, err error, httpStatus int) {

	webResponse := response.Response{
		Status: "error",
		Data:   err.Error(),
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Header("X-Content-Type-Options", "nosnniff")
	ctx.JSON(httpStatus, webResponse)

}

func WrapOk(ctx *gin.Context, data any, httpStatus int) {

	webResponse := response.Response{
		Status: "OK",
		Data:   data,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(httpStatus, webResponse)
}
