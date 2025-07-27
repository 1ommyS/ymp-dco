package group

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
	groupRepository "github.com/itpark/market/dco/internal/infrastructure/repository"
	customErrors "github.com/itpark/market/dco/internal/presentation/http/common"
	"github.com/itpark/market/dco/internal/presentation/http/group/dto"
	"github.com/itpark/market/dco/internal/service"
	"net/http"
)

type Handler struct {
	Service *service.GroupService
}

func Init(connection *db.DbConnection) *Handler {
	repository := groupRepository.NewGroupRepository(connection)
	newGroupService := service.NewGroupService(repository)

	return &Handler{
		Service: newGroupService,
	}
}

func (g *Handler) CreateGroup(ctx *gin.Context) {
	var groupDto dto.CreateGroupDto
	if err := ctx.ShouldBindJSON(&groupDto); err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid Request body", err))
		return
	}

	g.Service.CreateGroup(ctx, groupDto)

	ctx.Status(http.StatusCreated)
}

func (g *Handler) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, g.Service.GetAll(ctx))
}
