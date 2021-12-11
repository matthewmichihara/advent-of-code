use std::collections::BTreeSet;
use std::collections::HashMap;

fn main() {
    let input = include_str!("input.txt");
    let lines: Vec<(Vec<BTreeSet<char>>, Vec<BTreeSet<char>>)> = input
        .lines()
        .map(|line| {
            let parts: Vec<&str> = line.split(" | ").collect();
            let signals: Vec<BTreeSet<char>> = parts[0]
                .split(" ")
                .map(|signal| {
                    signal.chars().collect::<BTreeSet<char>>()
                })
                .collect();

            let outputs: Vec<BTreeSet<char>> = parts[1]
                .split(" ")
                .map(|output| {
                    output.chars().collect::<BTreeSet<char>>()
                })
                .collect();
            (signals, outputs)
        })
        .collect();

    let count: u32 = lines.iter()
        .map(|(signals, outputs)| {
            let mut digits: HashMap<&BTreeSet<char>, u8> = HashMap::new();

            let one: &BTreeSet<char> = signals.iter().find(|signal| signal.len() == 2).unwrap();
            digits.insert(one, 1);

            let seven: &BTreeSet<char> = signals.iter().find(|signal| signal.len() == 3).unwrap();
            digits.insert(seven, 7);

            let four: &BTreeSet<char> = signals.iter().find(|signal| signal.len() == 4).unwrap();
            digits.insert(four, 4);

            let eight: &BTreeSet<char> = signals.iter().find(|signal| signal.len() == 7).unwrap();
            digits.insert(eight, 8);

            let nine: &BTreeSet<char> = signals.iter()
                .find(|signal| signal.len() == 6 && signal.is_superset(four)).unwrap();
            digits.insert(nine, 9);
            
            let zero: &BTreeSet<char> = signals.iter()
                .find(|signal| {
                    signal.len() == 6 && 
                        signal.is_superset(one) && 
                        signal.symmetric_difference(nine).count() != 0
                })
                .unwrap();
            digits.insert(zero, 0);

            let six: &BTreeSet<char> = signals.iter()
                .find(|signal| {
                    signal.len() == 6 && 
                        signal.symmetric_difference(nine).count() != 0 && 
                        signal.symmetric_difference(zero).count() != 0
                })
                .unwrap();
            digits.insert(six, 6);

            let three: &BTreeSet<char> = signals.iter()
                .find(|signal| signal.len() == 5 && signal.is_superset(seven)).unwrap();
            digits.insert(three, 3);

            let five: &BTreeSet<char> = signals.iter()
                .find(|signal| signal.len() == 5 && signal.is_subset(six)).unwrap();
            digits.insert(five, 5);

            let two: &BTreeSet<char> = signals.iter()
                .find(|signal| {
                    signal.len() == 5 && 
                        signal.symmetric_difference(three).count() != 0 &&
                        signal.symmetric_difference(five).count() != 0
                })
                .unwrap();
            digits.insert(two, 2);

            let mut o: u32 = 0;
            for output in outputs {
                o = (o * 10) + *digits.get(output).unwrap() as u32;
            }
            o
        })
        .sum::<u32>();

    println!("{}", count);
}
