use std::env;
use std::io::BufRead;

use lib::io;
use regex::Regex;

mod stack;

#[derive(Clone, Copy)]
struct Move {
    amt: usize,
    from: usize,
    to: usize,
}

impl Move {
    fn apply<T: Copy>(self, stacks: &mut Vec<stack::Stack<T>>) {
        let src = &mut stacks[self.from-1];

        let items = src.popn(self.amt);
        for t in items {
            stacks[self.to-1].push(t);
        }
    }
}

fn main() {
    let move_re = Regex::new("move ([0-9]+) from ([0-9]+) to ([0-9]+)").unwrap();

    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));

    let (moves, vecs): (Vec<Move>, Vec<Vec<char>>) = reader.lines().map(
        |line| line.unwrap()
    ).fold((Vec::new(), Vec::new()),
        |(mut moves, mut vecs), line| {
            if line == "" {
                // Skip empty line that separates the grid from the instructions.
                (moves, vecs)
            } else if line.starts_with(" 1"){
                // Skip the line that just numbers the stacks.
                (moves, vecs)
            } else if line.starts_with("move") {
                let cap = move_re.captures(&line).unwrap();

                let mv = Move{
                    amt: (cap[1].parse()).unwrap(),
                    from: (cap[2].parse()).unwrap(),
                    to: (cap[3].parse()).unwrap(),
                };

                moves.push(mv);
                (moves, vecs)
            } else {
                for (idx, c) in line.char_indices() {
                    if idx == 0 {
                        continue;
                    }
                    // These lines have the shape of:
                    //
                    //      [D]    
                    //  [N] [C]
                    //  [Z] [M] [P]
                    //
                    // So, we want the characters (or empty spaces) at indices
                    // 1, 5, 9, ..., and so on.
                    match (idx - 1) % 4 {
                        0 => {
                            if vecs.len() <= ((idx - 1) / 4) {
                                vecs.push(Vec::new());
                            }

                            if c == ' ' {
                                continue;
                            }

                            vecs[(idx - 1) / 4].push(c);
                        },
                        _ => (),
                    };
                }

                (moves, vecs)
            }
        }
    );

    let mut stacks: Vec<stack::Stack<char>> = Vec::new();
    for mut v in vecs {
        v.reverse();
        stacks.push(stack::Stack::from(v));
    }

    for mv in moves {
        mv.apply(&mut stacks);
    }

    let mut msg = "".to_string();
    for stack in stacks {
        msg.push(stack.peek().unwrap());
    }

    println!("part1: {:?}", msg);
}
