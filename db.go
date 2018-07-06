// Very simple database package based on [reform](https://github.com/go-reform/reform).
package db

import (
	"database/sql"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects"
)

const (
	defaultDriver = "mysql"
	defaultDsn    = "/tmp/db.db"
	dsnParams     = "parseTime=true&clientFoundRows=true"
	secret        = "/run/secrets/db_dsn"
)

var (
	initialized = false
	sqlLink     *sql.DB
	rfLink      *reform.DB

	// ErrNoRows is returned when no rows are found
	ErrNoRows = reform.ErrNoRows

	// ErrInitialized is returned when the database is already initialized
	ErrInitialized = errors.New("database already initialized")
)

// Get returns database connection if the package was initialized
// using Init() method, returns nil otherwise. Make sure your are
// calling Init() before any queries.
func Get() *reform.DB {
	return rfLink
}

// Init is initializing the database using optional dsn and driver.
// It will use default values if nil is passed. This method can only
// be called once, or will return an error ErrInitialized.
func Init(dsn *string, driver *string) error {
	if initialized {
		return ErrInitialized
	}
	useDsn := getDsn()
	if dsn != nil {
		useDsn = *dsn
	}
	useDriver := defaultDriver
	if driver != nil {
		useDriver = *driver
	}
	var err error
	sqlLink, err = sql.Open(
		useDriver,
		useDsn)
	if err != nil {
		return err
	}
	rfLink = reform.NewDB(
		sqlLink,
		dialects.ForDriver(defaultDriver),
		nil)
	initialized = true
	return nil
}

func getDsn(file ...string) string {
	secretFile := secret
	if len(file) == 1 {
		secretFile = file[0]
	}
	s, err := ioutil.ReadFile(secretFile)
	if err != nil || s == nil {
		return formDsn(defaultDsn)
	}
	return formDsn(string(s))
}

func formDsn(dsn string) string {
	if strings.Index(dsn, "?") != -1 {
		return dsn + "&" + dsnParams
	}
	return dsn + "?" + dsnParams
}
