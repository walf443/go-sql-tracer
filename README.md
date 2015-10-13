# go-sql-tracer
## Example

```
  package main

  import (
    database/sql
	"github.com/go-sql-driver/mysql" # load driver before load go-sql-tracer
    _ "github.com/walf443/go-sql-tracer"
  )

  func main() {
    db, err := sql.Open('mysql-trace', dsn)
  }
```
