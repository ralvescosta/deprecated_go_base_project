package graphqlserver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/suite"
)

type GraphqlServerSuit struct {
	suite.Suite

	sut                 *graphqlServer
	gqlHandlerServerSpy *GqlHandlerServerSpy
}

func TestGraphqlserverSuit(t *testing.T) {
	suite.Run(t, new(GraphqlServerSuit))
}

func (pst *GraphqlServerSuit) SetupTest() {
	pst.gqlHandlerServerSpy = NewGqlHandlerServerSpy()
	pst.sut = &graphqlServer{pst.gqlHandlerServerSpy}
}

func (pst *GraphqlServerSuit) TestNewGraphQLServer() {
	s := NewGraphQLServer(pst.gqlHandlerServerSpy)

	pst.IsType(graphqlServer{}, s)
}

func (pst *GraphqlServerSuit) TestDefaultExecCorrectly() {
	pst.gqlHandlerServerSpy.On("AddTransport", ws)
	pst.gqlHandlerServerSpy.On("AddTransport", transport.Options{})
	pst.gqlHandlerServerSpy.On("AddTransport", transport.GET{})
	pst.gqlHandlerServerSpy.On("AddTransport", transport.POST{})
	pst.gqlHandlerServerSpy.On("AddTransport", transport.MultipartForm{})
	pst.gqlHandlerServerSpy.On("SetQueryCache", lru.New(100))
	pst.gqlHandlerServerSpy.On("Use", extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	pst.gqlHandlerServerSpy.On("Use", extension.Introspection{})

	pst.sut.Default()

	pst.gqlHandlerServerSpy.AssertExpectations(pst.T())
}

func (pst *GraphqlServerSuit) TestServeHTTP() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	pst.gqlHandlerServerSpy.On("ServeHTTP", res, req)

	pst.sut.ServeHTTP(res, req)

	pst.gqlHandlerServerSpy.AssertExpectations(pst.T())
}

func (pst *GraphqlServerSuit) TestInitFunc() {
	pst.sut.initFunc(context.Background(), transport.InitPayload{})
}
