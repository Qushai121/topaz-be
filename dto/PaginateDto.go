package dto

type PaginateDto[T any] struct {
	Data        T   `json:"data,omitempty"`
	TotalRecord int64 `json:"totalRecord"`
}
