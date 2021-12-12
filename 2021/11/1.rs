fn main() {
    let mut grid: Vec<Vec<i32>> = include_str!("input.txt")
        .lines()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap() as i32).collect())
        .collect();

    let mut flashes = 0;
    for i in 1..=100 {
        step(&mut grid);
        loop {
            let f = process_flashes(&mut grid);
            flashes += f;
            if f == 0 { break; }
        }
        reset_flashes(&mut grid);
    }

    println!("{}", flashes);
}

fn print_grid(grid: &Vec<Vec<i32>>) -> () {
    for row in grid {
        println!("{:?}", row);
    }
    println!("");
}

fn step(grid: &mut Vec<Vec<i32>>) -> () {
    for r in 0..grid.len() {
        let row = &grid[r];
        for c in 0..row.len() {
            grid[r][c] = grid[r][c] + 1;
        }
    }
}

fn process_flashes(grid: &mut Vec<Vec<i32>>) -> u32 {
    let mut flashes = 0;

    for r in 0..grid.len() {
        let row = &grid[r];
        for c in 0..row.len() {
            let level = grid[r][c];
            if level > 9 {
                incr_neighbors(grid, r as i32, c as i32);
                flashes += 1;
                grid[r][c] = -1;
            }
        }
    }

    flashes
}

fn reset_flashes(grid: &mut Vec<Vec<i32>>) -> () {
    for r in 0..grid.len() {
        let row = &grid[r];
        for c in 0..row.len() {
            let level = grid[r][c];
            if level == -1 {
                grid[r][c] = 0;
            }
        }
    }
}

fn incr_neighbors(grid: &mut Vec<Vec<i32>>, r: i32, c: i32) -> () {
    let neighbors: Vec<(i32, i32)> = vec![
        (r-1, c-1),
        (r-1, c),
        (r-1, c+1),
        (r, c-1),
        (r, c+1),
        (r+1, c-1),
        (r+1, c),
        (r+1, c+1)
    ].into_iter()
    .filter(|neighbor| is_valid(grid, neighbor.0, neighbor.1))
    .collect();

    for neighbor in neighbors {
        let (nr, nc) = (neighbor.0 as usize, neighbor.1 as usize);
        if grid[nr][nc] != -1 {
            grid[nr][nc] = grid[nr][nc] + 1;
        }
    }
}

fn is_valid(grid: &Vec<Vec<i32>>, r: i32, c: i32) -> bool {
    if r < 0 || c < 0 {
        return false;
    }

    if (r as usize) >= grid.len() {
        return false;
    }

    let row = &grid[r as usize];
    if (c as usize) >= row.len() {
        return false;
    }

    true
}
