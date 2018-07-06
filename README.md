[![Build Status](https://travis-ci.com/compo-io/db.svg?branch=master)](https://travis-ci.com/compo-io/db)
[![codecov](https://codecov.io/gh/compo-io/db/branch/master/graph/badge.svg)](https://codecov.io/gh/compo-io/db)
[![Go Report Card](https://goreportcard.com/badge/github.com/compo-io/db)](https://goreportcard.com/report/github.com/compo-io/db)

# db

Very simple database package based on [reform](https://github.com/go-reform/reform).

```
import "github.com/compo-io/db"
```

## Usage

```go
var (

	// ErrNoRows is returned when no rows are found
	ErrNoRows = reform.ErrNoRows

	// ErrInitialized is returned when the database is already initialized
	ErrInitialized = errors.New("database already initialized")
)
```

#### func  Get

```go
func Get() *reform.DB
```
Get returns database connection if the package was initialized using Init()
method, returns nil otherwise. Make sure your are calling Init() before any
queries.

#### func  Init

```go
func Init(dsn *string, driver *string) error
```
Init is initializing the database using optional dsn and driver. It will use
default values if nil is passed. This method can only be called once, or will
return an error ErrInitialized.

## Example
```go
package main

import (
    "log"

    "github.com/compo-io/db"
)

func main() {
    if err := db.Init(nil,nil); err != nil {
        log.Fatal(err)
    }

    db.Get().Query("SELECT * FROM users WHERE id = ?", 1)
}
```