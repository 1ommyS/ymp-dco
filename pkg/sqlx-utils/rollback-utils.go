package sqlx_utils

import (
	"database/sql"
	"github.com/itpark/market/dco/internal/telemetry/logging"
)

func HandleRollback(tx *sql.Tx, oldError error) error {
	err := tx.Rollback()

	if err != nil {
		logging.Error("Failed to rollback transaction: %v", err)
		return err
	}
	return oldError
}
