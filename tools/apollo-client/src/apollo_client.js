import { split, HttpLink } from '@apollo/client';
import { getMainDefinition } from '@apollo/client/utilities';
import { ApolloClient, InMemoryCache,  } from '@apollo/client';
import { WebSocketLink } from '@apollo/client/link/ws';

const BASE_URL = "localhost:3333/api/gql"

const wsLink = new WebSocketLink({
  uri: `wss://${BASE_URL}/subscriptions`,
  options: {
    reconnect: true,
    connectionParams: {
      accessToken: "token"
    }
  },
});

const httpLink = new HttpLink({
  uri: `https://${BASE_URL}/query`
});

const link = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink,
);

export const client = new ApolloClient({
  link,
  cache: new InMemoryCache()
});