use std::env;
use std::io::BufRead;
use std::ops::RangeInclusive;

use regex::Regex;

use lib::io;

fn fully_contains<T: Ord>(a: &RangeInclusive<T>, b: &RangeInclusive<T>) -> bool {
    vec![(a, b), (b, a)].iter().any(|(a, b)| {
        a.start() <= b.start() && a.end() >= b.end()
    })
}

fn main() {
    let assignment_re = Regex::new("([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)").unwrap();

    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));
    let assignments = reader.lines().map(
        |line| line.unwrap()
    ).map(
        |line| -> (RangeInclusive<i32>, RangeInclusive<i32>) {
            let cap = assignment_re.captures(line.as_str()).unwrap();
            (
                (cap[1].parse().unwrap())..=(cap[2].parse().unwrap()),
                (cap[3].parse().unwrap())..=(cap[4].parse().unwrap()),
            )
        }
    );

    let part1 = assignments.filter(|(r1, r2)| fully_contains(r1, r2) ).count();
    println!("part1: {:?}", part1);
}
