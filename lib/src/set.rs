use std::collections::HashSet;
use std::hash::Hash;

// Returns the HashSet representing the intersection of the list of given HashSets.
//
// This is logically equivalent to `sets.reduce(|acc, set| acc & set)`.
//
// Due to implementation detail constraints, this only works for HashSets
// containing types that implement Eq (in order to call .contains).
// Also, the result set contains elements of &T rather than T to avoid
// unnecessarily constraining T further to only types implementing Copy.
pub fn intersect_all<T: Eq + Hash>(sets: &[HashSet<T>]) -> HashSet<&T> {
    if sets.len() == 0 {
        return HashSet::new();
    }

    HashSet::from_iter(sets[0].iter().filter(move |item| sets[1..].iter().all(|set| set.contains(item))))
}