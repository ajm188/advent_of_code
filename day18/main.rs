#![feature(iter_arith)]
use std::io;
use std::io::prelude::*;

struct House { lights: Vec<Vec<bool>>, }
impl House {
    fn step(&self) -> House {
        let lights = (0..self.lights.len()).map(|i| {
            (0..self.lights.len()).map(|j| {
                let light = *self.lights.get(i).unwrap().get(j).unwrap();
                let neighbors = self.neighbors(i, j);
                let neighbors_on = neighbors.iter().filter(|b| **b).count();
                if light {
                    if neighbors_on == 2 || neighbors_on == 3 {
                        true
                    } else {
                        false
                    }
                } else {
                    if neighbors_on == 3 { true } else { false }
                }
            }).collect()
        }).collect();
        House { lights: lights, }
    }

    fn neighbors(&self, x: usize, y: usize) -> Vec<bool> {
        let indices =
            (-1..2).fold(vec![], |mut list, i| {
                let points = (-1..2).filter(|j| {
                    let x_i = x as i32 + i;
                    let y_j = y as i32 + j;
                    (i != 0 || *j != 0) && self.in_bounds(x_i, y_j)
                }).map(|j| (i, j));
                list.extend(points);
                list
            });
        let neighbors: Vec<bool> = indices.iter().map(|p| {
            let x_i = x as i32 + p.0;
            let y_j = y as i32 + p.1;
            *self.lights.get(x_i as usize).unwrap().get(y_j as usize).unwrap()
        }).collect();
        neighbors
    }

    fn lights_on(&self) -> u32 {
        (0..self.lights.len()).map(|i| {
            let on_in_row: u32 = (0..self.lights.len()).map(|j| {
                if *self.lights.get(i).unwrap().get(j).unwrap() {
                    1
                } else {
                    0
                }
            }).sum();
            on_in_row
        }).sum()
    }

    fn in_bounds(&self, x: i32, y: i32) -> bool {
        x >= 0 && y >= 0 &&
            x < self.lights.len() as i32 && y < self.lights.len() as i32
    }
}

fn parse(line: &String) -> Vec<bool> {
    line.chars().map(|c| c == '#').collect()
}

fn main() {
    let stdin = io::stdin();
    let lights: Vec<Vec<bool>> =
        stdin
        .lock()
        .lines()
        .map(|line| parse(&line.unwrap()))
        .collect();
    let house = House { lights: lights, };
    let last = (0..100).fold(house, |h, _| h.step());
    println!("{}", last.lights_on());
}
