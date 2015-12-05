#![feature(iter_arith)]
use std::env::args;
use std::vec::Vec;

fn str_to_int(s: String) -> i32 {
    i32::from_str_radix(&s, 10).ok().unwrap()
}

struct Rectangle {
    length: i32,
    width: i32,
    height: i32,
}

impl Rectangle {
    fn from_string_dims(dims: Vec<String>) -> Rectangle {
        let height = str_to_int(dims[2].clone());
        let width = str_to_int(dims[1].clone());
        let length = str_to_int(dims[0].clone());
        Rectangle {
            length: length,
            width: width,
            height: height,
        }
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

fn required_wrapping_paper(dimensions: String) -> i32 {
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
    if !s.is_empty() {
        dims.push(s.clone());
    }
    let rect = Rectangle::from_string_dims(dims);
    rect.surface_area() + rect.side_surface_areas().iter().min().unwrap()
}

fn main() {
    let amount: i32 = args().skip(1).map(required_wrapping_paper).sum();
    println!("{}", amount);
}
