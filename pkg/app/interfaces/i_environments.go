package interfaces

type IEnvironments interface {
	Configure() error
	GO_ENV() string
	DEV_ENV() string
	STAGING_ENV() string
	PROD_ENV() string
}
