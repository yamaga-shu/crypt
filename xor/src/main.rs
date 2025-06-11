use xor::{bits, bool};

fn main() {
    let plain_bool = true;
    let key_bool = false;

    // encrypt
    let cipher_bool = bool::xor(plain_bool, key_bool);
    println!(
        "Boolean XOR between {} and {} is {}",
        plain_bool, key_bool, cipher_bool
    );
    // Output:
    // Boolean XOR between true and false is true

    // decrypt
    let decrypted_bool = bool::xor(cipher_bool, key_bool);
    println!(
        "Decrypted boolean: {:#?} using key {:#?} is {}",
        cipher_bool, key_bool, decrypted_bool
    );
    // Output:
    // Decrypted boolean: true using key false is true

    let plain_bits = 0b1101_0011;
    let key_bits = 0b1010_1000;

    // encrypt
    let cipher_bits = bits::xor(plain_bits, key_bits);
    println!(
        "Bits XOR between {:08b} and {:08b} is {:08b}",
        plain_bits, key_bits, cipher_bits
    );
    // Output:
    // Bits XOR between 11010011 and 10101000 is 01111011

    // decrypt
    let decrypted_bits = bits::xor(cipher_bits, key_bits);
    println!(
        "Decrypted bits: {:08b} using key {:08b} is {:08b}",
        cipher_bits, key_bits, decrypted_bits
    );
    // Output:
    // Decrypted bits: 01111011 using key 10101000 is 11010011
}
