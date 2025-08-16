package group

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

func (repo *GroupRepository) CreateGroup(ctx context.Context, name string) error {
	query := `
		INSERT INTO dco.groups (title) VALUES ($1)
		`

	tx, err := repo.DbConnection.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, name)
	if err != nil {
		logging.Error(err.Error())
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			return err
		}

		return err
	}
	rowsAffected, _ := res.RowsAffected()
	logging.Debug(fmt.Sprintf("Created %d rows", rowsAffected))

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repo *GroupRepository) GetAll(ctx *gin.Context) ([]domain.Group, error) {
	var groups []domain.Group
	err := repo.DbConnection.DB.SelectContext(ctx, &groups, "SELECT id, title from dco.groups")
	if err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return groups, nil
}
