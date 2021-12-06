use std::collections::HashMap;

fn main() {
  let input = include_str!("input.txt");
  
  let lines: Vec<((i32, i32), (i32, i32))> = input.lines()
    .map(|line| {
      let points: Vec<&str> = line.split(" -> ").collect();
      let a: Vec<&str> = points[0].split(",").collect();
      let b: Vec<&str> = points[1].split(",").collect();
      (
        (
          a[0].parse().expect("can't parse number"),
          a[1].parse().expect("can't parse number")
        ), (
          b[0].parse().expect("can't parse number"),
          b[1].parse().expect("can't parse number")
        )
      )
    }).collect();

  // Store map of points -> hits
  // for each line
  //   populate all hits
  // count all hits where hit > 1

  let mut hits: HashMap<(i32, i32), i32> = HashMap::new();

  for line in &lines {
    if is_horizontal(line) {
      let (start, end) = if line.0.0 < line.1.0 {
        (line.0, line.1)
      } else {
        (line.1, line.0)
      };

      for x in start.0..=end.0 {
        *hits.entry((x, start.1)).or_insert(0) += 1;
      }
    } else if is_vertical(line) {
      let (start, end) = if line.0.1 < line.1.1 {
        (line.0, line.1)
      } else {
        (line.1, line.0)
      };

      for y in start.1..=end.1 {
        *hits.entry((start.0, y)).or_insert(0) += 1;
      }
    }
  } 

  let count = hits.iter().filter(|&(_, hits)| *hits > 1).count();

  println!("{}", count);
}

fn is_horizontal(line: &((i32, i32), (i32, i32))) -> bool {
  line.0.1 == line.1.1
}

fn is_vertical(line: &((i32, i32), (i32, i32))) -> bool {
  line.0.0 == line.1.0
}

fn get_points(line: &((i32, i32), (i32, i32))) -> Vec<(i32, i32)> {
  let (point1, point2) = line;
  let mut points: Vec<(i32, i32)> = Vec::new();

  let x_delta = if point1.0 < point2.0 {
    1
  } else {
    -1
  };

  let y_delta = if point1.1 < point2.1 {
    1
  } else {
    -1
  };

  if point1 == point2 {
    points.push(*point1);
    return points;
  }

  let mut point = point1;

  while point != point2 {
    points.push(*point);
    point.0 += x_delta;
    point.1 += y_delta;
  }
  
  points
}

