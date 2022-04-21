package graphqlserver

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type GraphqlServerSpy struct {
	mock.Mock
}

func (m GraphqlServerSpy) Default() {}

func (m GraphqlServerSpy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
