package main

import (
	"github.com/ralvescosta/base/cmd"
	"github.com/ralvescosta/base/cmd/api"
	"github.com/ralvescosta/base/cmd/migrator"
)

func main() {
	cmd.Execute(
		migrator.NewMigratorCmd(),
		api.NewHTTPServerCmd(),
	)
}
