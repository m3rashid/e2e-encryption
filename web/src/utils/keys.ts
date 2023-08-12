// import { Buffer } from 'buffer';

// export const convertPrivateKeyToBase64: (
//   key: CryptoKey
// ) => Promise<string> = async (key) => {

// };

export const genKey: () => Promise<{
  publicKey: string;
  privateKey: string;
}> = async () => {
  const keyPair = await crypto.subtle.generateKey(
    {
      name: 'RSA-OAEP',
      modulusLength: 4096,
      publicExponent: new Uint8Array([1, 0, 1]),
      hash: 'SHA-256',
    },
    true,
    ['encrypt', 'decrypt']
  );

  const publicKey = await crypto.subtle.exportKey('jwk', keyPair.publicKey);
  const privateKey = await crypto.subtle.exportKey('jwk', keyPair.privateKey);

  return {
    publicKey: btoa(JSON.stringify(publicKey)),
    privateKey: btoa(JSON.stringify(privateKey)),
  };
};
