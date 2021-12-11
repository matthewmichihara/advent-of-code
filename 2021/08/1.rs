fn main() {
    let input = include_str!("input.txt");
    let lines: Vec<(Vec<String>, Vec<String>)> = input
        .lines()
        .map(|line| {
            let parts: Vec<&str> = line.split(" | ").collect();
            let signals: Vec<String> = parts[0]
                .split(" ")
                .map(|signal| {
                    let mut chars: Vec<char> = signal.chars().collect();
                    chars.sort();
                    chars.iter().collect::<String>()
                })
                .collect();

            let outputs: Vec<String> = parts[1]
                .split(" ")
                .map(|output| {
                    let mut chars: Vec<char> = output.chars().collect();
                    chars.sort();
                    chars.iter().collect::<String>()
                })
                .collect();
            (signals, outputs)
        })
        .collect();

    let count: usize = lines.iter()
        .map(|(_, outputs)| {
            outputs.iter()
                .filter(|output| {
                    match output.len() {
                        2 | 3 | 4 | 7 => true,
                        _ => false,
                    }
                })
                .count()
        })
        .sum::<usize>();

    println!("{}", count);
}
