package documentdto

import (
	"mime/multipart"
	"strconv"
)

type DocumentRequestBodyDto struct {
	CategoryId string                  `json:"categoryId" form:"categoryId" validate:"required"`
	Name       string                  `json:"name" form:"name" validate:"required,min=3,max=50"`
	ContentRaw string                  `json:"contentRaw" form:"contentRaw" validate:"required,min=3"`
	Image      *multipart.FileHeader `json:"image" form:"image" validate:"required"`
}

func (d *DocumentRequestBodyDto) GetCategoryIdToUint() *uint {
	value, err := strconv.ParseUint(d.CategoryId, 10, 64)

	if err != nil {
		return nil
	}

	uintVal := uint(value)
	return &uintVal
}
