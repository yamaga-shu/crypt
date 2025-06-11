use xor::{bits, bool};

fn main() {
    let a = true;
    let b = false;

    println!("Boolean XOR between {} and {} is {}", a, b, bool::xor(a, b));
    // Output:
    // Boolean XOR between true and false is true

    let c = 0b1101_0011;
    let d = 0b1010_1000;

    println!(
        "Bits XOR between {:08b} and {:08b} is {:08b}",
        c,
        d,
        bits::xor(c, d)
    )
    // Output:
    // Bits XOR between 11010011 and 10101000 is 01111011
}
