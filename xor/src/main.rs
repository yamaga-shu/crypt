use xor::{bits, bool};

fn main() {
    let a = true;
    let b = false;

    println!("Boolean XOR between {} and {} is {}", a, b, bool::xor(a, b));

    let c = 0b1101_0011;
    let d = 0b1010_1000;

    println!("Bits XOR between {} and {} is {}", c, d, bits::xor(c, d))
}
