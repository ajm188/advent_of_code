use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::{Captures, Regex};

#[derive(Clone)]
struct Sue {
    id: u32,
    children: Option<u32>,
    cats: Option<u32>,
    samoyeds: Option<u32>,
    pomeranians: Option<u32>,
    akitas: Option<u32>,
    vizslas: Option<u32>,
    goldfish: Option<u32>,
    trees: Option<u32>,
    cars: Option<u32>,
    perfumes: Option<u32>,
}

fn parse(line: &String) -> Sue {
    let sue_re = Regex::new(r"Sue (\d+):").unwrap();
    let children_re = Regex::new(r"children: (\d+)").unwrap();
    let cats_re = Regex::new(r"cats: (\d+)").unwrap();
    let samoyeds_re = Regex::new(r"samoyeds: (\d+)").unwrap();
    let pomeranians_re = Regex::new(r"pomeranians: (\d+)").unwrap();
    let akitas_re = Regex::new(r"akitas: (\d+)").unwrap();
    let vizslas_re = Regex::new(r"vizslas: (\d+)").unwrap();
    let goldfish_re = Regex::new(r"goldfish: (\d+)").unwrap();
    let trees_re = Regex::new(r"trees: (\d+)").unwrap();
    let cars_re = Regex::new(r"cars: (\d+)").unwrap();
    let perfumes_re = Regex::new(r"perfumes: (\d+)").unwrap();

    let id_str = sue_re.captures(line).unwrap().at(1).unwrap();
    let id = u32::from_str_radix(id_str, 10).unwrap();
    let children = wrap_capture(children_re.captures(line));
    let cats = wrap_capture(cats_re.captures(line));
    let samoyeds = wrap_capture(samoyeds_re.captures(line));
    let pomeranians = wrap_capture(pomeranians_re.captures(line));
    let akitas = wrap_capture(akitas_re.captures(line));
    let vizslas = wrap_capture(vizslas_re.captures(line));
    let goldfish = wrap_capture(goldfish_re.captures(line));
    let trees = wrap_capture(trees_re.captures(line));
    let cars = wrap_capture(cars_re.captures(line));
    let perfumes = wrap_capture(perfumes_re.captures(line));
    Sue {
        id: id,
        children: children,
        cats: cats,
        samoyeds: samoyeds,
        pomeranians: pomeranians,
        akitas: akitas,
        vizslas: vizslas,
        goldfish: goldfish,
        trees: trees,
        cars: cars,
        perfumes: perfumes,
    }
}

fn wrap_capture(captures: Option<Captures>) -> Option<u32> {
    match captures {
        Some(c) => Some(u32::from_str_radix(c.at(1).unwrap(), 10).unwrap()),
        None    => None,
    }
}

fn main() {
    let stdin = io::stdin();
    let sues: Vec<Sue> =
        stdin
        .lock()
        .lines()
        .map(|line| parse(&line.unwrap()))
        .collect();
    let my_sue = Sue {
        id: 0,
        children: Some(3),
        cats: Some(7),
        samoyeds: Some(2),
        pomeranians: Some(3),
        akitas: Some(0),
        vizslas: Some(0),
        goldfish: Some(5),
        trees: Some(3),
        cars: Some(2),
        perfumes: Some(1),
    };
    let possible_sues: Vec<Sue> =
        sues
        .iter()
        .filter(|sue| {
            // filter out wrong # of children
            let children = match sue.children {
                Some(children) => children == my_sue.children.unwrap(),
                None    => true,
            };
            if !children { return false; }
            // filter out wrong # of cats
            let cats = match sue.cats {
                Some(cats) => cats > my_sue.cats.unwrap(),
                None       => true,
            };
            if !cats { return false; }
            // filter out wrong # of samoyeds
            let samoyeds = match sue.samoyeds {
                Some(samoyeds) => samoyeds == my_sue.samoyeds.unwrap(),
                None       => true,
            };
            if !samoyeds { return false; }
            // filter out wrong # of pomeranians
            let pomeranians = match sue.pomeranians {
                Some(pomeranians) => pomeranians < my_sue.pomeranians.unwrap(),
                None       => true,
            };
            if !pomeranians { return false; }
            // filter out wrong # of akitas
            let akitas = match sue.akitas {
                Some(akitas) => akitas == my_sue.akitas.unwrap(),
                None       => true,
            };
            if !akitas { return false; }
            // filter out wrong # of vizslas
            let vizslas = match sue.vizslas {
                Some(vizslas) => vizslas == my_sue.vizslas.unwrap(),
                None       => true,
            };
            if !vizslas { return false; }
            // filter out wrong # of goldfish
            let goldfish = match sue.goldfish {
                Some(goldfish) => goldfish < my_sue.goldfish.unwrap(),
                None       => true,
            };
            if !goldfish { return false; }
            // filter out wrong # of trees
            let trees = match sue.trees {
                Some(trees) => trees > my_sue.trees.unwrap(),
                None       => true,
            };
            if !trees { return false; }
            // filter out wrong # of cars
            let cars = match sue.cars {
                Some(cars) => cars == my_sue.cars.unwrap(),
                None       => true,
            };
            if !cars { return false; }
            // filter out wrong # of perfumes
            let perfumes = match sue.perfumes {
                Some(perfumes) => perfumes == my_sue.perfumes.unwrap(),
                None       => true,
            };
            if !perfumes { return false; }
            true
        })
        .map(|sue| sue.clone())
        .collect();
    for sue in possible_sues {
        println!("id: {}", sue.id);
    }
}
