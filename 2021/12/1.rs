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

    println!("{:?}", caves);

    let visited: HashSet<&str> = HashSet::new();
    let starting_path = vec!["start"];
    let paths = find_paths(&caves, visited, starting_path);
    for path in &paths {
        println!("{:?}", path);
    }
    println!("{}", paths.len());
}

fn find_paths<'a>(
    caves: &'a HashMap<&str, Vec<&str>>, 
    visited: HashSet<&'a str>,
    path: Vec<&'a str>
) -> HashSet<Vec<&'a str>> {
    let cave = path.last().unwrap();

    if *cave == "end" {
        return HashSet::from([path]);
    }

    if !is_big(cave) && visited.contains(cave) {
        return HashSet::new();
    }

    let mut new_visited = visited.clone();
    new_visited.insert(cave);

    let neighbors = neighbors(caves, cave);
    let mut paths = HashSet::new();

    for neighbor in neighbors {
        let mut path_clone = path.clone();
        path_clone.push(neighbor);
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
