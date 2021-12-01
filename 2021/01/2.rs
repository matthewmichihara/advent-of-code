use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let depths: Vec<i32> = reader.lines()
        .filter_map(Result::ok)
        .map(|line| line.parse::<i32>())
        .filter_map(Result::ok)
        .collect();

    let mut last_window = -1;
    let mut increases = 0;

    for i in 0..depths.len() - 3 {
        let window = depths[i] + depths[i+1] + depths[i+2];
        if last_window == window {
            continue;
        }

        if window > last_window {
            increases += 1;
        }

        last_window = window;
    }

    println!("Increases: {}", increases);
    Ok(())
}
