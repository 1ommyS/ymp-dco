package http

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
	"github.com/itpark/market/dco/internal/presentation/http/group"
	"github.com/itpark/market/dco/internal/presentation/http/segments"
)

func RegisterRoutes(engine *gin.Engine, db *db.DbConnection) *gin.Engine {
	api := engine.Group("/api/v1")

	{
		api.GET("/health", healthCheck)
	}

	groupRouter := group.NewGroupRouter(db)
	segmentRouter := segments.NewSegmentRouter(db)

	groupRouter.RegisterRoutes(api)
	segmentRouter.RegisterRoutes(api)

	return engine
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
