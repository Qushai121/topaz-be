package services

import (
	"net/http"

	"github.com/Qushai121/topaz-be/configs"
	"github.com/Qushai121/topaz-be/dto"
	documentdto "github.com/Qushai121/topaz-be/dto/documentDto"
	"github.com/Qushai121/topaz-be/models"
	"gorm.io/gorm"
)

type IDocumentService interface {
	GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[[]documentdto.DocumentListItem]], *dto.ErrorDto[any])
	CreateDocument(body *documentdto.DocumentRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any])
	UpdateDocument(documentId uint, userId uint, body *documentdto.DocumentRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any])
	DeleteDocument(documentId uint, userId uint) (*dto.SuccessDto[any], *dto.ErrorDto[any])
}

type documentService struct {
	topazDb *gorm.DB
}

func NewDocumentService(baseDb configs.BASEDB) IDocumentService {
	return &documentService{
		topazDb: baseDb.TOPAZDB,
	}
}

func (s *documentService) GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[[]documentdto.DocumentListItem]], *dto.ErrorDto[any]) {

	var documentList dto.PaginateDto[[]documentdto.DocumentListItem]

	query := s.topazDb.Model(&models.Document{})

	query = queryParams.GetQueryParamsToDbQuery(query, &documentList.TotalRecord)

	query.Find(&documentList.Data)

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Get document list failed", int(http.StatusInternalServerError), nil)
	}

	return dto.NewSuccessDto("Get document list successfully", int(http.StatusOK), documentList), nil
}

func (s *documentService) CreateDocument(body *documentdto.DocumentRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {

	document := models.Document{
		Name:               body.Name,
		ContentRaw:         body.ContentRaw,
		CategoryDocumentId: body.CategoryId,
		UserId:             body.UserId,
	}

	query := s.topazDb.Create(&document)

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Create new document failed", http.StatusInternalServerError, nil)
	}

	return dto.NewSuccessDto[any]("Create new document successfully", http.StatusCreated, nil), nil
}

// DeleteDocument implements IDocumentService.
func (s *documentService) DeleteDocument(documentId uint, userId uint) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {

	query := s.topazDb.Delete(models.Document{
		UserId: userId,
		Model: gorm.Model{
			ID: documentId,
		},
	})

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Delete document failed", http.StatusInternalServerError, nil)
	}

	return dto.NewSuccessDto[any]("Delete document successfully", http.StatusCreated, nil), nil
}

// UpdateDocument implements IDocumentService.
func (s *documentService) UpdateDocument(documentId uint, userId uint, body *documentdto.DocumentRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {
	body.UserId = userId

	query := s.topazDb.Model(&models.Document{}).Where("id = ? AND user_id = ?", documentId, userId).Updates(body)

	if query.Error != nil || query.RowsAffected == 0 {
		return nil, dto.NewErrorDto[any]("Update document failed", http.StatusInternalServerError, nil)
	}

	return dto.NewSuccessDto[any]("Update document successfully", http.StatusCreated, nil), nil
}
