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

    let mut polymer: HashMap<(char, char), u64> = HashMap::new();
    for i in 0..template.len() - 1 {
        let pair = (template[i], template[i + 1]);
        *polymer.entry(pair).or_insert(0) += 1;
    }

    for _ in 1..=40 {
        for (pair, count) in polymer.clone() {
            let element = rules.get(&pair).unwrap();
            let a = (pair.0, *element);
            let b = (*element, pair.1);
            *polymer.entry(a).or_insert(0) += count;
            *polymer.entry(b).or_insert(0) += count;
            *polymer.entry(pair).or_insert(0) -= count;
            if *polymer.get(&pair).unwrap() == 0 {
                polymer.remove(&pair);
            }
        }
    }

    let mut hist: HashMap<char, u64> = HashMap::new();
    *hist.entry(template[0]).or_insert(0) += 1;
    for (pair, count) in &polymer {
        let (_, c) = pair;
        *hist.entry(*c).or_insert(0) += count;
    }

    let max = hist.values().max().unwrap();
    let min = hist.values().min().unwrap();

    println!("{}", max - min);
}
