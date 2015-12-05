use std::env::args;

struct Santa {
    position: i32,
    first_time_in_basement: i32,
    num_movements: i32,
}

impl Santa {
    fn has_been_in_basement(&self) -> bool {
        self.first_time_in_basement >= 0
    }

    fn from_santa(santa: Santa, movement: i32) -> Santa {
        let pos = santa.position + movement;
        let movements = santa.num_movements + 1;
        let basement = if santa.has_been_in_basement() || pos >= 0 {
            santa.first_time_in_basement
        } else {
            movements
        };
        Santa {
            position: pos,
            first_time_in_basement: basement,
            num_movements: movements,
        }
    }
}

fn main() {
    let instructions = match args().nth(1) {
        Some(v) => v,
        None => "".to_string(),
    };
    let santa = Santa {
        position: 0,
        first_time_in_basement: -1,
        num_movements: 0,
    };
    let last_santa: Santa = instructions
        .chars()
        .map(|c: char| if c == '(' { 1 } else { -1 })
        .fold(santa, |s, i| Santa::from_santa(s, i));
    println!("{} {}", last_santa.position, last_santa.first_time_in_basement);
}
