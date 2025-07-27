package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itpark/market/dco/internal/config/db"
	"github.com/itpark/market/dco/internal/domain"
	"github.com/itpark/market/dco/internal/telemetry/logging"
)

type GroupRepository struct {
	DbConnection *db.DbConnection
}

func NewGroupRepository(dbConnection *db.DbConnection) *GroupRepository {
	return &GroupRepository{
		DbConnection: dbConnection,
	}
}

func (repo *GroupRepository) CreateGroup(ctx context.Context, name string) {
	query := `
		INSERT INTO dco.groups (title) VALUES ($1)
		`

	res, err := repo.DbConnection.DB.ExecContext(ctx, query, name)
	if err != nil {
		logging.Error(err.Error())
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))
}

func (repo *GroupRepository) GetAll(ctx *gin.Context) []domain.Group {
	var groups []domain.Group
	err := repo.DbConnection.DB.SelectContext(ctx, &groups, "SELECT id, title from dco.groups")
	if err != nil {
		logging.Error(err.Error())
	}
	return groups
}
