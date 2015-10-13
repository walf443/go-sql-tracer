package tracer

import (
	"github.com/shogo82148/go-sql-proxy"
	"database/sql"
	"log"
	"strings"
)

type Logger struct {}
var logger Logger

func (*Logger) Output(n int, out string) (err error) {
	log.Println(out)
	return err
}

func init() {
	logger := new(Logger)
	for _, driver := range(sql.Drivers()) {
		if strings.Contains(driver, "-trace") {
			continue
		}
		db, _ := sql.Open(driver, "")
		defer db.Close()
		sql.Register(driver + "-trace", proxy.NewTraceProxy(db.Driver(), logger))
	}
}
