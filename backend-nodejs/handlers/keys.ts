import { Request, Response } from 'express';
import { generateRsaKeyPair } from '../crypto/keys';

export const handleExchangeKeys = async (req: Request, res: Response) => {
  try {
    const { publicKey } = req.body;
    if (!publicKey) {
      return res.status(400).json({ error: 'Missing public key' });
    }
    const { privateKey, publicKey: serverPublickey } =
      await generateRsaKeyPair();
    console.log({ privateKey, serverPublickey });
    return res.json({ publicKey: serverPublickey });
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ error: 'Internal server error' });
  }
};
