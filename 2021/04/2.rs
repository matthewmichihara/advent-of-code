fn main() {
  let input = include_str!("input.txt");
  let parts: Vec<&str> = input.split("\n\n").collect();
  let nums: Vec<i32> = parts[0].split(",")
    .map(|n| n.parse().expect("Could not parse number"))
    .collect();

  let cards: Vec<Vec<Vec<i32>>> = parts[1..].iter().map(|part| {
    part.lines().map(|line| {
      line.split_whitespace()
        .map(|n| n.parse().expect("Could not parse number"))
        .collect()
    }).collect()
  }).collect();

  let mut marks: Vec<Vec<Vec<bool>>> = cards.iter().map(|card| {
    card.iter().map(|row| {
      row.iter().map(|_| false).collect()
    }).collect()
  }).collect();

  println!("{}", final_score(&nums, &cards, &mut marks));
}

fn final_score(nums: &Vec<i32>, cards: &Vec<Vec<Vec<i32>>>, marks: &mut Vec<Vec<Vec<bool>>>) -> i32 {
  let mut winning_card_indices = vec![false; cards.len()];

  for num in nums {
    for card_index in 0..cards.len() {
      if winning_card_indices[card_index] {
        continue;
      }
      let card = &cards[card_index];
      for r in 0..card.len() {
        for c in 0..card[r].len() {
          if card[r][c] == *num {
            marks[card_index][r][c] = true;
          }

          if is_win(&marks[card_index]) {
            winning_card_indices[card_index] = true;
            if winning_card_indices.iter().all(|&b| b) {
              return unmarked_num_sum(&card.to_vec(), &marks[card_index]) * num; 
            }
          }
        }
      }
    }
  }

  panic!("Should not get here");
}

fn is_win(card: &Vec<Vec<bool>>) -> bool {
  for row in card {
    if row.iter().all(|&n| n) {
      return true;
    }
  }

  for c in 0..card[0].len() {
    let mut all_true = true; 
    for row in card {
      if !row[c] {
        all_true = false;
      }
    }
    if all_true {
      return true;
    }
  }

  return false;
}

fn unmarked_num_sum(card: &Vec<Vec<i32>>, marks: &Vec<Vec<bool>>) -> i32 {
  let mut sum = 0;
  for r in 0..card.len() {
    for c in 0..card[r].len() {
      let num = card[r][c];
      if !marks[r][c] {
        sum += num;
      }
    }
  }
  return sum;
}
