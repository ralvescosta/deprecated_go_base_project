package environments

import "github.com/stretchr/testify/mock"

type EnvironmentsSpy struct {
	mock.Mock
}

func (pst EnvironmentsSpy) Configure() error {
	args := pst.Called()

	return args.Error(0)
}

func (pst EnvironmentsSpy) GO_ENV() string {
	return pst.Called().String(0)
}

func (pst EnvironmentsSpy) DEV_ENV() string {
	return pst.Called().String(0)
}

func (pst EnvironmentsSpy) STAGING_ENV() string {
	return pst.Called().String(0)
}

func (pst EnvironmentsSpy) PROD_ENV() string {
	return pst.Called().String(0)
}

func NewEnvironmentsSpy() *EnvironmentsSpy {
	return new(EnvironmentsSpy)
}
