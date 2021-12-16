fn main() {
    let packet = include_str!("input.txt").trim();
    println!("hex: {}", packet);

    let binary = packet
        .chars()
        .map(|c| hex_to_bin(c))
        .collect::<Vec<&str>>()
        .join("");
    println!("bin: {}", binary);

    let (res, _) = read_packet(&binary);
    println!("res: {}", res);
}

fn read_packet(binary: &str) -> (i64, usize) {
    let mut res = 0;
    let mut i = 0;

    let version = bin_to_dec(&binary[i..i + 3]);
    i += 3;
    println!("ver: {}", version);

    let ptype = bin_to_dec(&binary[i..i + 3]);
    i += 3;
    println!("typ: {}", ptype);

    // Literal
    if ptype == 4 {
        let mut literal: Vec<&str> = Vec::new();
        loop {
            let is_last_group = &binary[i..i + 1] == "0";
            i += 1;
            literal.push(&binary[i..i + 4]);
            i += 4;
            if is_last_group {
                break;
            }
        }
        let literal: i64 = bin_to_dec(&literal.join(""));
        println!("lit: {}", literal);
        println!("i: {}", i);
        return (literal, i);
    }

    // Operator
    let length_type_id = &binary[i..i + 1];
    i += 1;
    println!("lti: {}", length_type_id);

    let mut nums: Vec<i64> = Vec::new();
    match length_type_id {
        "0" => {
            let total_length_subpackets = bin_to_dec(&binary[i..i + 15]);
            i += 15;
            println!("tls: {}", total_length_subpackets);

            let total_length_index = i + total_length_subpackets as usize;

            while i < total_length_index {
                let (res, c) = read_packet(&binary[i..]);
                i += c;
                nums.push(res);
            }
        }
        "1" => {
            let num_subpackets = bin_to_dec(&binary[i..i + 11]);
            i += 11;
            println!("num subpackets: {}", num_subpackets);

            for _ in 0..num_subpackets {
                let (res, c) = read_packet(&binary[i..]);
                i += c;
                nums.push(res);
            }
        }
        _ => panic!("unhandled length type id: {}", length_type_id),
    }

    match ptype {
        0 => { // Sum
            res = nums.iter().sum();
        },
        1 => { // Product
            res = nums.iter().product();
        },
        2 => { // Minimum
            res = *nums.iter().min().unwrap();
        },
        3 => { // Maximum
            res = *nums.iter().max().unwrap();
        },
        5 => { // Greater than
            res = if nums[0] > nums[1] { 1 } else { 0 };
        },
        6 => { // Less than
            res = if nums[0] < nums[1] { 1 } else { 0 };
        },
        7 => { // Equal
            res = if nums[0] == nums[1] { 1 } else { 0 };
        },
        _ => {
            println!("not summing...");
        }
    }

    (res, i)
}

fn bin_to_dec(bin: &str) -> i64 {
    isize::from_str_radix(bin, 2).unwrap() as i64
}

fn hex_to_bin(hex: char) -> &'static str {
    match hex {
        '0' => "0000",
        '1' => "0001",
        '2' => "0010",
        '3' => "0011",
        '4' => "0100",
        '5' => "0101",
        '6' => "0110",
        '7' => "0111",
        '8' => "1000",
        '9' => "1001",
        'A' => "1010",
        'B' => "1011",
        'C' => "1100",
        'D' => "1101",
        'E' => "1110",
        'F' => "1111",
        _ => panic!("unhandled hex char: {}", hex),
    }
}
