package services

import (
	"github.com/Qushai121/topaz-be/dto"
	"gorm.io/gorm"
)

type IDocumentService interface {
	GetDocumentList() *dto.ErrorDto[any]
}

type documentService struct{
	dbTopaz *gorm.DB
}

func NewDocumentService(dbTopaz *gorm.DB) IDocumentService {
	return &documentService{
		dbTopaz: dbTopaz,
	}
}

func (s *documentService) GetDocumentList() *dto.ErrorDto[any] {
	return dto.InternalServerError()
}
