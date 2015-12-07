extern crate regex;

use regex::Regex;
use std::env::args;

fn has_bad_characters(s: &String) -> bool {
    let bad_re = Regex::new(r"(?:(ab)|(cd)|(pq)|(xy))").unwrap();
    bad_re.is_match(&s)
}

fn at_least_three_vowels(s: &String) -> bool {
    // regex wasn't working here for some reason, and it's a linear pass
    // anyways
    s.chars().filter(
        |c| match *c {
            'a' | 'e' | 'i' | 'o' | 'u' => true,
            _                           => false,
        }
    ).count() >= 3
}

fn at_least_one_duplicate(s: &String) -> bool {
    let mut iter = s.chars().peekable();
    while iter.peek().is_some() {
        let c = iter.next().unwrap();
        if iter.peek().is_some() && c == *iter.peek().unwrap() {
            return true
        }
    }
    false
}

fn is_nice(s: &String) -> bool {
    !has_bad_characters(s) &&
        at_least_three_vowels(s) &&
        at_least_one_duplicate(s)
}

fn main() {
    let num = args().skip(1).filter(is_nice).count();
    println!("{}", num);
}
