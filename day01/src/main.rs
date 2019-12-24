use std::env::args;

fn fuel_required(mass: i32) -> i32 { (mass / 3) - 2 }

fn full_fuel_required(mass: i32) -> i32 {
    let mut fuel = fuel_required(mass);
    let mut new_fuel = fuel_required(fuel);
    while new_fuel > 0 {
        fuel += new_fuel;
        new_fuel = fuel_required(new_fuel);
    }

    fuel
}

fn main() {
    let modules: Vec<i32> = args().skip(1).map(|arg| {
        arg.parse().unwrap()
    }).collect();

    let total_fuel: i32 = modules.iter().map(|module| { fuel_required(*module) }).sum();
    let full_total_fuel: i32 = modules.iter().map(|module| { full_fuel_required(*module) }).sum();

    println!("{}\n{}", total_fuel, full_total_fuel);
}
