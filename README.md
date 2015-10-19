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
    db, err := sql.Open('mysql:trace', dsn)
  }
```

## DESCRIPTION

go-sql-tracer register sql drivers that name has suffix ':trace' from your loaded drivers in package init.
So you should place loading sql driver libraries before import go-sql-tracer.

## SEE ALSO

http://github.com/shogo82148/go-sql-proxy

## Author

Copyright (c) 2015 Keiji Yoshimi

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
