use std::collections::HashSet;
use std::env;
use std::io::BufRead;

use lib::io;
use lib::set;

fn item_priority(c: char) -> i32 {
    let offset = match c {
        'a'..='z' => ('a' as i32) - 1,
        'A'..='Z' => ('A' as i32) - 1 - 26,
        _ => c as i32,
    };

    (c as i32) - offset
}

fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let halves = reader.lines().map(|line| 
        line.unwrap()
    ).map(|line| -> (Vec<char>, Vec<char>) {
        (line.chars().into_iter().take(line.len() / 2).collect(), line.chars().into_iter().skip(line.len() / 2).collect())
    });

    let commons = halves.map(|(l, r)| {
        let sets = vec![
            HashSet::from_iter(l.into_iter()),
            HashSet::from_iter(r.into_iter()),
        ];
        let items: Vec<&char> = set::intersect_all(&sets).into_iter().collect();

        *items[0]
    });
    
    /*.map(|x| { println!("{}", x); x })*/
    let part1 = commons.map(item_priority).sum::<i32>();
    
    println!("part1: {:?}", part1);
}
