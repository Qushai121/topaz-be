package services

import (
	"net/http"

	"github.com/Qushai121/topaz-be/dto"
	documentdto "github.com/Qushai121/topaz-be/dto/documentDto"
	"github.com/Qushai121/topaz-be/models"
	"gorm.io/gorm"
)

type IDocumentService interface {
	GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[[]documentdto.DocumentListItem]], *dto.ErrorDto[any])
}

type documentService struct {
	dbTopaz *gorm.DB
}

func NewDocumentService(dbTopaz *gorm.DB) IDocumentService {
	return &documentService{
		dbTopaz: dbTopaz,
	}
}

func (s *documentService) GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[[]documentdto.DocumentListItem]], *dto.ErrorDto[any]) {

	var documentList dto.PaginateDto[[]documentdto.DocumentListItem]

	res := s.dbTopaz.Model(&models.Document{}).Find(&documentList.Data)

	if res.Error != nil {
		return nil, dto.NewErrorDto[any]("Get document list failed", int(http.StatusInternalServerError), nil)
	}

	return dto.NewSuccessDto("Get document list successfully", int(http.StatusOK), documentList), nil
}
