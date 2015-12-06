#![feature(iter_arith)]
use std::env::args;
use std::vec::Vec;

fn str_to_int(s: String) -> i32 {
    i32::from_str_radix(&s, 10).ok().unwrap()
}

struct Present {
    length: i32,
    width: i32,
    height: i32,
}

impl Present {
    fn from_string_dimensions(dimensions: String) -> Present {
        let mut dims = Vec::with_capacity(3);
        let mut s = "".to_string();
        for c in dimensions.chars() {
            if c == 'x' {
                dims.push(s.clone());
                s.clear();
            } else {
                s.push(c);
            }
        }
        dims.push(s);
        let height = str_to_int(dims[2].clone());
        let width = str_to_int(dims[1].clone());
        let length = str_to_int(dims[0].clone());
        Present {
            length: length,
            width: width,
            height: height,
        }
    }

    fn required_wrapping_paper(&self) -> i32 {
        // I really like that I get to call a builtin "unwrap" method inside a
        // Present here.
        self.surface_area() + self.side_surface_areas().iter().min().unwrap()
    }

    fn surface_area(&self) -> i32 {
        self.side_surface_areas().iter().map(|s| s * 2).sum()
    }

    fn side_surface_areas(&self) -> Vec<i32> {
        vec![self.length * self.width,
             self.length * self.height,
             self.width * self.height]
    }
}

fn main() {
    let presents: Vec<Present> = args().skip(1).map(Present::from_string_dimensions).collect();
    let amount: i32 = presents.iter().map(|p| p.required_wrapping_paper()).sum();
    println!("{}", amount);
}
