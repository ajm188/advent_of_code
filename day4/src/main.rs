extern crate openssl;

use openssl::crypto::hash::{hash, Type};

fn hash_with_five_leading_zeros(input: String) -> bool {
    five_leading_zeros(hash(Type::MD5, input.as_bytes()))
}

fn five_leading_zeros(md5: Vec<u8>) -> bool {
    md5.iter().take(2).all(|x| x + 0 == 0) &&
        md5.iter().skip(2).take(1).all(|x| x + 0 < 10 as u8)
}

fn main() {
    let base = "bgvyzdsv";
    let num = (1..).find(|i|
                         hash_with_five_leading_zeros(String::from(base) + &(i.to_string())));
    println!("{}", num.unwrap());
}
