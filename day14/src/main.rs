use std::env::args;
use std::io;
use std::io::prelude::*;

extern crate regex;
use regex::Regex;

struct Reindeer {
    name: String,
    speed: u32, // in km / s
    cont_travel: u32,
    rest_time: u32,
}

impl Reindeer {
    fn new(name: String, speed: u32, cont_travel: u32,
           rest_time: u32) -> Reindeer {
        Reindeer {
            name: name,
            speed: speed,
            cont_travel: cont_travel,
            rest_time: rest_time,
        }
    }

    fn fly(&self, seconds: u32) -> u32 {
        let mut dist = 0;
        let mut elapsed_time = 0;
        while elapsed_time < seconds {
            let time_left = seconds - elapsed_time;
            if self.cont_travel < time_left {
                dist += self.speed * self.cont_travel;
            } else {
                dist += self.speed * time_left;
                break;
            }

            if self.rest_time < (time_left - self.cont_travel) {
                elapsed_time += self.cont_travel + self.rest_time;
            } else {
                break;
            }
        }
        dist
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

    let name = caps.at(1).unwrap();
    let speed = u32::from_str_radix(caps.at(2).unwrap(), 10).unwrap();
    let cont_travel = u32::from_str_radix(caps.at(3).unwrap(), 10).unwrap();
    let rest_time = u32::from_str_radix(caps.at(4).unwrap(), 10).unwrap();

    Reindeer::new(name.to_string(), speed, cont_travel, rest_time)
}

fn main() {
    let flight_time_arg = args().skip(1).next().unwrap_or("2503".to_string());
    let flight_time = u32::from_str_radix(&flight_time_arg, 10).unwrap();

    let stdin = io::stdin();
    let lines = stdin.lock().lines();

    let reindeers = lines.map(|line| parse_line(&line.unwrap()));
    let distances = reindeers.map(|rd| rd.fly(flight_time));
    let max_dist = distances.max();
    println!("{}", max_dist.unwrap());
}
