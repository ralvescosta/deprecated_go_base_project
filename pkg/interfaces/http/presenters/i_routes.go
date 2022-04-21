package presenters

import httpServer "github.com/ralvescosta/base/pkg/infra/http_server"

type IRoutes interface {
	Register(httpServer httpServer.IHTTPServer)
}
