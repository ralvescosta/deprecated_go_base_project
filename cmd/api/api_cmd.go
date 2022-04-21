package api

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ralvescosta/base/pkg/infra/environments"
)

func NewHTTPServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "GoLang Base Application API Server Command",
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
			container.graphqlServer.Default()
			container.marketsRoutes.Register(container.httpServer)
			container.graphqlRoutes.Register(container.httpServer, container.graphqlServer)
			container.httpServer.Setup()

			if err := container.httpServer.Run(); err != nil {
				log.Fatal(err)
			}
		},
	}
}
