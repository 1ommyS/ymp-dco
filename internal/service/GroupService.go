package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/domain"
	"github.com/itpark/market/dco/internal/infrastructure/repository"
	"github.com/itpark/market/dco/internal/presentation/http/group/dto"
)

type GroupService struct {
	Repository *repository.GroupRepository
}

func NewGroupService(repository *repository.GroupRepository) *GroupService {
	return &GroupService{
		Repository: repository,
	}
}

func (groupService *GroupService) CreateGroup(ctx context.Context, groupDto dto.CreateGroupDto) {
	groupService.Repository.CreateGroup(ctx, groupDto.Name)
}

func (groupService *GroupService) GetAll(ctx *gin.Context) []domain.Group {
	return groupService.Repository.GetAll(ctx)
}
