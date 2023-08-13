### Simple End-to-End Encryption Implementation

Very simple end to end encryption implementation using AES-256-CBC and RSA-4096.

#### Usage

1. On initial request, Client creates a random RSA public-private key pair and sends its public key to the server.
2. The server takes this key, stores this key (as part of the session) and sends its own public key to the client.
3. The client stores the public key sent by the server (as part of the session) in memory.
4. For any request, client sends the request body encrypted with server's public key.
5. The server decrypts the request body using its private key and processes the request.
6. The server sends the response body encrypted with client's public key.
7. In this way, the client decrypts the response body using its private key.
8. The client and server can now communicate securely. Not even the browser devtools can see the request/response body.
