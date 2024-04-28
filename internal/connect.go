package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func Connect(f *YamlConfig) error {
	var err error
	var dns string = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", f.Host, f.Port, f.User, f.Password, f.Dbname, f.Sslmode)
	db, err = sql.Open("postgres", dns)

	if err != nil {
		return err
	}

	return db.Ping()
}
