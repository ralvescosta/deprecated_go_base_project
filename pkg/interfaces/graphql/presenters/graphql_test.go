package presenters

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	graphqlserver "github.com/ralvescosta/base/pkg/infra/graphql_server"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
	"github.com/ralvescosta/base/pkg/infra/logger"
	"github.com/stretchr/testify/suite"
)

type GraphQLPresenterSuit struct {
	suite.Suite

	sut           GraphqlRoutes
	httpServer    *httpServer.HTTPServerSpy
	graphqlServer *graphqlserver.GraphqlServerSpy
}

func TestGraphQLPresenterTestSuit(t *testing.T) {
	suite.Run(t, new(GraphQLPresenterSuit))
}

func (pst *GraphQLPresenterSuit) SetupTest() {
	pst.httpServer = httpServer.NewHTTPServerSpy()
	pst.graphqlServer = graphqlserver.NewGraphqlServerSpy()
	pst.sut = GraphqlRoutes{}
}

func (pst *GraphQLPresenterSuit) TestRegsiterExecuteCorrectly() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	pst.httpServer.On("RegisterRoute", "POST", "/api/v1/gql/query").Return(nil).Once()
	pst.httpServer.On("RegisterRoute", "GET", "/api/v1/gql/subscriptions").Return(nil).Once()
	pst.httpServer.On("RegisterRoute", "GET", "/api/v1/gql/playground").Return(nil).Once()
	pst.graphqlServer.On("ServeHTTP", c.Writer, c.Request)

	pst.sut.Register(pst.httpServer, pst.graphqlServer)

	for _, h := range pst.httpServer.Handlers {
		h(c)
	}

	pst.httpServer.AssertExpectations(pst.T())
	pst.graphqlServer.AssertExpectations(pst.T())
}

func (pst *GraphQLPresenterSuit) TestNewGraphqlPresenter() {
	routes := NewGraphQLRoutes(logger.NewLoggerSpy())
	pst.IsType(GraphqlRoutes{}, routes)
}
