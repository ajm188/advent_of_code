use std::env;
use std::io::BufRead;

use itertools::Itertools;
use itertools::FoldWhile;

use lib::io;

fn input(mut reader: Box<dyn BufRead>) -> String {
    let mut line = String::new();
    reader.read_line(&mut line).expect("failed to read input line");

    line
}

fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let line = input(reader);
    let marker = line.chars().tuple_windows().fold_while(0, |count, (a, b, c, d)| {
        let ret = count + 1;

        if a == b || a == c || a == d {
            FoldWhile::Continue(ret)
        } else if b == c || b == d {
            FoldWhile::Continue(ret)
        } else if c == d {
            FoldWhile::Continue(ret)
        } else {
            FoldWhile::Done(ret + 3)
        }
    }).into_inner();
    println!("part1: {:?}", marker);
}
