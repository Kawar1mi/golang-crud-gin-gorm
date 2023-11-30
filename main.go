package main

import (
	"os"

	_ "github.com/Kawar1mi/golang-crud-gin-gorm/docs"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/config"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/handler"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/helper"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/model"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/processor"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/router"
	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/storage"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /v1
func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tag{})

	tagStorage := storage.NewTagsPostgresStorage(db)
	tagProcessor := processor.NewTagsProcessor(tagStorage, validate)
	tagHandler := handler.NewTagsHandler(tagProcessor)

	router := router.NewRouter(tagHandler)

	addr := os.Getenv("ADDR")

	log.Info().Msgf("starting server on: %s", addr)

	err := router.Run(addr)
	helper.FatalIfError(err)
}
