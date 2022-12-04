use std::env;
use std::io::BufRead;

use regex::Regex;

use lib::io;

enum MatchResult {
    WIN,
    LOSE,
    DRAW,
}

impl MatchResult {
    fn score(self) -> i64 {
        match self {
            Self::WIN => 6,
            Self::LOSE => 0,
            Self::DRAW => 3,
        }
    }
}

#[derive(Clone, Copy)]
enum Move {
    ROCK,
    PAPER,
    SCISSORS,
}


impl Move {
    fn parse(raw: &str) -> Result<Move, String> {
        match raw {
            "A"|"X" => Ok(Self::ROCK),
            "B"|"Y" => Ok(Self::PAPER),
            "C"|"Z" => Ok(Self::SCISSORS),
            _ => Err(format!("unknown input: {:?}", raw)),
        }
    }

    fn throw(self, other: Move) -> MatchResult {
        match self {
            Self::ROCK => match other {
                Self::ROCK => MatchResult::DRAW,
                Self::PAPER => MatchResult::LOSE,
                Self::SCISSORS => MatchResult::WIN,
            },
            Self::PAPER => match other {
                Self::ROCK => MatchResult::WIN,
                Self::PAPER => MatchResult::DRAW,
                Self::SCISSORS => MatchResult::LOSE,
            },
            Self::SCISSORS => match other {
                Self::ROCK => MatchResult::LOSE,
                Self::PAPER => MatchResult::WIN,
                Self::SCISSORS => MatchResult::DRAW,
            }
        }
    }

    fn score(self) -> i64 {
        match self {
            Self::ROCK => 1,
            Self::PAPER => 2,
            Self::SCISSORS => 3,
        }
    }
}

struct Round {
    opponent: Move,
    response: Move,
}

impl Round {
    fn new(opponent: Move, response: Move) -> Round {
        Round{
            opponent: opponent,
            response: response,
        }
    }

    fn score(self) -> i64 {
        self.response.score() + self.response.throw(self.opponent).score()
    }
}

fn main() {
    let reader: Box<dyn BufRead> = io::new_reader(env::args().nth(1));
    let round_re = Regex::new("([ABCXYZ]) ([ABCXYZ])").unwrap();

    let text: String = reader.lines().map(|l| l.unwrap()).collect();
    let rounds = round_re.captures_iter(&text).map(|cap| {
        Round::new(Move::parse(&cap[1]).unwrap(), Move::parse(&cap[2]).unwrap())
    });

    let advised_score = rounds.map(|r| r.score()).sum::<i64>();
    println!("part 1: {:?}", advised_score);
}
