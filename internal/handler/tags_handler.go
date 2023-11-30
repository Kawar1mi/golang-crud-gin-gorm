package handler

import (
	"net/http"
	"strconv"

	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/data/request"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/helper"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/processor"
	"github.com/gin-gonic/gin"
)

type TagsHandler struct {
	processor *processor.TagsProcessor
}

func NewTagsHandler(processor *processor.TagsProcessor) *TagsHandler {
	handler := new(TagsHandler)
	handler.processor = processor
	return handler
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			201 {object} response.Response{}
// @Router			/tags [post]
func (h *TagsHandler) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagRequest{}

	err := ctx.ShouldBindJSON(&createTagRequest)
	if err != nil {
		helper.LogError(err)
		WrapError(ctx, err)
		return
	}

	id, err := h.processor.Create(createTagRequest)
	if err != nil {
		WrapError(ctx, err)
		return
	}

	data := map[string]int{"id": id}

	WrapOk(ctx, data, http.StatusCreated)
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (h *TagsHandler) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagRequest{}

	err := ctx.ShouldBindJSON(&updateTagRequest)
	if err != nil {
		helper.LogError(err)
		WrapError(ctx, err)
		return
	}

	id, ok := getTagId(ctx)
	if !ok {
		return
	}

	updateTagRequest.Id = id

	err = h.processor.Update(updateTagRequest)
	if err != nil {
		WrapError(ctx, err)
		return
	}

	WrapOk(ctx, nil, http.StatusOK)
}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Param			tagId path string true "delete tags by id"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [delete]
func (h *TagsHandler) Delete(ctx *gin.Context) {

	id, ok := getTagId(ctx)
	if !ok {
		return
	}

	err := h.processor.Delete(id)
	if err != nil {
		WrapError(ctx, err)
		return
	}

	WrapOk(ctx, nil, http.StatusOK)
}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Description			Return the tahs whoes tagId valu mathes id.
// @Param				tagId path string true "find tags by id"
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (h *TagsHandler) GetById(ctx *gin.Context) {

	id, ok := getTagId(ctx)
	if !ok {
		return
	}

	tagResponse, err := h.processor.GetById(id)
	if err != nil {
		WrapError(ctx, err)
		return
	}

	WrapOk(ctx, tagResponse, http.StatusOK)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (h *TagsHandler) GetAll(ctx *gin.Context) {

	tagsResponse, err := h.processor.GetAll()
	if err != nil {
		WrapError(ctx, err)
		return
	}

	WrapOk(ctx, tagsResponse, http.StatusOK)
}

func getTagId(ctx *gin.Context) (int, bool) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	if err != nil {
		helper.LogError(err)
		WrapError(ctx, err)
		return 0, false
	}

	return id, true
}
