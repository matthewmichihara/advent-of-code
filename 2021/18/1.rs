#[derive(Debug)]
enum Node<'a> {
    Composite { 
        left: &'a Node<'a>, 
        right: &'a Node<'a>,
    },
    Literal { value: u8 },
}

fn main() {
    let a = "[1,2]";
    let b = "[[3,4],5]";
    let (na, _) = parse(a);
    let (nb, _) = parse(b);
    let nc = add(&na, &nb);
    print_node(&nc, "");

    let x = find_deep_pair(&nb, 2).unwrap();
    print_node(x, "");
}

fn parse(s: &str) -> (Node, usize) {
    let mut i = 0;
    let c = &s[0..1];
    match c {
        "[" => {
            i += 1; // starting '['
            let (left, offset) = parse(&s[i..]);
            i += offset + 1; // ','
            let (right, offset) = parse(&s[i..]);
            i += offset + 1; // ']'
            let node = Node::Composite { left: &left, right: &right };
            return (node, i);
        },
        _ => {
            let node = Node::Literal { value: c.parse::<u8>().unwrap() };
            i += 1;
            return (node, i);
        }
    }
}

fn print_node(n: &Node, prefix: &str) -> () {
    match n {
        Node::Composite { left, right } => {
            println!("{}*", prefix);
            print_node(left, &format!("{}{}", prefix, "  "));
            print_node(right, &format!("{}{}", prefix, "  "));
        }
        Node::Literal { value } => {
            println!("{}{}", prefix, value);
        }
    }
}

fn add<'a>(a: &'a Node, b: &'a Node) -> Node<'a> {
    Node::Composite {
        left: a,
        right: b
    }
}

fn find_deep_pair<'a> (n: &'a Node, depth: u8) -> Option<&'a Node<'a>> {
    let is_normal_pair = match n {
        Node::Composite { left, right } => {
            let left_is_literal = match left {
                Node::Literal { value } => true,
                _ => false
            };
            let right_is_literal = match right {
                Node::Literal { value } => true,
                _ => false
            };
            left_is_literal && right_is_literal
        },
        _ => false,
    };

    if is_normal_pair && depth <= 0 {
        return Some(n);
    }

    match n {
        Node::Composite { left, right } => {
            match find_deep_pair(left, depth - 1) {
                Some(l) => return Some(l),
                None => {}
            }
            match find_deep_pair(right, depth - 1) {
                Some(r) => return Some(r),
                None => {}
            }
        },
        _ => {}
    }

    return None
}
