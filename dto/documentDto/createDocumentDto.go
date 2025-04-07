package documentdto

type DocumentRequestBodyDto struct {
	CategoryId uint `json:"categoryId" validate:"required"`
	UserId     uint
	Name       string `json:"name" validate:"required,min=3,max=50"`
	ContentRaw string `json:"contentRaw" validate:"required,min=3"`
}