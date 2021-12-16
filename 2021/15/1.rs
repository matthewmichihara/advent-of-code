use std::collections::BinaryHeap;
use std::collections::HashSet;
use std::cmp::Ordering;

#[derive(Debug, Eq, PartialEq)]
struct Node {
    risk: i32,
    pos: (i32, i32),
}

impl Ord for Node {
    fn cmp(&self, other: &Self) -> Ordering {
        other.risk.cmp(&self.risk)
    }
}

impl PartialOrd for Node {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

fn main() {
    let cavern: Vec<Vec<i32>> = include_str!("input.txt")
        .lines()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap() as i32).collect())
        .collect();

    let start = (0, 0);
    let end = ((cavern[0].len() - 1) as i32, (cavern.len() - 1) as i32);
    
    let sum = min_sum(&cavern, &start, &end);
    println!("{}", sum);
}

fn min_sum(cavern: &Vec<Vec<i32>>, start_pos: &(i32, i32), end_pos: &(i32, i32)) -> i32 {
    let start = Node { risk: 0, pos: start_pos.clone() };

    let mut heap: BinaryHeap<Node> = BinaryHeap::new();
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    heap.push(start);

    while let Some(Node { risk, pos }) = heap.pop() {
        if pos == *end_pos {
            return risk;
        }

        if visited.contains(&pos) {
            continue;
        }

        for n in neighbors(cavern, &pos) {
            let new_risk = risk + cavern[n.1 as usize][n.0 as usize];
            heap.push(Node { risk: new_risk, pos: n });
        }

        visited.insert(pos);
    }

    0
}

fn neighbors(cavern: &Vec<Vec<i32>>, pos: &(i32, i32)) -> Vec<(i32, i32)> {
    let n = vec![
        (pos.0 - 1, pos.1),
        (pos.0, pos.1 - 1),
        (pos.0 + 1, pos.1),
        (pos.0, pos.1 + 1),
    ];
    n.into_iter()
        .filter(|p| {
            p.0 >= 0 && p.0 < cavern[0].len() as i32 && p.1 >= 0 && p.1 < cavern.len() as i32
        })
        .collect()
}
