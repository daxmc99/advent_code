use std::fs;
use std::io::{BufRead, BufReader};

fn main() {
    let filename = "day1_input.txt";

    println!("Reading {}", filename);

    let mut fuel_total = 0;

    let contents = fs::File::open(filename).expect("Something went wrong reading the file");
    let buffered = BufReader::new(contents);

    for line in buffered.lines() {
        let line = line.expect("Unable read file");
        let module_mass = line.parse::<i32>().unwrap();
        fuel_total += fuel_calc(module_mass);
        //println!("{} ", line);
    }
    println!("Total fuel is: {}", fuel_total)
}

fn fuel_calc(mass: i32) -> i32 {
    let mut x = mass / 3; // integer division rounds down
    x = x - 2;
    return x;
}
