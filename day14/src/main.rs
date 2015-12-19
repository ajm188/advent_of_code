use std::env::args;
use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

struct Reindeer {
    speed: u32,
    cont_travel: u32,
    rest_time: u32,
    distance: u32,
    rem_travel: u32,
    rem_rest: u32,
    score: u32
}

impl Reindeer {
    fn new(speed: u32, cont_travel: u32,
           rest_time: u32) -> Reindeer {
        Reindeer {
            speed: speed,
            cont_travel: cont_travel,
            rest_time: rest_time,
            distance: 0,
            rem_travel: cont_travel,
            rem_rest: rest_time,
            score: 0,
        }
    }

    fn fly(&self) -> Reindeer {
        if self.rem_travel > 0 {
            self.dup_with(self.distance + self.speed,
                          self.rem_travel - 1,
                          self.rem_rest,
                          self.score)
        } else if self.rem_rest > 0 {
            self.dup_with(self.distance,
                          self.rem_travel,
                          self.rem_rest - 1,
                          self.score)
        } else {
            self.dup_with(self.distance + self.speed,
                          self.cont_travel - 1,
                          self.rest_time,
                          self.score)
        }
    }

    fn dup_with(&self, distance: u32, rem_travel: u32,
                rem_rest: u32, score: u32) -> Reindeer {
        Reindeer {
            speed: self.speed,
            cont_travel: self.cont_travel,
            rest_time: self.rest_time,
            distance: distance,
            rem_travel: rem_travel,
            rem_rest: rem_rest,
            score: score,
        }
    }

    fn award(&self, dist: u32) -> Reindeer {
        let score_adj = if self.distance == dist { 1 } else { 0 };
        self.dup_with(self.distance,
                      self.rem_travel,
                      self.rem_rest,
                      self.score + score_adj)
    }
}

fn parse_line(line: &String) -> Reindeer {
    let re = Regex::new(
        r"(\w*) .* fly (\d+) .* for (\d+) .*, .* for (\d+) .*"
    ).unwrap();
    let caps = match re.captures(line) {
        Some(cap) => cap,
        None      => panic!("could not parse input"),
    };

    let speed = u32::from_str_radix(caps.at(2).unwrap(), 10).unwrap();
    let cont_travel = u32::from_str_radix(caps.at(3).unwrap(), 10).unwrap();
    let rest_time = u32::from_str_radix(caps.at(4).unwrap(), 10).unwrap();

    Reindeer::new(speed, cont_travel, rest_time)
}

fn main() {
    let flight_time_arg = args().skip(1).next().unwrap_or("2503".to_string());
    let flight_time = u32::from_str_radix(&flight_time_arg, 10).unwrap();

    let stdin = io::stdin();
    let lines = stdin.lock().lines();

    let mut reindeers: Vec<Reindeer> = lines.map(|line| parse_line(&line.unwrap())).collect();
    for _ in 0..flight_time {
        reindeers = reindeers.iter().map(|rd| rd.fly()).collect();
        let max_dist = reindeers.iter().map(|rd| rd.distance).max().unwrap();
        reindeers = reindeers.iter().map(|rd| rd.award(max_dist)).collect();
    }
    let max_dist = reindeers.iter().map(|rd| rd.distance).max();
    let max_score = reindeers.iter().map(|rd| rd.score).max();

    println!("{}", max_dist.unwrap());
    println!("{}", max_score.unwrap());
}
