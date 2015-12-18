/*
fn seq_counts<T: Eq + PartialEq + Clone>(list: Vec<T>) -> Vec<(T, usize)> {
    let len = list.len();
    let mut vec = vec![];
    let mut items = 0;
    while items < len {
        let first = list.iter().skip(items).next().unwrap();
        let count = list.iter().skip(items).take_while(|&x| first == x).count();
        vec.push((first.clone(), count));
        items += count;
    }
    vec
}
*/

fn seq_counts(list: Vec<char>) -> Vec<(char, usize)> {
    let mut vec = vec![];
    let mut items = 0;
    while items < list.len() {
        let first = list.iter().skip(items).next().unwrap();
        let count = list.iter().skip(items).take_while(|&x| first == x).count();
        vec.push((first.clone(), count));
        items += count;
    }
    vec
}

fn concat(substrings: Vec<String>) -> String {
    substrings.iter().fold("".to_string(), |a, s| a + &s)
}

fn main() {
    let chars: Vec<char> = "3113322113".chars().collect();
    //let chars: Vec<char> = "1".chars().collect();
    let result = (0..40).fold(
        chars,
        |s, _| {
            let subseq: Vec<String> = seq_counts(s).iter().map(|tup| {
                let c: char = tup.0;
                let count: usize = tup.1;
                format!("{}{}", count, c)
            }).collect();
            let string: String = concat(subseq);
            string.chars().collect()
        }
    );
    println!("{}", result.len());
    /*
    for c in result {
        print!("{}", c);
    }
    println!("");
    */
}
