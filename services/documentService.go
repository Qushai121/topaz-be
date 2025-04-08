package services

import (
	"net/http"

	"github.com/Qushai121/topaz-be/configs"
	"github.com/Qushai121/topaz-be/dto"
	documentdto "github.com/Qushai121/topaz-be/dto/documentDto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/Qushai121/topaz-be/models"
	"github.com/Qushai121/topaz-be/utils"

	// "github.com/Qushai121/topaz-be/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IDocumentService interface {
	GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[documentdto.DocumentListItem]], *dto.ErrorDto[any])
	CreateDocument(ctx *fiber.Ctx, body *documentdto.DocumentRequestBodyDto, localDataUser *entities.UserTokenPayload) (*dto.SuccessDto[any], *dto.ErrorDto[any])
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

func (s *documentService) GetDocumentList(queryParams *documentdto.GetDocumentListQueryParamsDto) (*dto.SuccessDto[dto.PaginateDto[documentdto.DocumentListItem]], *dto.ErrorDto[any]) {

	var documentList dto.PaginateDto[documentdto.DocumentListItem]

	query := s.topazDb.Model(&models.Document{})

	query = queryParams.GetBaseQueryParamsToDbQuery(query, &documentList.TotalRecord)

	query.Find(&documentList.Data)

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Get document list failed", int(http.StatusInternalServerError), nil)
	}

	return dto.NewSuccessDto("Get document list successfully", int(http.StatusOK), documentList), nil
}

func (s *documentService) CreateDocument(ctx *fiber.Ctx, body *documentdto.DocumentRequestBodyDto, localDataUser *entities.UserTokenPayload) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {
	// Save model document
	document := models.Document{
		Name:               body.Name,
		ContentRaw:         body.ContentRaw,
		CategoryDocumentId: *body.GetCategoryIdToUint(),
		UserId:             localDataUser.UserId,
	}

	query := s.topazDb.Create(&document)

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Create new document failed", http.StatusInternalServerError, nil)
	}
	// Save model document

	// Safe File Functionality
	imagePath, errImagePath := utils.SaveFileToPath(body.Image, utils.DocumenFileDir, ctx)

	if errImagePath != nil {
		return nil, dto.NewErrorDto[any]("Create new document failed", http.StatusInternalServerError, nil)
	}

	fileStorage := models.FileStorage{
		Filename: body.Image.Filename,
		FilePath: *imagePath,
	}

	queryFile := s.topazDb.Model(&document).Association("FileStorage").Append(&fileStorage)

	if queryFile != nil {
		s.topazDb.Delete(&document)
		utils.RemoveFileFromPath(*imagePath)
		return nil, dto.NewErrorDto[any]("Create new document failed", http.StatusInternalServerError, nil)
	}
	// Safe File Functionality

	return dto.NewSuccessDto[any]("Create new document successfully", http.StatusCreated, nil), nil
}

func (s *documentService) DeleteDocument(documentId uint, userId uint) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {

	query := s.topazDb.Where("id = ? AND user_id = ?", documentId, userId).Delete(models.Document{})

	if query.Error != nil {
		return nil, dto.NewErrorDto[any]("Delete document failed", http.StatusInternalServerError, nil)
	}

	return dto.NewSuccessDto[any]("Delete document successfully", http.StatusCreated, nil), nil
}

func (s *documentService) UpdateDocument(documentId uint, userId uint, body *documentdto.DocumentRequestBodyDto) (*dto.SuccessDto[any], *dto.ErrorDto[any]) {

	query := s.topazDb.Model(&models.Document{}).Where("id = ? AND user_id = ?", documentId, userId).Updates(body)

	if query.Error != nil || query.RowsAffected == 0 {
		return nil, dto.NewErrorDto[any]("Update document failed", http.StatusInternalServerError, nil)
	}

	return dto.NewSuccessDto[any]("Update document successfully", http.StatusCreated, nil), nil
}
