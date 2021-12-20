use std::cmp;

#[derive(Debug)]
struct State {
    x: i32,
    y: i32,
    x_vel: i32,
    y_vel: i32,
}

enum Result {
    Continue,
    InTarget,
    Impossible,
}

fn main() {
    let (x0, x1, y0, y1) = {
        let input = include_str!("input.txt").trim();
        let range = input.split(':').collect::<Vec<&str>>()[1];
        let coords = range.split(',').map(|c| c.trim()).collect::<Vec<&str>>();
        let xs = coords[0].split('=').collect::<Vec<&str>>()[1].split("..").collect::<Vec<&str>>();
        let ys = coords[1].split('=').collect::<Vec<&str>>()[1].split("..").collect::<Vec<&str>>();
        (
            xs[0].parse::<i32>().unwrap(), 
            xs[1].parse::<i32>().unwrap(),
            ys[0].parse::<i32>().unwrap(), 
            ys[1].parse::<i32>().unwrap()
        )
    };

    let mut sum = 0;

    let x_bound = cmp::max(x0.abs(), x1.abs());
    let y_bound = cmp::max(y0.abs(), y1.abs());

    for x_vel in -x_bound..=x_bound {
        for y_vel in -y_bound..1000 {
            let mut state = State {
                x: 0,
                y: 0,
                x_vel: x_vel,
                y_vel: y_vel,
            };

            loop {
                match should_continue(&state, x0, x1, y0, y1) {
                    Result::InTarget => {
                        sum += 1;
                        break;
                    },
                    Result::Impossible => {
                        break;
                    },
                    _ => {},
                }

                state = tick(&state);
            }
        }
    }

    println!("sum: {}", sum);
}

fn tick(state: &State) -> State {
    let new_x = state.x + state.x_vel;
    let new_y = state.y + state.y_vel;
    let new_x_vel = if state.x_vel > 0 {
        state.x_vel - 1
    } else if state.x_vel < 0 {
        state.x_vel + 1
    } else {
        0
    };
    let new_y_vel = state.y_vel - 1;

    State {
        x: new_x,
        y: new_y,
        x_vel: new_x_vel,
        y_vel: new_y_vel,
    }
}

fn should_continue(state: &State, x0: i32, x1: i32, y0: i32, y1: i32) -> Result {
    if state.x >= x0 && state.x <= x1 && state.y >= y0 && state.y <= y1 {
        return Result::InTarget;
    }

    if state.x_vel >= 0 && state.x > x1 {
        return Result::Impossible;
    }

    if state.x_vel <= 0 && state.x < x0 {
        return Result::Impossible;
    }

    if state.y_vel <= 0 && state.y < y0 {
        return Result::Impossible;
    }

    Result::Continue
}
