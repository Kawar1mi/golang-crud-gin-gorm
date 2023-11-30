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
