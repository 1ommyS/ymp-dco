package router

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
	groupHandlers "github.com/itpark/market/dco/internal/presentation/http/group"
)

type GroupRouter struct {
	DbConnection *db.DbConnection
}

func NewGroupRouter(DbConnection *db.DbConnection) *GroupRouter {
	return &GroupRouter{
		DbConnection: DbConnection,
	}
}

func (groupRouter *GroupRouter) RegisterRoutes(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("groups")

	{
		groupHandler := groupHandlers.Init(groupRouter.DbConnection)
		group.POST("/", groupHandler.CreateGroup)
		group.GET("/", groupHandler.FindAll)
	}

}
