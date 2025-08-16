package segments

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
	segmentRepository "github.com/itpark/market/dco/internal/infrastructure/repository/segment"
	customErrors "github.com/itpark/market/dco/internal/presentation/http/common"
	"github.com/itpark/market/dco/internal/presentation/http/segments/dto"
	"github.com/itpark/market/dco/internal/service"
	"net/http"
)

type Handler struct {
	Service *service.SegmentService
}

func InitSegmentHandler(connection *db.DbConnection) *Handler {
	repository := segmentRepository.NewSegmentRepository(connection)
	newSegmentService := service.NewSegmentService(repository)

	return &Handler{
		Service: newSegmentService,
	}
}

func (h *Handler) CreateSegment(ctx *gin.Context) {
	var dto dto.CreateSegmentDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Invalid Request body", err))
		return
	}
	err := h.Service.CreateSegmentDto(ctx, dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Failed to create segment", err))
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) FindAll(ctx *gin.Context) {
	segments, err := h.Service.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, customErrors.CreateError("Failed to get segments", err))
	}

	ctx.JSON(http.StatusOK, segments)
}

func (h *Handler) FindByGroupTitleAndClientId(context *gin.Context) {
	groupTitle := context.Param("groupTitle")
	clientId := context.Param("clientId")

	segment, err := h.Service.GetSegmentByClientIdAndGroupTitle(context, groupTitle, clientId)

	if err != nil {
		context.JSON(http.StatusNotFound, customErrors.CreateError("Segment not found", nil))
		return
	}

	context.JSON(http.StatusOK, segment.Response)
}

func (h *Handler) AttachUserToSegment(context *gin.Context) {
	groupTitle := context.Param("groupTitle")
	clientId := context.Param("clientId")

	_, err := h.Service.AttachUserToSegment(context, groupTitle, clientId)

	if err != nil {
		context.JSON(http.StatusNotFound, customErrors.CreateError("Segment not found", nil))
		return
	}

	context.Status(http.StatusOK)
	return
}
