import { useEffect } from 'react';

// const genKey = async () => {
//   const keys = await new SubtleCrypto().generateKey({ name: 'RSA' }, true, [
// 'encrypt',
// 'decrypt',
//   ]);
//   console.log(keys);
// };

const genKey = async () => {
  const keys = await window.crypto.subtle.generateKey(
    {
      name: 'RSA-OAEP',
      modulusLength: 2048,
      publicExponent: new Uint8Array([1, 0, 1]),
      hash: 'SHA-256',
    },
    true,
    ['encrypt', 'decrypt']
  );
  console.log(keys.privateKey.exportKey('jwk'));
};

function App() {
  useEffect(() => {
    genKey();
    // create a public private key pair
    // send public key to the server
    // server will send their own public key
    // store the keys
    // use keys in further communication
  }, []);

  return <div>hello</div>;
}

export default App;
