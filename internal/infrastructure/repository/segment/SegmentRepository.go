package segment

import (
	"context"
	"github.com/google/uuid"
	"github.com/itpark/market/dco/internal/config/db"
	"github.com/itpark/market/dco/internal/domain"
	"github.com/itpark/market/dco/internal/infrastructure/repository/segment/dto"
	"github.com/itpark/market/dco/internal/telemetry/logging"
	sqlx_utils "github.com/itpark/market/dco/pkg/sqlx-utils"
)

type SegmentRepository struct {
	DbConnection *db.DbConnection
}

func NewSegmentRepository(dbConnection *db.DbConnection) *SegmentRepository {
	return &SegmentRepository{
		DbConnection: dbConnection,
	}
}

func (r *SegmentRepository) CreateSegment(ctx context.Context, model *domain.Segments) error {
	query := `INSERT INTO dco.segments (title, group_id, p, response) VALUES ($1, $2, $3, $4)`

	tx, err := r.DbConnection.DB.BeginTx(ctx, nil)

	if err != nil {
		logging.Error("Failed to begin transaction: %v", err)
		return err
	}

	_, err = tx.ExecContext(ctx, query, model.Title, model.GroupId, model.P, model.Response)

	if err != nil {
		logging.Error("Failed to create segment: ", err)
		return sqlx_utils.HandleRollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *SegmentRepository) GetAll(ctx context.Context) ([]dto.GetAllSegmentsDbResult, error) {
	var data []dto.GetAllSegmentsDbResult
	query := `select s.title, g.title as group_name, s.p, s.response from dco.segments s join dco.groups g on s.group_id = g.id`

	err := r.DbConnection.DB.SelectContext(ctx, &data, query)

	if err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return data, nil
}

func (r *SegmentRepository) AttachUserToSegment(ctx context.Context, clientId string, segmentId uuid.UUID) error {
	query := `insert into dco.clients_segments(client_id, segment_id) values ($1, $2)`
	tx, err := r.DbConnection.DB.BeginTx(ctx, nil)
	if err != nil {
		logging.Error("Failed to begin transaction: %v", err)
		return err
	}
	_, err = tx.ExecContext(ctx, query, clientId, segmentId)
	if err != nil {
		logging.Error("Failed to attach user to segment: %v", err)
		return sqlx_utils.HandleRollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *SegmentRepository) GetSegmentByClientIdAndGroupTitle(ctx context.Context, clientId, groupTitle string) (domain.Segments, error) {
	var result domain.Segments

	query := `select * from dco.clients_segments cs
join dco.segments s on cs.segment_id = s.id
join dco.groups g on s.group_id = g.id
where cs.client_id = $1 and g.title = $2`

	err := r.DbConnection.DB.SelectContext(ctx, &result, query, clientId, groupTitle)

	if err != nil {
		logging.Error(err.Error())
		return domain.Segments{}, err
	}

	return result, nil
}

func (r *SegmentRepository) GetSegmentsByGroupTitle(ctx context.Context, groupTitle string) ([]domain.Segments, error) {
	var result []domain.Segments

	query := `select s.group_id, s.id, s.p, s.response, s.title from dco.groups g join dco.segments s on g.id = s.group_id where g.title = $1`

	err := r.DbConnection.DB.SelectContext(ctx, &result, query, groupTitle)

	if err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return result, nil
}
