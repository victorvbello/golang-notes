package mariadb

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MariDBConfig struct {
	URL      string
	DBName   string
	User     string
	Port     int
	Password string
}

func SQLOpenConnection(conf MariDBConfig) (*sqlx.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.User, conf.Password, conf.URL, strconv.Itoa(conf.Port), conf.DBName)
	return sqlx.Connect("mysql", conn)
}

func SQLCloseConnection(db *sqlx.DB) error {
	return db.Close()
}
