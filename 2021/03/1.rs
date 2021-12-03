fn main() {
    let reader = include_str!("input.txt");

    let mut ones = vec![0; 12];
    let mut num_lines = 0;

    for line in reader.lines() {
        for i in 0..line.len() {
            let bits: Vec<char> = line.chars().collect();
            let bit = bits[i];
            if bit == '1' {
                ones[i] += 1;
            }
        }
        num_lines += 1;
    }

    let mut gam = String::from("");
    let mut eps = String::from("");
    for i in 0..ones.len() {
        if ones[i] > num_lines/2 {
            gam.push_str("1");
            eps.push_str("0");
        } else {
            gam.push_str("0");
            eps.push_str("1");
        }
    }

    let gam = isize::from_str_radix(&gam, 2).expect("Could not convert to decimal");
    let eps = isize::from_str_radix(&eps, 2).expect("Could not convert to decimal");
    println!("{}", gam * eps);
}
