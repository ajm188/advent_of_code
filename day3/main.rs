use std::collections::HashSet;
use std::env::args;

#[derive(PartialEq, Eq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn update(&self, instr: char) -> Point {
        match instr {
            '^' => Point{x: self.x, y: self.y + 1},
            'v' => Point{x: self.x, y: self.y - 1},
            '<' => Point{x: self.x - 1, y: self.y},
            '>' => Point{x: self.x + 1, y: self.y},
            _   => unreachable!()
        }
    }
}

fn main() {
    let instructions = match args().nth(1) {
        Some(v) => v,
        None    => "".to_string(),
    };
    let starting_point = Point{x: 0, y: 0};
    let mut points = vec![starting_point];
    for tuple in instructions.chars().zip(0..(instructions.len())) {
        let (instr, i) = match tuple { (a@_, b@_) => (a, b) };
        let next_point = points[i].update(instr);
        points.push(next_point);
    }

    let mut point_set = HashSet::new();
    for point in points {
        point_set.insert(point);
    }
    println!("{}", point_set.len());
}
