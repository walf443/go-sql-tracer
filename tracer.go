package tracer

import (
	"github.com/shogo82148/go-sql-proxy"
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
)

type Logger struct {}

func (*Logger) Output(n int, out string) (err error) {
	log.Println(out)
	return err
}

func init() {
	logger := new(Logger)
	sql.Register("mysql-trace", proxy.NewTraceProxy(&mysql.MySQLDriver{}, logger))
	sql.Register("sqlite-trace", proxy.NewTraceProxy(&sqlite3.SQLiteDriver{}, logger))
}
