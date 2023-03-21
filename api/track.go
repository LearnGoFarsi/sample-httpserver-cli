package api

type Track struct {
	Id     *int64  `db:"track_id" json:"id"`
	Name   *string `db:"track_name" json:"name"`
	Artist *string `db:"artist" json:"artist"`
	Length *int    `db:"track_length" json:"length"`
}
