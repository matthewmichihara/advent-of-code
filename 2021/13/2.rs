use std::collections::HashSet;

fn main() {
    let input_parts: Vec<&str> = include_str!("input.txt").split("\n\n").collect();
    let mut dots: HashSet<(i32, i32)> = input_parts[0].lines()
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

    for fold in folds {
        dots = dots.into_iter()
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
    }

    display(&dots);
}

fn display(dots: &HashSet<(i32, i32)>) -> () {
    let max_x: i32 = *dots.iter().map(|(x, _)| x).max().unwrap();
    let max_y: i32 = *dots.iter().map(|(_, y)| y).max().unwrap();

    let mut grid: Vec<Vec<char>> = (0..=max_y).map(|_| vec!['.'; (max_x+1) as usize]).collect();
    for (x, y) in dots {
        grid[*y as usize][*x as usize] = '#';
    }

    for line in &grid {
        for c in line {
            print!("{}", c); 
        }
        println!()
    }
}
