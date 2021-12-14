use std::collections::HashMap;

fn main() {
    let input_parts: Vec<&str> = include_str!("input.txt").split("\n\n").collect();
    let template: Vec<char> = input_parts[0].chars().collect();
    let rules: HashMap<(char, char), char> = input_parts[1]
        .lines()
        .map(|line| {
            let rule: Vec<&str> = line.split(" -> ").collect();
            let pair: Vec<char> = rule[0].chars().collect();
            let element: Vec<char> = rule[1].chars().collect();
            ((pair[0], pair[1]), element[0])
        })
        .collect();

    let mut polymer = template.clone();
    for _ in 0..10 {
        let mut next_polymer: Vec<char> = vec![polymer[0]];
        for i in 0..polymer.len() - 1 {
            let pair = (polymer[i], polymer[i + 1]);
            let element = rules.get(&pair).unwrap();
            next_polymer.push(*element);
            next_polymer.push(pair.1);
        }
        polymer = next_polymer;
    }

    let mut hist: HashMap<char, u64> = HashMap::new();
    for c in &polymer {
        *hist.entry(*c).or_insert(0) += 1;
    }

    let max = hist.values().max().unwrap();
    let min = hist.values().min().unwrap();

    println!("{}", max - min);
}
