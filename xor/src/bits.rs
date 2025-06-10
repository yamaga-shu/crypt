pub fn xor(a: u8, b: u8) -> u8 {
    a ^ b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_bits() {
        assert_eq!(xor(0b1101_0011, 0b1010_1000), 0b0111_1011);
        assert_eq!(xor(0b1111_1111, 0b0000_0000), 0b1111_1111);
        assert_eq!(xor(0b0000_0000, 0b1111_1111), 0b1111_1111);
        assert_eq!(xor(0b1010_1010, 0b0101_0101), 0b1111_1111);
    }
}
