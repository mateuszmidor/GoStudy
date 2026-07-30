package outbox

import (
	"context"
	"database/sql"
	"errors"
)

// RunInTransaction executes all functions within a database transaction
func RunInTransaction(ctx context.Context, db *sql.DB, fns ...func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, fn := range fns {
		err = fn(tx)
		if err != nil {
			break
		}
	}

	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}
	return err
}
