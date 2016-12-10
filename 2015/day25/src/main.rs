use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

struct CodeGenerator { code: u64, }
impl Iterator for CodeGenerator {
    type Item = u64;
    fn next(&mut self) -> Option<Self::Item> {
        self.code *= 252533;
        self.code %= 33554393;
        Some(self.code)
    }
}

fn parse(input: &String) -> (u32, u32) {
    let format = Regex::new(r".*row (\d+), column (\d+).*").unwrap();
    let caps = match format.captures(input) {
        Some(c) => c,
        None    => panic!("could not parse input"),
    };
    let row = u32::from_str_radix(caps.at(1).unwrap(), 10).unwrap();
    let col = u32::from_str_radix(caps.at(2).unwrap(), 10).unwrap();
    (row, col)
}

fn main() {
    let stdin = io::stdin();
    let input = stdin.lock().lines().next().unwrap();
    let (target_row, target_col) = parse(&input.unwrap());
    let mut code = 20151125;
    let mut generator = CodeGenerator { code: code, };
    let mut row = 1;
    let mut col = 1;
    loop {
        println!("({}, {}); {}", row, col, code);
        if row == target_row && col == target_col {
            break;
        }
        code = generator.next().unwrap();
        if row == 1 {
            row = col + 1;
            col = 1;
        } else {
            row -= 1;
            col += 1;
        }
    }
}
