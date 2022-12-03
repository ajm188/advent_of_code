use std::collections::BinaryHeap;
use std::{io::BufRead, env};
use std::iter;

use lib::io;

#[derive(Clone)]
struct Elf {
    foodstuffs: Vec<i64>
}

impl Elf {
    fn new() -> Elf {
        Elf { foodstuffs: vec![] }
    }

    fn calories(&self) -> i64 {
        self.foodstuffs.iter().map(|x| *x).sum()
    }
}

fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let mut elves: Vec<Elf> = vec![];
    let mut current_elf = Elf::new();

    for line in reader.lines() {
        match line.unwrap().as_str() {
            "" => {
                elves.push(current_elf.clone());
                current_elf = Elf::new();
            },
            l => {
                let calories: i64 = l.to_string().parse().unwrap();
                current_elf.foodstuffs.push(calories);
            }
        }
    }

    let mut sorted_calories: BinaryHeap<_> = elves.iter().map(|elf| elf.calories()).collect();

    println!("part 1: {:?}", sorted_calories.peek().unwrap());
    println!("part 2: {:?}", iter::repeat(0).take(3).map(|_| sorted_calories.pop().unwrap()).sum::<i64>());
}
