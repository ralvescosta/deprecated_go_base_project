package presenters

import httpServer "markets/pkg/infra/http_server"

type IRoutes interface {
	Register(httpServer httpServer.IHTTPServer)
}
