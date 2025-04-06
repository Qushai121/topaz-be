package documentdto

import "github.com/Qushai121/topaz-be/dto"

type GetDocumentListQueryParamsDto struct {
	dto.BaseListQueryParamsDto
}

type DocumentListItem struct {
	Name       string `query:"name"`
	ContentRaw string `query:"contentRaw"`
}
