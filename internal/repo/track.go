package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan"
	"github.com/learngofarsi/go-basics-project/api"
)

const (
	UPSERT_TRACK = `
			INSERT INTO tracks AS t (track_id, track_name, artist, track_length) VALUES($1, $2, $3, $4) ON CONFLICT(track_id)
			DO UPDATE
			SET (track_id, track_name, artist, track_length) = ROW (excluded.*)
			WHERE  (t.*) IS DISTINCT FROM (excluded.*)
			RETURNING *;
	`
	SELECT_ALL   = "SELECT * FROM tracks"
	SELECT_BY_ID = "SELECT * FROM tracks WHERE track_id = %s"
)

type Repo struct {
	*sql.DB
}

func NewTrackRepo(db *sql.DB) Repo {
	return Repo{db}
}

func (r Repo) Upsert(ctx context.Context, track *api.Track) (err error) {

	var tx *sql.Tx

	defer func() {
		if err == nil {
			err = tx.Commit()
		}

		if err != nil {
			if e := tx.Rollback(); e != nil {
				err = fmt.Errorf("%w failed to rollback upsert transaction %w", err, e)
			}
		}

	}()

	tx, err = r.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	_, err = tx.ExecContext(ctx, UPSERT_TRACK, *track.Id, *track.Name, *track.Artist, *track.Length)

	return
}

func (r Repo) Get(ctx context.Context) (ts []api.Track, err error) {

	rows, err := r.QueryContext(ctx, SELECT_ALL)
	if err != nil {
		return
	}
	defer rows.Close()

	err = scan.Rows(&ts, rows)

	return
}

func (r Repo) GetById(ctx context.Context, id string) (t api.Track, err error) {

	row, err := r.QueryContext(ctx, fmt.Sprintf(SELECT_BY_ID, id))
	if err != nil {
		return
	}
	defer row.Close()

	err = scan.Row(&t, row)

	return
}
