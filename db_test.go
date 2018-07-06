package db

import (
	"testing"

	_ "github.com/proullon/ramsql/driver"
	"github.com/stretchr/testify/suite"
)

func TestDB(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}

type DBTestSuite struct {
	suite.Suite
}

