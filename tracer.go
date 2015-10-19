package tracer

import (
	"github.com/shogo82148/go-sql-proxy"
	"database/sql"
	"log"
	"strings"
)

type Logger struct {
	output func(int, string) error
}
var logger *Logger

func (l *Logger) Output(n int, out string) (err error) {
	l.output(n, out)
	return err
}

func ChangeOutput(f func(int, string) error) {
	logger.output = f
}

func init() {
	logger = new(Logger)
	logger.output = func(n int, out string) error {
	  var err error
	  log.Println(out)
	  return err
	}
	for _, driver := range(sql.Drivers()) {
		if strings.Contains(driver, ":trace") {
			continue
		}
		db, _ := sql.Open(driver, "")
		defer db.Close()
		sql.Register(driver + ":trace", proxy.NewTraceProxy(db.Driver(), logger))
	}
}

