package environments

import (
	"fmt"
	"os"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"

	"github.com/ralvescosta/dotenv"
)

var dotEnvConfig = dotenv.Configure

type env struct {
	_GO_ENV      string
	_DEV_ENV     string
	_STAGING_ENV string
	_PROD_ENV    string
	_PROF_ENV    string
}

func (pst env) Configure() error {
	err := dotEnvConfig(fmt.Sprintf(".env.%s", pst._GO_ENV))
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	return nil
}

func (pst env) GO_ENV() string {
	return pst._GO_ENV
}

func (pst env) DEV_ENV() string {
	return pst._DEV_ENV
}

func (pst env) STAGING_ENV() string {
	return pst._STAGING_ENV
}

func (pst env) PROD_ENV() string {
	return pst._PROD_ENV
}

func (pst env) PROFILING_ENV() string {
	return pst._PROF_ENV
}

func NewEnvironment() interfaces.IEnvironments {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		goEnv = "development"
	}

	profEnv := os.Getenv("APP_PROFILING")
	if profEnv == "" {
		profEnv = "disabled"
	}

	return env{
		_GO_ENV:      goEnv,
		_PROF_ENV:    profEnv,
		_DEV_ENV:     "development",
		_STAGING_ENV: "staging",
		_PROD_ENV:    "production",
	}
}
