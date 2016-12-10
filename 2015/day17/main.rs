#![feature(iter_arith)]
use std::env::args;

fn powerset<T: Clone>(t: Vec<T>) -> Vec<Vec<T>> {
    pset(t, vec![vec![]])
}

fn pset<T: Clone>(t: Vec<T>, s: Vec<Vec<T>>) -> Vec<Vec<T>> {
    if t.len() == 0 {
        s
    } else {
        let head = t.first().unwrap();
        let tail = t.iter().skip(1).map(|a| a.clone()).collect();
        let mut ps: Vec<Vec<T>> = s.iter().map(|v| {
            let mut vec = v.clone();
            vec.push(head.clone());
            vec
        }).collect();
        ps.extend(s);
        pset(tail, ps)
    }
}

fn main() {
    let eggnog = 150;
    let containers =
        args()
        .skip(1)
        .map(|arg| u32::from_str_radix(&arg, 10).unwrap())
        .collect();
    let arrangements: Vec<Vec<u32>> = 
        powerset(containers)
        .iter()
        .filter(|v| {
            let total: u32 = v.iter().sum();
            total == eggnog
        })
        .map(|v| v.clone())
        .collect();
    let smallest_arrangement_size =
        arrangements.iter().map(|v| v.len()).min().unwrap_or(0);
    let num_smallest_arrangements =
        arrangements
        .iter()
        .filter(|v| v.len() == smallest_arrangement_size)
        .count();
    println!("{}", num_smallest_arrangements);
}
