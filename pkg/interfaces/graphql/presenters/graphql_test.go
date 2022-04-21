package presenters

import (
	"testing"

	graphqlserver "github.com/ralvescosta/base/pkg/infra/graphql_server"
	httpServer "github.com/ralvescosta/base/pkg/infra/http_server"
	"github.com/stretchr/testify/suite"
)

type GraphQLPresenterSuit struct {
	suite.Suite

	sut           GraphqlRoutes
	httpServer    httpServer.HTTPServerSpy
	graphqlServer graphqlserver.GraphqlServerSpy
}

func TestGraphQLPresenterTestSuit(t *testing.T) {
	suite.Run(t, new(GraphQLPresenterSuit))
}

func (pst *GraphQLPresenterSuit) SetupTest() {
	pst.httpServer = httpServer.HTTPServerSpy{}
	pst.graphqlServer = graphqlserver.GraphqlServerSpy{}
	pst.sut = GraphqlRoutes{}
}

func (pst *GraphQLPresenterSuit) TestRegsiterExecuteCorrectly() {
	pst.httpServer.On("RegisterRoute", "POST", "/api/gql/query").Return(nil).Once()
	pst.httpServer.On("RegisterRoute", "GET", "/api/gql/subscriptions").Return(nil).Once()
	pst.httpServer.On("RegisterRoute", "GET", "/api/gql/playground").Return(nil).Once()

	pst.sut.Register(pst.httpServer, pst.graphqlServer)

	pst.httpServer.AssertExpectations(pst.T())
}
