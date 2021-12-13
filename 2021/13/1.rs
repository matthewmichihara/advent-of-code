use std::collections::HashSet;

fn main() {
    let input_parts: Vec<&str> = include_str!("input.txt").split("\n\n").collect();
    let dots: HashSet<(i32, i32)> = input_parts[0].lines()
        .map(|line| {
            let split: Vec<&str> = line.split(',').collect();
            (split[0].parse::<i32>().unwrap(), split[1].parse::<i32>().unwrap())
        })
        .collect();

    let folds: Vec<(char, i32)> = input_parts[1].lines()
        .map(|line| {
            let split: Vec<&str> = line.split('=').collect();
            (split[0].chars().last().unwrap(), split[1].parse().unwrap())
        })
        .collect();

    let fold = folds[0];

    let folded: HashSet<(i32, i32)> = dots.into_iter()
        .map(|(x, y)| {
            match fold {
                ('x', pos) => {
                    if pos < x {
                        (x - 2 * (x - pos), y)
                    } else {
                        (x, y)
                    }
                },
                ('y', pos) => {
                    if pos < y {
                        (x, y - 2 * (y - pos))
                    } else {
                        (x, y)
                    }
                },
                _ => panic!("something went wrong"),
            }
        })
        .collect();

    println!("{}", folded.len());
}
