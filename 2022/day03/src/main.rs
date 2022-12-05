use std::collections::HashSet;
use std::env;
use std::io::BufRead;

use lib::io;


fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let halves = reader.lines().map(|line| 
        line.unwrap()
    ).map(|line| -> (Vec<char>, Vec<char>) {
        (line.chars().into_iter().take(line.len() / 2).collect(), line.chars().into_iter().skip(line.len() / 2).collect())
    });

    let commons = halves.map(|(l, r)| {
        let mut lhs = HashSet::new();
        for c in &l {
            lhs.insert(*c);
        }

        let mut rhs = HashSet::new();
        for c in &r {
            rhs.insert(*c);
        }

        let mut i: Vec<char> = Vec::new();
        for c in lhs.intersection(&rhs).take(1) {
            i.insert(0, *c);
        }

        i[0].to_owned()
    });
    
    let part1 = commons.map(|c: char| -> i32 {
        let base = c as i32;
        match c {
            'a'..='z' => base - ('a' as i32) + 1,
            'A'..='Z' => base - ('A' as i32) + 1 + 26,
            _ => base,
        }
    })/*.map(|x| { println!("{}", x); x })*/.sum::<i32>();
    
    println!("part1: {:?}", part1);
}
