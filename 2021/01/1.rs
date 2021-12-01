use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut last_depth = -1;
    let mut increases = 0;
    for line in reader.lines() {
        let depth = line?.parse::<i32>().unwrap();
        if last_depth == -1 {
            last_depth = depth;
            continue;
        }

        if depth > last_depth{
            increases += 1;
        }

        last_depth = depth;
    }

    println!("Increases: {}", increases);

    Ok(())
}
