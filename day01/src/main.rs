use std::env::args;

fn fuel_required(mass: i32) -> i32 { (mass / 3) - 2 }

fn main() {
    let modules = args().skip(1).map(|arg| {
        arg.parse().unwrap()
    });

    let total_fuel: i32 = modules.map(|module| { fuel_required(module) }).sum();

    println!("{}", total_fuel);
}
