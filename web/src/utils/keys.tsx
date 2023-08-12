import { Buffer } from 'buffer';

export const convertKeyToBase64 = async (key: CryptoKey) => {
  const exported = await crypto.subtle.exportKey('pkcs8', key);
  const buff = Buffer.from(exported).toString('base64');
  return buff.toString();
};

export const genKey = async () => {
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

  const privateKey = await convertKeyToBase64(keys.privateKey);
  const publicKey = await convertKeyToBase64(keys.publicKey);
  return { privateKey, publicKey };
};
