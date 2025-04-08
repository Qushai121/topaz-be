package utils

import (
	"log"
	"mime/multipart"
	"reflect"

	"github.com/Qushai121/topaz-be/entities"
	"github.com/gofiber/fiber/v2"
)

func ParseMultipartRequest[T any](ctx *fiber.Ctx, data *T, fileFields *[]entities.FileField) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	// Parse regular fields
	if errParser := ctx.BodyParser(data); errParser != nil {
		log.Println(errParser.Error())
		return errParser
	}

	// No files to assign
	if fileFields == nil {
		return nil
	}

	// Use reflection to set file fields manually
	dataVal := reflect.ValueOf(data).Elem()
	dataType := dataVal.Type()

	for _, fileField := range *fileFields {
		files := form.File[fileField.Field]

		if fileField.IsMultiple {
			for i := 0; i < dataType.NumField(); i++ {
				field := dataType.Field(i)
				if field.Tag.Get("form") == fileField.Field && field.Type == reflect.TypeOf([]*multipart.FileHeader{}) {
					dataVal.Field(i).Set(reflect.ValueOf(files))
					break
				}
			}

		} else {
			file := files[0]

			// Find the struct field and set it
			for i := 0; i < dataType.NumField(); i++ {
				field := dataType.Field(i)
				if field.Tag.Get("form") == fileField.Field && field.Type == reflect.TypeOf(&multipart.FileHeader{}) {
					dataVal.Field(i).Set(reflect.ValueOf(file))
					break
				}
			}
		}

	}

	return nil
}
