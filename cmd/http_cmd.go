package cmd

import "markets/pkg/infra/environments"

func HTTPServer() error {
	env := environments.NewEnvironment()
	if err := env.Configure(); err != nil {
		return err
	}

	container, err := NewHTTPContainer(env)
	if err != nil {
		return err
	}

	container.httpServer.Default()
	container.marketsRoutes.Register(container.httpServer)
	container.httpServer.Setup()

	if err := container.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
