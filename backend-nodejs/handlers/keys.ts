import { Request, Response } from 'express';
import { generateRsaKeyPair } from '../crypto/keys';

export type Session = { clientPublicKey: string; serverPrivateKey: string };

export const session = new Map<string, Session>();

export const handleExchangeKeys = async (req: Request, res: Response) => {
  try {
    const { publicKey, clientId } = req.body;
    if (!publicKey || !clientId) {
      return res.status(400).json({ error: 'Missing public key or client ID' });
    }

    const { privateKey, publicKey: serverPublickey } =
      await generateRsaKeyPair();

    session.set(clientId, {
      clientPublicKey: publicKey,
      serverPrivateKey: privateKey,
    });

    return res.json({ publicKey: serverPublickey });
  } catch (err: any) {
    console.log(err);
    return res.status(500).json({ error: 'Internal server error' });
  }
};
