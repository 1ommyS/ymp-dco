package group

import (
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
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
		groupHandler := InitGroupHandler(groupRouter.DbConnection)
		group.POST("/", groupHandler.CreateGroup)
		group.GET("/", groupHandler.FindAll)
	}

}
