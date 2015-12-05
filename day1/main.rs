#![feature(iter_arith)]

use std::env::args;

fn main() {
    let instructions = match args().nth(1) {
        Some(v) => v,
        None => "".to_string(),
    };
    let floor: i32 = instructions
        .chars()
        .map(|c: char| if c == '(' { 1 } else { -1 })
        .sum();
    println!("{}", floor);
}
