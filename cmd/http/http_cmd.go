package http

import (
	"log"

	"github.com/ralvescosta/base/pkg/infra/environments"

	"github.com/spf13/cobra"
)

func NewHTTPServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "GoLang Base Application HTTP Server Command",
		Run: func(cmd *cobra.Command, args []string) {
			env := environments.NewEnvironment()
			if err := env.Configure(); err != nil {
				log.Fatal(err)
			}

			container, err := NewHTTPContainer(env)
			if err != nil {
				log.Fatal(err)
			}

			container.httpServer.Default()
			container.marketsRoutes.Register(container.httpServer)
			container.httpServer.Setup()

			if err := container.httpServer.Run(); err != nil {
				log.Fatal(err)
			}
		},
	}
}
