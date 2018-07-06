package db

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/suite"
	"gopkg.in/reform.v1"
)

func TestDB(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}

type DBTestSuite struct {
	suite.Suite
}

func (t *DBTestSuite) TestInit() {
	t.False(initialized)
	dsn1 := "http://unknown:wrong:dsn2"
	driver1 := "unknown_driver"
	t.Error(Init(&dsn1, &driver1))
	dsn2 := "DBTestSuite"
	driver2 := "ramsql"
	t.NoError(Init(&dsn2, &driver2))
	err := Init(nil, nil)
	t.Error(err)
	t.Equal(ErrInitialized, err)
}

func (t *DBTestSuite) TestGet() {
	if initialized {
		t.NotNil(Get())
		t.IsType(new(reform.DB), Get())
	} else {
		t.Nil(Get())
	}
}

func (t *DBTestSuite) TestGetDsn() {
	file := path.Join(os.TempDir(), "some_db_secret_file_for_tests")
	t.NoError(ioutil.WriteFile(file, []byte("some data"), 0644))
	dsn1 := getDsn()
	t.Equal(defaultDsn+"?"+dsnParams, dsn1)
	dsn2 := getDsn(file)
	t.Equal("some data?"+dsnParams, dsn2)
}

func (t *DBTestSuite) TestFormDsn() {
	t.Equal("dsn?"+dsnParams, formDsn("dsn"))
	t.Equal("dsn?a=b&"+dsnParams, formDsn("dsn?a=b"))
}
