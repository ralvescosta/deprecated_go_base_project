package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	apm "go.elastic.co/apm/module/apmgin/v2"
)

type IHTTPServer interface {
	Default()
	RegisterRoute(method string, path string, handlers ...gin.HandlerFunc) error
	Setup()
	Run() error
}

type HTTPServer struct {
	env      interfaces.IEnvironments
	addr     string
	logger   interfaces.ILogger
	router   *gin.Engine
	server   *http.Server
	shotdown chan bool
}

var httpServerWrapper = gin.New

func (pst *HTTPServer) Default() {
	pst.router = httpServerWrapper()
	pst.router.Use(GinLogger(pst.logger))
	pst.router.Use(apm.Middleware(pst.router)) //apm also carry about the recovery
	pst.router.SetTrustedProxies(nil)
}

func (hs HTTPServer) RegisterRoute(method string, path string, handlers ...gin.HandlerFunc) error {
	switch method {
	case "POST":
		hs.router.POST(path, handlers...)
	case "GET":
		hs.router.GET(path, handlers...)
	case "PUT":
		hs.router.PUT(path, handlers...)
	case "PATCH":
		hs.router.PATCH(path, handlers...)
	case "DELETE":
		hs.router.DELETE(path, handlers...)
	default:
		return errors.NewInternalError("http method not allowed")
	}
	return nil
}

func (pst *HTTPServer) Setup() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	pst.addr = fmt.Sprintf("%s:%s", host, port)

	pst.server = &http.Server{
		Addr:    pst.addr,
		Handler: pst.router,
	}

	go pst.gracefullShutdown()
}

func (pst HTTPServer) Run() error {
	if pst.env.GO_ENV() != pst.env.PROD_ENV() {
		certPath := os.Getenv("TLS_CERT_PATH")
		keyPath := os.Getenv("TLS_KEY_PATH")

		pst.logger.Info(fmt.Sprintf("[HttpServer::Run] - Server running at: https://%s", pst.addr))
		err := pst.server.ListenAndServeTLS(certPath, keyPath)

		return errors.NewInternalError(err.Error())
	}

	pst.logger.Info(fmt.Sprintf("[HttpServer::Run] - Server running at: http://%s", pst.addr))
	err := pst.server.ListenAndServe()

	return errors.NewInternalError(err.Error())
}

func (pst HTTPServer) gracefullShutdown() {
	<-pst.shotdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pst.server.Shutdown(ctx); err != nil {
		pst.logger.Error("[HttpServer::GracefullShutdown] - could'ent shutdown properly")
		return
	}
}

func NewHTTPServer(environments interfaces.IEnvironments, logger interfaces.ILogger, shotdown chan bool) IHTTPServer {
	return &HTTPServer{
		env:      environments,
		logger:   logger,
		shotdown: shotdown,
	}
}
