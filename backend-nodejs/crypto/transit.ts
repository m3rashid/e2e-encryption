import crypto from 'crypto';

export const handleEncrypt = async (data: any, clientPublicKey: string) => {
  return new Promise((resolve, reject) => {
    try {
      const encrypted = crypto.publicEncrypt(
        {
          key: clientPublicKey,
          padding: crypto.constants.RSA_PKCS1_OAEP_PADDING,
          oaepHash: 'sha256',
        },
        Buffer.from(data)
      );
      resolve(encrypted.toString('base64'));
    } catch (err) {
      reject(err);
    }
  });
};

export const handleReceiveEncrypted = async (
  encryptedBase64: string,
  serverPrivateKey: string
) => {
  return new Promise((resolve, reject) => {
    try {
      const decrypted = crypto.privateDecrypt(
        {
          key: serverPrivateKey,
          padding: crypto.constants.RSA_PKCS1_OAEP_PADDING,
          oaepHash: 'sha256',
        },
        Buffer.from(encryptedBase64, 'base64')
      );
      resolve(decrypted.toString());
    } catch (err) {
      reject(err);
    }
  });
};
