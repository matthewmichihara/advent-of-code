fn main() {
    let reader = include_str!("input.txt");

    let mut pos = 0;
    let mut depth = 0;
    for line in reader.lines() {
        let split: Vec<&str> = line.split_whitespace().collect();
        let dir = split[0];
        let units: u32 = split[1].parse().expect("Could not parse int");

        match dir {
            "forward" => pos += units,
            "down" => depth += units,
            "up" => depth -= units,
            _ => panic!("Unhandled dir")
        }
    }

    println!("{}", pos * depth);
}
