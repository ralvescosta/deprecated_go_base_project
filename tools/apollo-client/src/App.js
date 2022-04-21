import { useEffect } from 'react';
import { ApolloProvider } from '@apollo/client';
import { gql } from '@apollo/client';
import { useSubscription } from '@apollo/client';
import { client } from './apollo_client'
import './App.css';

const Component = () => {
    const { data } = useSubscription(
    gql`
    subscription {
      marketCreated {
        long
        lat
        registro
      }
    }
  `);

  useEffect(() => {
    console.log(data)
  }, [data])

  return (
    <div className="App">
      <h1>Ai Pai Para...</h1> 
    </div>
  )
}

function App() {
  return (
    <ApolloProvider client = {client}>
      <Component />
    </ApolloProvider>
  );
}

export default App;
