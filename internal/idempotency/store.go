package idempotency

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) IsProcessed(
	ctx context.Context,
	service string,
	eventID string,
) (bool, error) {

	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM processed_events
			WHERE service = $1
			  AND event_id = $2
		)
	`

	err := s.db.GetContext(
		ctx,
		&exists,
		query,
		service,
		eventID,
	)

	return exists, err
}

func (s *Store) MarkProcessed(
	ctx context.Context,
	service string,
	eventID string,
) error {

	query := `
		INSERT INTO processed_events (
			service,
			event_id
		)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`

	_, err := s.db.ExecContext(
		ctx,
		query,
		service,
		eventID,
	)

	return err
}
