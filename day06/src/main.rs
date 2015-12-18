#![feature(iter_arith)]
#![feature(convert)]
extern crate regex;

use std::io;
use std::io::prelude::*; // I still don't understand why I need this
use std::ops::Range;
use regex::Regex;

trait Contains<T> {
    fn contains(&self, T) -> bool;
}

impl<T: PartialOrd> Contains<T> for Range<T> {
    fn contains(&self, e: T) -> bool {
        self.start <= e && e < self.end
    }
}

fn count_where<I,F>(iterator: I, f: F) -> u32
where I: Iterator, F: Fn(&I::Item) -> bool {
    iterator.filter(f).count() as u32
}

struct HouseLights { matrix: Vec<Vec<u32>>, }
struct Point<T> { x: T, y: T }

impl HouseLights {
    fn init(dim: u32) -> HouseLights {
        let mat: Vec<Vec<u32>> = 
            (0..dim).map(|_| (0..dim).map(|_| 0).collect()).collect();
        HouseLights { matrix: mat, }
    }

    fn total_brightness(&self) -> u32 {
        // sadly, rust can't seem to infer that adding a bunch of nested u32's
        // will result in a u32
        // self.matrix.iter().map(|row| row.iter().sum()).sum()
        let row_sums: Vec<u32> = self.matrix.iter().map(|row| row.iter().sum()).collect();
        row_sums.iter().sum()
    }

    /*
    fn count_lights(&self, state: bool) -> u32 {
        self.matrix.iter()
            .map(|row| count_where(row.iter(), |l| state == **l))
            .sum()
    }
    */

    fn switch<F>(&self, p1: Point<usize>, p2: Point<usize>,
                 func: F) -> HouseLights where F: Fn(u32) -> u32 {
        let dim = self.matrix.len();
        let x_range = p1.x..(p2.x + 1);
        let y_range = p1.y..(p2.y + 1);

        let mat = self.matrix.iter().zip(0..dim)
            .map(|t_y| {
                let (row, y) = t_y;
                row.iter().zip(0..dim).map(|t_x| {
                    let (state, x) = t_x;
                    if x_range.contains(x) && y_range.contains(y) {
                        func(*state)
                    } else {
                        *state
                    }
                }).collect()
            }).collect();
        HouseLights{ matrix: mat, }
    }
}

fn unpack(s: String) -> (Point<usize>, Point<usize>, String) {
    let re = Regex::new(r"(\w+(?: \w+)?) (\d+),(\d+) .* (\d+),(\d+)").unwrap();
    let (action, x1, y1, x2, y2) = match re.captures(&s) {
        Some(cap) => (cap.at(1).unwrap(), cap.at(2).unwrap(),
                      cap.at(3).unwrap(), cap.at(4).unwrap(),
                      cap.at(5).unwrap()),
        None      => panic!("could not parse input line"),
    };
    let p1 = Point {
        x: usize::from_str_radix(&x1, 10).unwrap(),
        y: usize::from_str_radix(&y1, 10).unwrap(),
    };
    let p2 = Point {
        x: usize::from_str_radix(&x2, 10).unwrap(),
        y: usize::from_str_radix(&y2, 10).unwrap(),
    };
    (p1, p2, action.to_string())
}

fn main() {
    let stdin = io::stdin();

    let h_1 = HouseLights::init(1000);
    let lines = stdin.lock().lines().map(|line| line.unwrap());
    let h_n = lines.map(unpack)
        .fold(h_1, |h, t| match t.2.as_str() {
            "turn on" => h.switch(t.0, t.1, |b: u32| -> u32 { b + 1 }),
            "turn off" => h.switch(t.0, t.1, |b: u32| -> u32 { if b < 1 { 0 } else { b - 1 } }),
            "toggle" => h.switch(t.0, t.1, |b: u32| -> u32 { b + 2 }),
            _ => panic!("could not parse input line"),
        });
    //println!("{}", h_n.count_lights(true));
    println!("{}", h_n.total_brightness());
}
