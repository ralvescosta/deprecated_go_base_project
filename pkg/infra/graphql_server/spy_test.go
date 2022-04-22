package graphqlserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/stretchr/testify/suite"
)

type GraphqlServerSpySuit struct {
	suite.Suite

	server  GraphqlServerSpy
	handler GqlHandlerServerSpy
}

func TestGraphqlServerSpyTestSuit(t *testing.T) {
	suite.Run(t, new(GraphqlServerSpySuit))
}

func (pst *GraphqlServerSpySuit) SetupTest() {
	pst.server = GraphqlServerSpy{}
	pst.handler = GqlHandlerServerSpy{}
}

func (pst *GraphqlServerSpySuit) TestServerDefault() {
	pst.server.Default()
}

func (pst *GraphqlServerSpySuit) TestServerServeHTTP() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	pst.server.On("ServeHTTP", res, req)

	pst.server.ServeHTTP(res, req)
}

func (pst *GraphqlServerSpySuit) TestServerNewGraphqlServerSpy() {
	spy := NewGraphqlServerSpy()
	pst.IsType(&GraphqlServerSpy{}, spy)
}

func (pst *GraphqlServerSpySuit) TestHandlerNewGqlHandlerServerSpy() {
	spy := NewGqlHandlerServerSpy()
	pst.IsType(&GqlHandlerServerSpy{}, spy)
}

func (pst *GraphqlServerSpySuit) TestHandler() {
	t := transport.GET{}
	pst.handler.On("AddTransport", t)
	pst.handler.AddTransport(t)

	c := lru.New(10)
	pst.handler.On("SetQueryCache", c)
	pst.handler.SetQueryCache(c)

	h := extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	}
	pst.handler.On("Use", h)
	pst.handler.Use(h)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	pst.handler.On("ServeHTTP", res, req)
	pst.handler.ServeHTTP(res, req)

	pst.handler.AssertExpectations(pst.T())
}
