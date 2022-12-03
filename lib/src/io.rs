use std::fs;
use std::io;
use std::io::{BufRead, BufReader};

// Returns a new BufRead implementation which either reads from a file with the
// given filename, or stdin if filename is None.
pub fn new_reader(filename: Option<String>) -> Box<dyn BufRead +'static> {
    let reader: Box<dyn BufRead> = match filename {
        None => Box::new(BufReader::new(io::stdin())),
        Some(filename) => Box::new(BufReader::new(fs::File::open(filename).unwrap())),
    };

    reader
}