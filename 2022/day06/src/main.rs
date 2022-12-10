use std::collections::HashMap;
use std::collections::VecDeque;
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

fn find_marker(s: &String, n: usize) -> usize {
    let mut map: HashMap<char, usize> = HashMap::new();
    let mut stack = VecDeque::with_capacity(s.len());

    let inc = |map: &mut HashMap<char, usize>, c: char| {
        match map.get(&c) {
            Some(count) => { map.insert(c, count + 1); },
            _ => { map.insert(c, 1); },
        };
    };

    return s.char_indices().fold_while(0, |_, (i, c)| {
        if stack.len() < n {
            stack.push_back(c);
            inc(&mut map, c);

            return FoldWhile::Continue(i);
        }

        if map.len() == n {
            return FoldWhile::Done(i);
        }

        stack.push_back(c);
        inc(&mut map, c);

        let popped = stack.pop_front().unwrap();
        let count = map.get(&popped).unwrap();
        if *count > 1 {
            map.insert(popped, count - 1);
        } else {
            map.remove(&popped);
        }

        FoldWhile::Continue(i)
    }).into_inner()
}

#[cfg(test)]
mod tests {
    use super::*;

    macro_rules! test_marker {
        ($name:ident, $($line:expr, $seq:expr, $expected:expr),+) => {
            #[test]
            fn $name() {
                $({
                    let (actual, expected) = (find_marker(&$line.to_string(), $seq), $expected);
                    assert_eq!(actual, expected, "find_marker({}, {})", $line, $seq);
                })*
            }
        }
    }

    macro_rules! test_start_of_packet {
        ($($line:expr, $expected:expr),+) => {
            test_marker!(test_find_marker_4, $($line, 4, $expected),*);
        }
    }

    macro_rules! test_start_of_message {
        ($($line:expr, $expected:expr),+) => {
            test_marker!(test_find_marker_14, $($line, 14, $expected),*);
        }
    }

    test_start_of_packet!(
        "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7,
        "bvwbjplbgvbhsrlpgdmjqwftvncz", 5,
        "nppdvjthqldpwncqszvftbrmjlhg", 6,
        "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10,
        "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11
    );

    test_start_of_message!(
        "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19,
        "bvwbjplbgvbhsrlpgdmjqwftvncz", 23,
        "nppdvjthqldpwncqszvftbrmjlhg", 23,
        "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29,
        "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26
    );
}

fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let line = input(reader);
    println!("part1: {:?}", find_marker(&line, 4));
    println!("part2: {:?}", find_marker(&line, 14));
}
