use std::collections::VecDeque;

fn main() {
    let input = include_str!("input.txt");
    let mut points: Vec<i64> = input
        .lines()
        .map(|line| line.chars().collect())
        .filter(|line| !is_corrupted(line))
        .map(|line| get_points(&line))
        .collect();

    points.sort();
    let mid = points.len() / 2;
    println!("{}", points[mid]);
}

fn get_points(line: &Vec<char>) -> i64 { 
    let mut stack: VecDeque<char> = VecDeque::new();
    for c in line {
        match c {
            '(' | '[' | '{' | '<' => stack.push_back(*c),
            ')' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '(' {
                        panic!("something went wrong");
                    }
                }
            },
            ']' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '[' {
                        panic!("something went wrong");
                    }
                }
            },
            '}' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '{' {
                        panic!("something went wrong");
                    }
                }
            },
            '>' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '<' {
                        panic!("something went wrong");
                    }
                }
            },
            _ => panic!("something went wrong"),
        }
    }

    let mut points = 0;
    loop {
        match stack.pop_back() {
            Some(popped) => {
                let point = match popped {
                    '(' => 1,
                    '[' => 2,
                    '{' => 3,
                    '<' => 4,
                    _ => panic!("something went wrong"),
                };
                points = (5 * points) + point;
            },
            None => return points,
        }
    }
}

fn is_corrupted(line: &Vec<char>) -> bool { 
    let mut stack: VecDeque<char> = VecDeque::new();
    for c in line {
        match c {
            '(' | '[' | '{' | '<' => stack.push_back(*c),
            ')' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '(' {
                        return true;
                    }
                }
            },
            ']' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '[' {
                        return true;
                    }
                }
            },
            '}' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '{' {
                        return true;
                    }
                }
            },
            '>' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '<' {
                        return true;
                    }
                }
            },
            _ => panic!("something went wrong"),
        }
    }

    return false;
}
