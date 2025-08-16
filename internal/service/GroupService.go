package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/infrastructure/repository/group"
	"github.com/itpark/market/dco/internal/presentation/http/group/dto"
)

type GroupService struct {
	Repository *group.GroupRepository
}

func NewGroupService(repository *group.GroupRepository) *GroupService {
	return &GroupService{
		Repository: repository,
	}
}

func (groupService *GroupService) CreateGroup(ctx context.Context, groupDto dto.CreateGroupDto) error {
	return groupService.Repository.CreateGroup(ctx, groupDto.Name)

}

func (groupService *GroupService) GetAll(ctx *gin.Context) ([]dto.GetGroupDto, error) {
	groups, err := groupService.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewGetGroupDtoListFromModel(groups), nil
}
