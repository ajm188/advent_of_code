#![feature(iter_arith)]
use std::env::args;
use std::fs::File;
use std::io::Read;

extern crate rustc_serialize;
use rustc_serialize::json::Json;

fn is_red(tup: (&String, &Json)) -> bool {
    match tup.1.clone() {
        Json::String(s) => s == String::from("red"),
        _               => false,
    }
}

fn sum(data: &Json) -> i64 {
    match data.clone() {
        Json::I64(v)      => v,
        Json::U64(v)      => v as i64,
        Json::Array(list) => list.iter().map(|j| sum(j)).sum(),
        Json::Object(obj) =>
            if obj.iter().any(is_red) {
                0
            } else {
                obj.iter().map(|tup| sum(tup.1)).sum()
            },
        _                 => 0,
    }
}

fn main() {
    let fname = args().nth(1).unwrap_or("input.txt".to_string());
    let mut file = File::open(fname).unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    let data = Json::from_str(&contents).unwrap();
    let count = sum(&data);
    println!("{}", count);
}
