use std::cmp;

fn main() {
    let input = include_str!("input.txt");
    let crabs: Vec<i32> = input.trim().split(",").map(|n| n.parse().unwrap()).collect();
    
    let mut min_fuel = i32::MAX;
    for crab in &crabs {
        min_fuel = cmp::min(min_fuel, fuel(&crabs, &crab));            
    }

    println!("{}", min_fuel);
}

fn fuel(crabs: &Vec<i32>, pos: &i32) -> i32 {
    crabs.iter().fold(0, |accum, c| {
        let diff = (c - pos).abs();
        let sum = (diff * (diff + 1)) / 2;
        accum + sum
    })
}
