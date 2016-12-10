use std::collections::HashSet;
use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

fn replace(rep: (Regex, String), molecule: &String) -> HashSet<String> {
    let (reg, r) = rep;
    let mut replacements = HashSet::new();
    for (start, end) in reg.find_iter(molecule) {
        let front = molecule[0..start].to_string();
        let back = molecule[end..].to_string();
        let repl = front + &r + &back;
        replacements.insert(repl);
    }
    replacements
}

fn parse_replacement(line: &String) -> (Regex, String) {
    let rep_re = Regex::new(r"(.*) => (.*)").unwrap();
    match rep_re.captures(line) {
        Some(caps) => (Regex::new(caps.at(1).unwrap()).unwrap(),
                       caps.at(2).unwrap().to_string()),
        None => panic!("ahh"),
    }
}

fn main() {
    let stdin = io::stdin();
    let lines: Vec<String> =
        stdin.lock().lines().map(|l| l.unwrap()).collect();
    let molecule = lines.iter().last().unwrap();
    let replacements: Vec<(Regex, String)> =
        lines.iter().take_while(|l| *l != "")
        .map(|l| parse_replacement(l)).collect();
    let mut derivatives = HashSet::new();
    for rep in replacements {
        let repl = replace(rep, &molecule);
        derivatives = repl.union(&derivatives).cloned().collect();
    }
    println!("{}", derivatives.len());
}
