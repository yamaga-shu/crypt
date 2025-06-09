fn main() {
    let a = true;
    let b = false;

    println!("XOR between {} and {} is {}", a, b, xor(a, b));
}

// XOR演算を行う関数
fn xor(a: bool, b: bool) -> bool {
    a ^ b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_xor() {
        assert_eq!(xor(true, true), false);
        assert_eq!(xor(true, false), true);
        assert_eq!(xor(false, true), true);
        assert_eq!(xor(false, false), false);
    }
}
