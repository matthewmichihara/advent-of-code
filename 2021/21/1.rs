fn main() -> () {
  println!("hello world");

  let mut p1_score = 0;
  let mut p2_score = 0;

  let mut p1_pos = 8;
  let mut p2_pos = 9;

  let mut dice_value = 1;
  let mut dice_rolls = 0;

  loop {
    let p1_roll_1 = dice_value;
    let mut p1_roll_2 = (dice_value + 1) % 100;
    if p1_roll_2 == 0 {
      p1_roll_2 = 100;
    }
    let mut p1_roll_3 = (dice_value + 2) % 100;
    if p1_roll_3 == 0 {
      p1_roll_3 = 100;
    }
    let p1_roll = p1_roll_1 + p1_roll_2 + p1_roll_3;
    dice_rolls += 3;
    dice_value += 3;
    if dice_value == 0 {
      dice_value = 100;
    }
    p1_pos = (p1_pos + p1_roll) % 10;
    if p1_pos == 0 {
      p1_pos = 10;
    }

    p1_score += p1_pos;
    println!("p1 moves to space {} for score of {}", p1_pos, p1_score);

    if p1_score >= 1000 {
      println!("p1 win. score: {}", p1_score); 
      break;
    }
    
    let p2_roll_1 = dice_value;
    let mut p2_roll_2 = (dice_value + 1) % 100;
    if p2_roll_2 == 0 {
      p2_roll_2 = 100;
    }
    let mut p2_roll_3 = (dice_value + 2) % 100;
    if p2_roll_3 == 0 {
      p2_roll_3 = 100;
    }
    let p2_roll = p2_roll_1 + p2_roll_2 + p2_roll_3;
    dice_rolls += 3;
    dice_value += 3;
    if dice_value == 0 {
      dice_value = 100;
    }
    p2_pos = (p2_pos + p2_roll) % 10;
    if p2_pos == 0 {
      p2_pos = 10;
    }

    p2_score += p2_pos;
    println!("p2 moves to space {} for score of {}", p2_pos, p2_score);

    if p2_score >= 1000 {
      println!("p2 win. score: {}", p2_score); 
      break;
    }
  }

  println!("dice rolls: {}", dice_rolls);
}
