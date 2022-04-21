package graphqlserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GraphqlServerSpySuit struct {
	suite.Suite

	sut GraphqlServerSpy
}

func TestGraphqlServerSpyTestSuit(t *testing.T) {
	suite.Run(t, new(GraphqlServerSpySuit))
}

func (pst *GraphqlServerSpySuit) SetupTest() {
	pst.sut = GraphqlServerSpy{}
}

func (pst *GraphqlServerSpySuit) TestDefault() {
	pst.sut.Default()
}

func (pst *GraphqlServerSpySuit) TestServeHTTP() {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	pst.sut.On("ServeHTTP", res, req)

	pst.sut.ServeHTTP(res, req)
}

func (pst *GraphqlServerSpySuit) TestNewGraphqlServerSpy() {
	spy := NewGraphqlServerSpy()
	pst.IsType(&GraphqlServerSpy{}, spy)
}
