package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MigratorTestSuite struct {
	suite.Suite
}

func TestMigratorTestSuite(t *testing.T) {
	suite.Run(t, new(MigratorTestSuite))
}

func (s *MigratorTestSuite) TestListMigrations() {
	migrations, err := ListMigrations()

	s.NoError(err)
	s.NotNil(migrations)
}
