use std::collections::HashMap;

fn main() {
    let input = include_str!("input.txt");
    let mut fish_counts: HashMap<u32, u64> = HashMap::new();
    let fishes: Vec<u32> = input.trim().split(",").map(|n| n.parse().unwrap()).collect();
    for fish in &fishes {
       *fish_counts.entry(*fish).or_insert(0) += 1; 
    }

    let mut day = 0;
    while day < 256 {
        let new_fishes: u64 = *fish_counts.get(&0).unwrap_or(&0);
        let mut new_fish_counts: HashMap<u32, u64> = HashMap::new();
        
        for (key, value) in &fish_counts {
            if *key == 0 {
                *new_fish_counts.entry(6).or_insert(0) += *value;
            } else {
                *new_fish_counts.entry(key-1).or_insert(0) += *value;
            }
        }
        
        if new_fishes > 0 {
            new_fish_counts.insert(8, new_fishes);
        }

        fish_counts = new_fish_counts;
        day += 1;
    }

    let count: u64 = fish_counts.iter().map(|(_, value)| value).sum();

    println!("{}", count);
}
