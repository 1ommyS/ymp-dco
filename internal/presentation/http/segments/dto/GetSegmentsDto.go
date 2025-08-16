package dto

import (
	segments_db_dto "github.com/itpark/market/dco/internal/infrastructure/repository/segment/dto"
)

type GetSegmentsDto struct {
	Title     string `json:"title"`
	P         uint32 `json:"p"`
	GroupName string `json:"group_name"`
	Response  string `json:"response"`
}

func NewGetSegmentDtoFromModel(segment segments_db_dto.GetAllSegmentsDbResult) GetSegmentsDto {
	return GetSegmentsDto{
		Title:     segment.Title,
		P:         segment.P,
		Response:  segment.Response,
		GroupName: segment.GroupName,
	}
}

func NewGetSegmentDtoListFromModel(segments []segments_db_dto.GetAllSegmentsDbResult) []GetSegmentsDto {
	var dtos []GetSegmentsDto
	for _, segment := range segments {
		segmentDto := NewGetSegmentDtoFromModel(segment)
		dtos = append(dtos, segmentDto)
	}
	return dtos
}
