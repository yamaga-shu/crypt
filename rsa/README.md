## RSA Encryption

RSA (Rivest-Shamir-Adleman) is one of the first public-key cryptosystems and is widely used for secure data transmission. The security of RSA relies on the practical difficulty of factoring the product of two large prime numbers, the factoring problem.

### Key Concepts

1. **Public and Private Keys**: RSA uses a pair of keys. The public key is used for encryption, and the private key is used for decryption. The public key can be shared with anyone, while the private key must be kept secret.

2. **Encryption and Decryption**:
   - **Encryption**: A message is encrypted using the recipient's public key. This ensures that only the recipient, who has the corresponding private key, can decrypt the message.
   - **Decryption**: The recipient uses their private key to decrypt the message.

3. **Key Generation**:
   - Choose two distinct large random prime numbers \( p \) and \( q \).
   - Compute \( n = pq \). \( n \) is used as the modulus for both the public and private keys.
   - Compute the totient \(\phi(n) = (p-1)(q-1)\).
   - Choose an integer \( e \) such that \( 1 < e < \phi(n) \) and \( e \) is coprime to \(\phi(n)\). \( e \) is the public exponent.
   - Determine \( d \) as \( d \equiv e^{-1} \mod \phi(n) \). \( d \) is the private exponent.

4. **Security**: The security of RSA is based on the difficulty of factoring the large number \( n \) into its prime components. As of now, no efficient algorithm exists for factoring large numbers, making RSA secure when large keys are used.

### Applications

- **Secure Communications**: RSA is used to secure data transmission over the internet, such as in SSL/TLS protocols.
- **Digital Signatures**: RSA can be used to sign a message, providing authenticity and integrity.

### Limitations

- **Performance**: RSA is computationally intensive, which can make it slower than symmetric key algorithms.
- **Key Size**: To maintain security, RSA keys need to be large, typically 2048 bits or more.