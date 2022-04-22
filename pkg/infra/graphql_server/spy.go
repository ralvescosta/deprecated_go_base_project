package graphqlserver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/mock"
)

type GraphqlServerSpy struct {
	mock.Mock
}

func (m GraphqlServerSpy) Default() {}

func (m GraphqlServerSpy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func NewGraphqlServerSpy() *GraphqlServerSpy {
	return new(GraphqlServerSpy)
}

type GqlHandlerServerSpy struct {
	mock.Mock
}

func (pst *GqlHandlerServerSpy) AddTransport(t graphql.Transport) {
	pst.Called(t)
}

func (pst *GqlHandlerServerSpy) SetQueryCache(c graphql.Cache) {
	pst.Called(c)
}

func (pst *GqlHandlerServerSpy) Use(h graphql.HandlerExtension) {
	pst.Called(h)
}

func (pst *GqlHandlerServerSpy) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	pst.Called(res, req)
}

func NewGqlHandlerServerSpy() *GqlHandlerServerSpy {
	return new(GqlHandlerServerSpy)
}
