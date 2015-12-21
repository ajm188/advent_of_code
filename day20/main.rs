#![feature(iter_arith)]
fn presents_received(house: u32) -> u32 {
    let factors = (1..house + 1).filter(|i| house % i == 0);
    factors.map(|i| i * 10).sum()
}

fn main() {
    let amount = 36000000;
    // In the best case, a number is divisible by every number from 1 to itself
    // (n). In this case, it would receive 10 (1 + 2 + ... + n) presents, or
    // 10 ( n (n + 1) / 2), or 5 (n^2 + n). The i for which this quantity is at
    // least this amount, is the first house for which this is even possible.
    let lower = (1..).find(|i| 5 * (i * i + i) >= amount).unwrap();
    let house = (lower..).find(|&house| presents_received(house) >= amount);
    println!("{}", house.unwrap());
}
