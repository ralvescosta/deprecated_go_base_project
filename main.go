package main

import (
	"github.com/ralvescosta/base/cmd"
	"github.com/ralvescosta/base/cmd/api"
	"github.com/ralvescosta/base/cmd/seeders"
)

func main() {
	cmd.Execute(
		seeders.NewSeedersCmd(),
		api.NewHTTPServerCmd(),
	)
}
