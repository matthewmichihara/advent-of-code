use std::iter::FromIterator;

fn main() {
    println!("{}", oxygen_generator_rating() * co2_scrubber_rating());
}

fn oxygen_generator_rating() -> isize {
    let reader = include_str!("input.txt");

    // Read everything all diagnostics into vector.
    let mut diagnostics: Vec<Vec<char>> = reader.lines()
        .map(|diagnostic| diagnostic.chars().collect())
        .collect();

    let diagnostic_length = diagnostics[0].len();

    for i in 0..diagnostic_length {
        let mut ones = 0;
        for diagnostic in &diagnostics {
           if diagnostic[i] == '1' {
                ones += 1;
           }
        }

        let zeroes = diagnostics.len() - ones;
        let most_common_value = if ones >= zeroes { '1' } else { '0' };
        diagnostics.retain(|diagnostic| diagnostic[i] == most_common_value);

        if diagnostics.len() == 1 {
            let rating = String::from_iter(&diagnostics[0]);
            return isize::from_str_radix(&rating, 2).expect("Could not convert to decimal");
        }
    }

    panic!("Invalid state");
}

fn co2_scrubber_rating() -> isize {
    let reader = include_str!("input.txt");

    // Read everything all diagnostics into vector.
    let mut diagnostics: Vec<Vec<char>> = reader.lines()
        .map(|diagnostic| diagnostic.chars().collect())
        .collect();

    let diagnostic_length = diagnostics[0].len();

    for i in 0..diagnostic_length {
        let mut zeroes = 0;
        for diagnostic in &diagnostics {
           if diagnostic[i] == '0' {
                zeroes += 1;
           }
        }

        let ones = diagnostics.len() - zeroes;
        let least_common_value = if zeroes <= ones { '0' } else { '1' };
        diagnostics.retain(|diagnostic| diagnostic[i] == least_common_value);

        if diagnostics.len() == 1 {
            let rating = String::from_iter(&diagnostics[0]);
            return isize::from_str_radix(&rating, 2).expect("Could not convert to decimal");
        }
    }

    panic!("Invalid state");
}
