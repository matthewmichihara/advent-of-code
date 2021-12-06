fn main() {
    let input = include_str!("input.txt");
    let mut fishes: Vec<i32> = input.trim().split(",").map(|n| n.parse().unwrap()).collect();

    let mut day = 1;
    while day <= 80 {
        let mut new_fishes = 0;
        for fish in fishes.iter_mut() {
            if *fish == 0 {
                *fish = 6;
                new_fishes += 1;
            } else {
                *fish -= 1;
            }
        }
        for _ in 0..new_fishes {
            fishes.push(8);
        }

        day += 1;
    }

    println!("{}", fishes.len());
}
