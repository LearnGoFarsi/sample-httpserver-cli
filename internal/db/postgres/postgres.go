package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"text/template"
	"time"

	_ "github.com/lib/pq"

	"github.com/learngofarsi/go-basics-project/pkg/config"
)

const connString = "postgres://{{.User}}:{{.Pass}}@{{.Host}}:{{.Port}}/{{.DbName}}?sslmode=disable"

func buildConnectionStringOrPanic(cnf config.Postgres) string {

	sb := strings.Builder{}
	temp := template.Must(template.New("ConnString").Parse(connString))
	if err := temp.Execute(&sb, cnf); err != nil {
		panic(err)
	}

	return sb.String()
}

func NewPostgres(cnf config.Postgres) (*sql.DB, error) {

	conn := buildConnectionStringOrPanic(cnf)

	// pg driver: github.com/lib/pq
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Print("Failed to ping the database")
		return db, fmt.Errorf("Could not ping the database %w", err)
	}

	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)

	return db, nil
}
