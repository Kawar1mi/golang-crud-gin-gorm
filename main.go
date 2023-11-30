package main

import (
	"net/http"
	"os"

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

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tag{})

	tagStorage := storage.NewTagsPostgresStorage(db)
	tagProcessor := processor.NewTagsProcessor(tagStorage, validate)
	tagHandler := handler.NewTagsHandler(tagProcessor)

	router := router.NewRouter(tagHandler)

	addr := os.Getenv("ADDR")
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Info().Msgf("starting server on: %s", addr)

	err := server.ListenAndServe()
	helper.FatalIfError(err)
}
