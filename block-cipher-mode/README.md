# Overview of Block Cipher Modes

| Name | Merit | Demerit | Remarks |
|------|-------|---------|---------|
| ECB (Electronic Codebook) | Simple and fast encryption of individual blocks. | Identical plaintext blocks produce identical ciphertext blocks, revealing patterns. | Not recommended. |
| CBC (Cipher Block Chaining) | Enhances security by chaining blocks, making pattern detection difficult. | Requires padding and is vulnerable to certain attacks if not implemented correctly. | Recommended for general use, but with caution regarding padding and IV management. Recommended by [CRYPTREC / Practical Cryptography] |
| CFB (Cipher Feedback) | Allows encryption of data in units smaller than the block size, suitable for varying data sizes. | Error propagation: a single bit error in ciphertext affects subsequent blocks. | Useful for network streams and real-time data encryption. Recommended by [CRYPTREC] |
| OFB (Output Feedback) | Converts block cipher into a stream cipher, resistant to error propagation. | Requires careful management of IV to ensure security. | Suitable for noisy communication channels and pre-processing of data. Recommended by [CRYPTREC] |
| CTR (Counter) | Highly parallelizable, efficient for high-speed applications, and allows random access to encrypted data. | Requires a unique nonce and counter for each encryption to maintain security. | Recommended for high-performance applications. Recommended by [CRYPTREC / Practical Cryptography] |
