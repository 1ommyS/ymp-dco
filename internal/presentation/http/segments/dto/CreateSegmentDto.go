package dto

import (
	"github.com/google/uuid"
	"github.com/itpark/market/dco/internal/domain"
)

type CreateSegmentDto struct {
	Title    string    `json:"title"`
	GroupId  uuid.UUID `json:"group_id"`
	P        uint32    `json:"p"`
	Response string    `json:"response"`
}

func NewCreateSegmentDto(title string, groupId uuid.UUID, p uint32, response string) *CreateSegmentDto {
	return &CreateSegmentDto{Title: title, GroupId: groupId, P: p, Response: response}
}

func (segment CreateSegmentDto) ToModel() *domain.Segments {
	return &domain.Segments{
		Title:    segment.Title,
		GroupId:  segment.GroupId,
		P:        segment.P,
		Response: segment.Response,
	}
}
