import crypto from 'crypto';

export const generateRsaKeyPair: () => Promise<{
  publicKey: string;
  privateKey: string;
}> = async () => {
  return new Promise((resolve, reject) => {
    crypto.generateKeyPair(
      'rsa',
      {
        modulusLength: 4096,
        publicKeyEncoding: { type: 'spki', format: 'pem' },
        privateKeyEncoding: { type: 'pkcs8', format: 'pem' },
      },
      (err, publicKey, privateKey) => {
        if (err) return reject(err);
        return resolve({ publicKey, privateKey });
      }
    );
  });
};
