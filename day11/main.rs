use std::env::args;

struct PasswordGenerator { password: String, }
impl Iterator for PasswordGenerator {
    type Item = String;

    fn next(&mut self) -> Option<Self::Item> {
        let mut pchars: Vec<char> = self.password.chars().collect();
        pchars.reverse();
        self.password =
            pchars
            .iter()
            .fold((true, "".to_string()), |tup, c| {
                let (bump, chars) = tup;
                if bump {
                    let next = if *c == 'z' {
                        'a'
                    } else {
                        std::char::from_u32(*c as u32 + 1).unwrap()
                    };
                    (next == 'a', format!("{}", next) + &chars)
                } else {
                    (bump, format!("{}", c) + &chars)
                }
            }).1;
        Some(self.password.clone())
    }
}

fn is_valid_password(password: &String) -> bool {
    let chars: Vec<char> = password.chars().collect();
    let has_straight = (0..(chars.len() - 2)).any(|i| {
        let first = *chars.get(i).unwrap() as u32;
        let second = *chars.get(i + 1).unwrap() as u32;
        let third = *chars.get(i + 2).unwrap() as u32;
        (third as i32 - second as i32 == 1) &&
            (second as i32 - first as i32 == 1)
    });
    let has_repeat = (0..(chars.len() - 3)).any(|i| {
        chars.get(i).unwrap() == chars.get(i + 1).unwrap() &&
        ((i + 2)..(chars.len() - 1)).any(|j| {
            chars.get(j).unwrap() == chars.get(j + 1).unwrap()
        })
    });
    let has_bad_letters = chars.iter().any(|c| {
        match *c {
            'i' | 'o' | 'l' => true,
            _               => false,
        }
    });
    has_straight && has_repeat && !has_bad_letters
}

fn main() {
    let password = args().nth(1).unwrap();
    let mut pg = PasswordGenerator { password: password.clone(), };
    let next_pw = pg.find(|password| is_valid_password(password));
    println!("{}", next_pw.unwrap());
    let next_next_pw = pg.find(|password| is_valid_password(password));
    println!("{}", next_next_pw.unwrap());
}
