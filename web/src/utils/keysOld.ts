import { Buffer } from 'buffer';

export const convertPrivateKeyToBase64 = async (key: CryptoKey) => {
  const exported = await crypto.subtle.exportKey('pkcs8', key);
  const buff = Buffer.from(exported).toString('base64');
  return buff;
};

export const genKey = async () => {
  const keys = await window.crypto.subtle.generateKey(
    {
      name: 'RSA-OAEP',
      modulusLength: 2048,
      publicExponent: new Uint8Array([0x01, 0x00, 0x01]),
      hash: { name: 'SHA-256' },
    },
    true,
    ['encrypt', 'decrypt']
  );

  const privateKey = await convertPrivateKeyToBase64(keys.privateKey);
  const publicKey = keys.publicKey;
  return { privateKey, publicKey };
};
