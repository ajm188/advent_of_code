#![feature(str_char)]
extern crate regex;

use regex::Regex;
use std::env::args;

fn has_bad_characters(s: &String) -> bool {
    let bad_re = Regex::new(r"(?:(ab)|(cd)|(pq)|(xy))").unwrap();
    bad_re.is_match(&s)
}

fn at_least_three_vowels(s: &String) -> bool {
    let vowel_re = Regex::new(r"(.*[aeiou].*){3}").unwrap();
    vowel_re.is_match(s)
}

fn at_least_one_duplicate(s: &String) -> bool {
    (0..(s.len() - 1)).any(|i| s.char_at(i) == s.char_at(i + 1))
}

fn is_nice_part_one(s: &String) -> bool {
    !has_bad_characters(s) &&
        at_least_three_vowels(s) &&
        at_least_one_duplicate(s)
}

fn is_nice_part_two(s: &String) -> bool {
    sandwiched(s) && repeat_pair(s)
}

fn sandwiched(s: &String) -> bool {
    // rust only supports RE2 regexes, so not backreffing on capture groups
    // otherwise this can just be: (.+).\1
    (1..(s.len() - 1)).any(|i| s.char_at(i - 1) == s.char_at(i + 1))
}

fn repeat_pair(s: &String) -> bool {
    // posix regex: (.)(.).*\1\2
    (0..(s.len() - 3)).any(
        |i| {
            let c1 = s.char_at(i);
            let c2 = s.char_at(i + 1);
            ((i + 2)..(s.len() - 1)).any(
                |j| {
                    let c3 = s.char_at(j);
                    let c4 = s.char_at(j + 1);
                    c1 == c3 && c2 == c4
                }
            )
        }
    )
}

fn main() {
    let num_part_one = args().skip(1).filter(is_nice_part_one).count();
    let num_part_two = args().skip(1).filter(is_nice_part_two).count();
    println!("{} {}", num_part_one, num_part_two);
}
