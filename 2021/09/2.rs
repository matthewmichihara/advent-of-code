use std::collections::BTreeSet;
use std::collections::VecDeque;

fn main() {
    let input = include_str!("input.txt");
    let map: Vec<Vec<u32>> = input
        .lines()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap()).collect())
        .collect();
    let mut basins: BTreeSet<BTreeSet<Pos>> = BTreeSet::new();
    let mut visited: BTreeSet<Pos> = BTreeSet::new();

    for r in 0..map.len() {
        let row = &map[r];
        for c in 0..row.len() {
            let p = Pos { r: r, c: c };

            // Get basin
            if let Some(basin) = get_basin(&map, &mut visited, p) {
                basins.insert(basin);
            }
        }
    }

    let mut lens: Vec<usize> = basins.iter().map(|basin| basin.len()).collect();
    lens.sort();
    lens.reverse();
    let prod: usize = lens.iter().take(3).product();
    println!("{}", prod);
}

// lol
#[derive(PartialEq, Eq, PartialOrd, Ord, Copy, Clone, Debug)]
struct Pos {
    r: usize,
    c: usize,
}

fn get_basin(map: &Vec<Vec<u32>>, visited: &mut BTreeSet<Pos>, pos: Pos) -> Option<BTreeSet<Pos>> {
    if visited.contains(&pos) {
        return None;
    }

    if map[pos.r][pos.c] == 9 {
        return None;
    }

    let mut basin: BTreeSet<Pos> = BTreeSet::new();
    let mut q: VecDeque<Pos> = VecDeque::new();
    q.push_back(pos);

    loop {
        match q.pop_front() {
            Some(p) => {
                if visited.contains(&p) {
                    continue;
                }

                visited.insert(p);
                basin.insert(p);

                // left
                if p.c > 0 {
                    if map[p.r][p.c - 1] != 9 {
                        q.push_back(Pos { r: p.r, c: p.c - 1 });
                    }
                }

                // right
                if p.c < map[0].len() - 1 {
                    if map[p.r][p.c + 1] != 9 {
                        q.push_back(Pos { r: p.r, c: p.c + 1 });
                    }
                }

                // top
                if p.r > 0 {
                    if map[p.r - 1][p.c] != 9 {
                        q.push_back(Pos { r: p.r - 1, c: p.c });
                    }
                }

                // bottom
                if p.r < map.len() - 1 {
                    if map[p.r + 1][p.c] != 9 {
                        q.push_back(Pos { r: p.r + 1, c: p.c });
                    }
                }
            }
            None => break,
        }
    }

    return Some(basin);
}
