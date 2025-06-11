use xor::{bits, bool};

fn main() {
    let plain_bool = true;
    let key_bool = false;

    // encrypt
    let ciphered_bool = bool::xor(plain_bool, key_bool);
    println!(
        "Boolean XOR between {} and {} is {}",
        plain_bool, key_bool, ciphered_bool
    );
    // Output:
    // Boolean XOR between true and false is true

    // decrypt
    let decrypted_bool = bool::xor(ciphered_bool, key_bool);
    println!(
        "decrypted_bool == plain_bool: {}",
        decrypted_bool == plain_bool
    );
    // Output:
    // decrypted_bool == plain_bool: true

    let plain_bits = 0b1101_0011;
    let key_bits = 0b1010_1000;

    // encrypt
    let ciphered_bits = bits::xor(plain_bits, key_bits);
    println!(
        "Bits XOR between {:08b} and {:08b} is {:08b}",
        plain_bits, key_bits, ciphered_bits
    );
    // Output:
    // Bits XOR between 11010011 and 10101000 is 01111011

    // decrypt
    let decrypted_bits = bits::xor(ciphered_bits, key_bits);
    println!(
        "decripted_bits == plain_bits: {}",
        decrypted_bits == plain_bits
    );
    // Output:
    // decripted_bits == plain_bits: true
}
