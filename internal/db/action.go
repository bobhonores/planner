package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bobhonores/planner/internal/action"
	"github.com/google/uuid"
)

type ActionRow struct {
	ID          string
	Name        sql.NullString
	Description sql.NullString
	Done        sql.NullBool
	Created     sql.NullTime
	Updated     sql.NullTime
}

func (d *Database) Get(
	ctx context.Context,
	id string,
) (action.Action, error) {
	var actionRow ActionRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, name, description, done, created, updated
		FROM actions
		WHERE id = $1`,
		id,
	)
	err := row.Scan(
		&actionRow.ID,
		&actionRow.Name,
		&actionRow.Description,
		&actionRow.Done,
		&actionRow.Created,
		&actionRow.Updated,
	)
	if err != nil {
		return action.Action{}, fmt.Errorf("error fetching the action by uuid: %w", err)
	}

	actionID, err := uuid.Parse(actionRow.ID)
	if err != nil {
		return action.Action{}, fmt.Errorf("error parsing id: %w", err)
	}

	return action.Action{
		ID:          actionID,
		Name:        actionRow.Name.String,
		Description: actionRow.Description.String,
		Done:        actionRow.Done.Bool,
		Created:     actionRow.Created.Time,
		Updated:     actionRow.Updated.Time,
	}, nil
}

func (d *Database) Insert(
	ctx context.Context,
	act action.Action,
) (action.Action, error) {
	actionRow := ActionRow{
		ID:          act.ID.String(),
		Name:        sql.NullString{String: act.Name, Valid: true},
		Description: sql.NullString{String: act.Description, Valid: true},
		Done:        sql.NullBool{Bool: act.Done, Valid: true},
		Created:     sql.NullTime{Time: act.Created, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO actions
		(id, name, description, done, created)
		VALUES
		(:id, :name, :description, :done, :created)`,
		actionRow,
	)

	if err != nil {
		return action.Action{}, fmt.Errorf("failed to insert action: %w", err)
	}
	if err := rows.Close(); err != nil {
		return action.Action{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return act, nil
}

func (d *Database) Delete(
	ctx context.Context,
	id string,
) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM actions WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete action: %w", err)
	}

	return nil
}

func (d *Database) Update(
	ctx context.Context,
	act action.Action,
) (action.Action, error) {
	actRow := ActionRow{
		ID:          act.ID.String(),
		Name:        sql.NullString{String: act.Name, Valid: true},
		Description: sql.NullString{String: act.Description, Valid: true},
		Done:        sql.NullBool{Bool: act.Done, Valid: true},
		Updated:     sql.NullTime{Time: act.Updated, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE actions
		SET
		name = :name,
		description = :description,
		done = :done,
		updated = :updated
		WHERE id = :id`,
		actRow,
	)

	if err != nil {
		return action.Action{}, fmt.Errorf("failed to update action: %w", err)
	}
	if err := rows.Close(); err != nil {
		return action.Action{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return action.Action{
		ID:          act.ID,
		Name:        act.Name,
		Description: act.Description,
		Done:        act.Done,
		Created:     act.Created,
		Updated:     act.Updated,
	}, nil
}
