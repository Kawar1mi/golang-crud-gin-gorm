package storage

import (
	"fmt"

	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/helper"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/model"
	"gorm.io/gorm"
)

type TagsPostgresStorage struct {
	db *gorm.DB
}

func NewTagsPostgresStorage(db *gorm.DB) *TagsPostgresStorage {
	storage := new(TagsPostgresStorage)
	storage.db = db
	return storage
}

func (t *TagsPostgresStorage) Create(tag model.Tag) (int, error) {
	result := t.db.Create(&tag)
	if result.Error != nil {
		helper.LogError(result.Error)
		return 0, result.Error
	}

	return tag.Id, nil
}

func (t *TagsPostgresStorage) Update(tag model.Tag) error {
	result := t.db.Model(&tag).Where("id = ?", tag.Id).Update("name", tag.Name)
	if result.Error != nil {
		helper.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (t *TagsPostgresStorage) Delete(tagId int) error {
	var tag model.Tag
	result := t.db.Where("id = ?", tagId).Delete(&tag)
	if result.Error != nil {
		helper.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (t *TagsPostgresStorage) GetAll() ([]model.Tag, error) {
	var tags []model.Tag

	result := t.db.Find(&tags)
	if result.Error != nil {
		helper.LogError(result.Error)
		return tags, result.Error
	}

	return tags, nil
}

func (t *TagsPostgresStorage) GetById(tagId int) (model.Tag, error) {
	var tag model.Tag

	result := t.db.Find(&tag, tagId)
	if result.Error != nil {
		helper.LogError(result.Error)
		return tag, result.Error
	}

	if result.RowsAffected == 0 {
		err := fmt.Errorf("tag not found by id %v", tagId)
		helper.LogError(err)
		return tag, err
	}

	return tag, nil
}
