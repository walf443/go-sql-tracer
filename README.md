# go-sql-tracer
## Example

```
  package main

  import (
    database/sql
    _ "github.com/walf443/go-sql-tracer"
  )

  func main() {
    db, err := sql.Open('mysql-trace', dsn)
  }
```
