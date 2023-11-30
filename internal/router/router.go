package router

import (
	"net/http"

	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *handler.TagsHandler) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/v1")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.GetAll)
	tagsRouter.GET("/:tagId", tagsController.GetById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
