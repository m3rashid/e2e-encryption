import { useEffect } from 'react';
import { genKey } from './utils/keys';

const App = () => {
  useEffect(() => {
    genKey()
      .then(({ privateKey, publicKey }) => {
        console.log({ privateKey, publicKey });
      })
      .catch(console.log);
    // create a public private key pair
    // send public key to the server
    // server will send their own public key
    // store the keys
    // use keys in further communication
  }, []);

  return <div className='bg-red-500'>hello</div>;
};

export default App;
