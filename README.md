# Overview
[![Twitter](https://img.shields.io/badge/author-%40MachielMolenaar-blue.svg)](https://twitter.com/MachielMolenaar)
[![GoDoc](https://godoc.org/github.com/Machiel/querybag?status.svg)](https://godoc.org/github.com/Machiel/querybag)
[![Build Status](https://travis-ci.org/Machiel/querybag.svg?branch=master)](https://travis-ci.org/Machiel/querybag)
[![Coverage Status](https://coveralls.io/repos/Machiel/querybag/badge.svg?branch=master&service=github)](https://coveralls.io/github/Machiel/querybag?branch=master)

Querybag is a library that loads all SQL files from a specific directory, and it
allows you to access them by using their file name, as in the example below.

# License
Querybag is licensed under a MIT license.

# Installation
A simple `go get github.com/Machiel/querybag` should suffice.

# Usage

## Example

`queries/retrieve_user_by_id.sql`

```sql
SELECT *
FROM user
WHERE id = ?
```

```go
package main

import (
  "fmt"
  "database/sql"

  "github.com/Machiel/querybag"
)

var db *sql.DB

func main() {

  userID := 12
  b, err := querybag.New("queries") // Load all SQL files from queries directory

  if err != nil {
    panic(err)
  }

  rows, err := db.Query(b.Get("retrieve_user_by_id"), userID)

  // Business as usual

}
```
