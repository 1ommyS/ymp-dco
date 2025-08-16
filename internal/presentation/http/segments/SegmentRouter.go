package segments

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
)

type SegmentRouter struct {
	DbConnection *db.DbConnection
}

func NewSegmentRouter(DbConnection *db.DbConnection) *SegmentRouter {
	return &SegmentRouter{
		DbConnection: DbConnection,
	}
}

func (segmentRouter *SegmentRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("segments")

	{
		segmentHandler := InitSegmentHandler(segmentRouter.DbConnection)
		group.POST("/", segmentHandler.CreateSegment)
		group.GET("/", segmentHandler.FindAll)
		group.GET("/:groupTitle/:clientId", segmentHandler.FindByGroupTitleAndClientId)
		group.POST("/:groupTitle/:clientId", segmentHandler.AttachUserToSegment)
	}

}
