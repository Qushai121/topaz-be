package documentdto

import "github.com/Qushai121/topaz-be/dto"

type GetDocumentListQueryParamsDto struct {
	dto.BaseListQueryParamsDto
}

type DocumentListItem struct {
	ID         uint   `query:"id"`
	Name       string `query:"name"`
	ContentRaw string `query:"contentRaw"`
}
