export const stringifyKey: (key: CryptoKey) => Promise<string> = async (
  key
) => {
  const exported = await crypto.subtle.exportKey('jwk', key);
  return btoa(JSON.stringify(exported));
};

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

  const publicKey = await stringifyKey(keyPair.publicKey);
  const privateKey = await stringifyKey(keyPair.privateKey);

  return { publicKey, privateKey };
};
