package initialize

import (
	"database/sql"

	"github.com/LeMinh0706/lala-song/util"
)

func Postgres(config util.Config) (*sql.DB, error) {
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
