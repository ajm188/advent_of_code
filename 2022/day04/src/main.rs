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

fn overlaps<T: Ord>(a: &RangeInclusive<T>, b: &RangeInclusive<T>) -> bool {
    vec![(a, b), (b, a)].iter().any(|(a, b)| {
        (a.start() <= b.start() && a.end() >= b.start()) ||
            (a.start() >= b.start() && a.end() <= b.end())
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

    let (part1, part2) = assignments.fold((0, 0), |(p1, p2), (r1, r2)| {
        if fully_contains(&r1, &r2) {
            ((p1 + 1), (p2 + 1))
        } else if overlaps(&r1, &r2) {
            (p1, p2 + 1)
        } else {
            (p1, p2)
        }
    });
    println!("part1: {:?}", part1);
    println!("part2: {:?}", part2);
}
