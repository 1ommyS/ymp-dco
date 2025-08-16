package dto

import "github.com/itpark/market/dco/internal/domain"

type GetGroupDto struct {
	Title string `json:"title"`
}

func NewGetGroupDtoFromModel(group *domain.Group) *GetGroupDto {
	return &GetGroupDto{
		Title: group.Title,
	}
}

func NewGetGroupDtoListFromModel(groups []domain.Group) []GetGroupDto {
	var dtos []GetGroupDto
	for group := range groups {
		groupDto := NewGetGroupDtoFromModel(&groups[group])
		dtos = append(dtos, *groupDto)
	}
	return dtos
}
