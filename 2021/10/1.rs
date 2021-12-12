use std::collections::VecDeque;

fn main() {
    let input = include_str!("input.txt");
    let lines: Vec<Vec<char>> = input
        .lines()
        .map(|line| line.chars().collect())
        .collect();

    let mut points = 0;
    for line in &lines {
        points += get_points(line);
    }

    println!("{}", points);
}

fn get_points(line: &Vec<char>) -> i32 {
    let mut stack: VecDeque<char> = VecDeque::new();
    for c in line {
        match c {
            '(' | '[' | '{' | '<' => stack.push_back(*c),
            ')' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '(' {
                        return 3;
                    }
                }
            },
            ']' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '[' {
                        return 57;
                    }
                }
            },
            '}' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '{' {
                        return 1197;
                    }
                }
            },
            '>' => {
                if let Some(popped) = stack.pop_back() {
                    if popped != '<' {
                        return 25137;
                    }
                }
            },
            _ => panic!("something went wrong"),
        }
    }

    return 0;
}
