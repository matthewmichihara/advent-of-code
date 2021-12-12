use std::collections::HashMap;
use std::collections::HashSet;

fn main() {
    let mut caves: HashMap<&str, Vec<&str>> = HashMap::new();
    for line in include_str!("input.txt").lines() {
        let split: Vec<&str> = line.split('-').collect();
        let (a, b) = (split[0], split[1]);

        let tos = caves.entry(a).or_insert(Vec::new());
        tos.push(b);
        let tos = caves.entry(b).or_insert(Vec::new());
        tos.push(a);
    }

    let visited: HashMap<&str, u8> = HashMap::new();
    let starting_path = vec!["start"];
    let paths = find_paths(&caves, visited, starting_path);
    let mut sorted: Vec<&Vec<&str>> = paths.iter().collect();
    sorted.sort();
    for path in sorted {
        println!("{:?}", path);
    }
    println!("{}", paths.len());
}

fn find_paths<'a>(
    caves: &'a HashMap<&str, Vec<&str>>, 
    visited: HashMap<&'a str, u8>,
    path: Vec<&'a str>
) -> HashSet<Vec<&'a str>> {
    let cave = path.last().unwrap();

    if *cave == "end" {
        return HashSet::from([path]);
    }

    // can't go back to start
    if path.len() > 1 && *cave == "start" {
        return HashSet::new();
    }

    if !is_big(cave) && visited.contains_key(cave) && *visited.get(cave).unwrap() >= 3 {
        return HashSet::new();
    }

    if visited.iter().map(|(_,v)| v).filter(|v| *v >= &2).count() > 1 {
        return HashSet::new();
    }


    let neighbors = neighbors(caves, cave);
    let mut paths = HashSet::new();

    for neighbor in neighbors {
        let mut path_clone = path.clone();
        path_clone.push(neighbor);

        let mut new_visited = visited.clone();
        if !is_big(neighbor) {
            *new_visited.entry(neighbor).or_insert(0) += 1;
        }

        let new_paths = find_paths(caves, new_visited.clone(), path_clone);
        paths.extend(new_paths);
    }

    paths
}

fn neighbors<'a>(caves: &'a HashMap<&str, Vec<&str>>, cave: &str) -> Vec<&'a str> {
    let mut r: Vec<&str> = Vec::new();
    for n in caves.get(cave).unwrap() {
        r.push(n);
    }
    r
}

fn is_big(s: &str) -> bool {
    s.chars().all(|c| c.is_uppercase())
}
