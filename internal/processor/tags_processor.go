package processor

import (
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/data/request"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/data/response"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/helper"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/model"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/storage"
	"github.com/go-playground/validator/v10"
)

type TagsProcessor struct {
	storage  *storage.TagsPostgresStorage
	validate *validator.Validate
}

func NewTagsProcessor(storage *storage.TagsPostgresStorage, validate *validator.Validate) *TagsProcessor {
	processor := new(TagsProcessor)
	processor.storage = storage
	processor.validate = validate
	return processor
}

func (t *TagsProcessor) Create(tag request.CreateTagRequest) (int, error) {
	err := t.validate.Struct(tag)
	if err != nil {
		helper.LogError(err)
		return 0, err
	}

	tagModel := model.Tag{
		Name: tag.Name,
	}

	return t.storage.Create(tagModel)
}

func (t *TagsProcessor) Update(tag request.UpdateTagRequest) error {

	err := t.validate.Struct(tag)
	if err != nil {
		helper.LogError(err)
		return err
	}

	tagModel := model.Tag{
		Name: tag.Name,
		Id:   tag.Id,
	}

	return t.storage.Update(tagModel)
}

func (t *TagsProcessor) Delete(tagId int) error {
	return t.storage.Delete(tagId)
}

func (t *TagsProcessor) GetAll() ([]response.TagResponse, error) {
	var tags []response.TagResponse

	result, err := t.storage.GetAll()
	if err != nil {
		return tags, err
	}

	for _, v := range result {
		tag := response.TagResponse{
			Id:   v.Id,
			Name: v.Name,
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (t *TagsProcessor) GetById(tagId int) (response.TagResponse, error) {
	var tag response.TagResponse

	result, err := t.storage.GetById(tagId)
	if err != nil {
		return tag, err
	}

	tag.Id, tag.Name = result.Id, result.Name

	return tag, nil
}
