use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

fn count_diff(text: &String) -> usize {
    let hex_re = Regex::new(r"\\x[a-f0-9]{2}").unwrap();
    let no_hex = hex_re.replace_all(&text, " ");
    let esc_re = Regex::new(r"\\.").unwrap();
    let no_esc = esc_re.replace_all(&no_hex, " ");
    no_esc.len()
}

fn count_escaped(text: &String) -> usize {
    text.chars()
        .fold(
            0 as usize,
            |acc, c| acc + match c {
                '"' | '\\' => 2,
                _          => 1,
            }
        )
}

fn main() {
    let stdin = io::stdin();
    let input = stdin.lock().lines();
    let lens = input.map(|res| {
        let s = res.unwrap();
        (s.len(), count_diff(&s) - 2, count_escaped(&s) + 2)
    });
    let totals = lens.fold(
        (0, 0, 0),
        |acc, e| (acc.0 + e.0, acc.1 + e.1, acc.2 + e.2)
    );
    println!("({}, {} {})", totals.0, totals.1, totals.2);
    println!("{}", totals.0 - totals.1);
    println!("{}", totals.2 - totals.0);
}
