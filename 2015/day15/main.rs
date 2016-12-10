#![feature(iter_arith)]
trait Ingredient {
    fn capacity() -> i64;
    fn durability() -> i64;
    fn flavor() -> i64;
    fn texture() -> i64;
    fn calories() -> i64;
}

struct Sprinkles;
struct PeanutButter;
struct Frosting;
struct Sugar;

// just gonna hard code this stuff
impl Ingredient for Sprinkles {
    fn capacity() -> i64 { 5 }
    fn durability() -> i64 { -1 }
    fn flavor() -> i64 { 0 }
    fn texture() -> i64 { 0 }
    fn calories() -> i64 { 5 }
}
impl Ingredient for PeanutButter {
    fn capacity() -> i64 { -1 }
    fn durability() -> i64 { 3 }
    fn flavor() -> i64 { 0 }
    fn texture() -> i64 { 0 }
    fn calories() -> i64 { 1 }
}
impl Ingredient for Frosting {
    fn capacity() -> i64 { 0 }
    fn durability() -> i64 { -1 }
    fn flavor() -> i64 { 4 }
    fn texture() -> i64 { 0 }
    fn calories() -> i64 { 6 }
}
impl Ingredient for Sugar {
    fn capacity() -> i64 { -1 }
    fn durability() -> i64 { 0 }
    fn flavor() -> i64 { 0 }
    fn texture() -> i64 { 2 }
    fn calories() -> i64 { 8 }
}

fn capacity(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    Sprinkles::capacity() * sprinkles +
        PeanutButter::capacity() * peanut_butter +
        Frosting::capacity() * frosting +
        Sugar::capacity() * sugar
}
fn durability(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    Sprinkles::durability() * sprinkles +
        PeanutButter::durability() * peanut_butter +
        Frosting::durability() * frosting +
        Sugar::durability() * sugar
}
fn flavor(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    Sprinkles::flavor() * sprinkles +
        PeanutButter::flavor() * peanut_butter +
        Frosting::flavor() * frosting +
        Sugar::flavor() * sugar
}
fn texture(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    Sprinkles::texture() * sprinkles +
        PeanutButter::texture() * peanut_butter +
        Frosting::texture() * frosting +
        Sugar::texture() * sugar
}
fn calories(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    Sprinkles::calories() * sprinkles +
        PeanutButter::calories() * peanut_butter +
        Frosting::calories() * frosting +
        Sugar::calories() * sugar
}

fn score(sprinkles: i64, peanut_butter: i64, frosting: i64, sugar: i64) -> i64 {
    let amounts = vec![
        capacity(sprinkles, peanut_butter, frosting, sugar),
        durability(sprinkles, peanut_butter, frosting, sugar),
        flavor(sprinkles, peanut_butter, frosting, sugar),
        texture(sprinkles, peanut_butter, frosting, sugar),
    ];
    if amounts.iter().any(|v| *v < 0) {
        0
    } else {
        amounts.iter().product()
    }
}

fn main() {
    let mut bs = None;
    let mut best = (25, 25, 25, 25);
    for sprinkles in (0..101) {
        for pb in (0..101) {
            for frosting in (0..101) {
                for sugar in (0..101) {
                    if sprinkles + pb + frosting + sugar != 100 {
                        continue;
                    }
                    if calories(sprinkles, pb, frosting, sugar) != 500 {
                        continue;
                    }
                    let this_score = score(sprinkles, pb, frosting, sugar);
                    if bs.is_none() || bs.unwrap() < this_score {
                        bs = Some(this_score);
                        best = (sprinkles, pb, frosting, sugar);
                    }
                }
            }
        }
    }
    println!("({}, {}, {}, {}): {}",
        best.0, best.1, best.2, best.3, bs.unwrap());
}
