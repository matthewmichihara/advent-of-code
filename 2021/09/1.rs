fn main() {
    let input = include_str!("input.txt");
    let map: Vec<Vec<u32>> = input
        .lines()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap()).collect())
        .collect();

    let mut sum = 0;
    for r in 0..map.len() {
        let row = &map[r];
        for c in 0..row.len() {
            let p = Pos { r: r, c: c };
            if is_low(&map, p) {
                sum = sum + map[r][c] + 1;
            }
        }
    }

    println!("{}", sum);
}

struct Pos {
    r: usize,
    c: usize
}

fn is_low(map: &Vec<Vec<u32>>, p: Pos) -> bool {
    let v = map[p.r][p.c];
    
    // left
    if p.c > 0 {
        if map[p.r][p.c - 1] <= v {
            return false;
        }
    }

    // right
    if p.c < map[0].len() - 1 {
        if map[p.r][p.c + 1] <= v {
            return false;
        }
    }

    // top
    if p.r > 0 {
        if map[p.r - 1][p.c] <= v {
            return false;
        }
    }

    // bottom
    if p.r < map.len() - 1 {
        if map[p.r + 1][p.c] <= v {
            return false;
        }
    }

    return true;
}
