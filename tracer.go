package tracer

import (
	"github.com/shogo82148/go-sql-proxy"
	"database/sql"
	"log"
)

type Logger struct {}

func (*Logger) Output(n int, out string) (err error) {
	log.Println(out)
	return err
}

func init() {
	logger := new(Logger)
	for _, driver := range(sql.Drivers()) {
		db, _ := sql.Open(driver, "")
		defer db.Close()
		sql.Register(driver + "-trace", proxy.NewTraceProxy(db.Driver(), logger))
	}
}
