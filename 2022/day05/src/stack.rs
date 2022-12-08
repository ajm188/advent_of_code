pub struct Stack<T> {
    s: Vec<T>,
}

impl<T: Copy> Stack<T> {
    pub fn from(v: Vec<T>) -> Self {
        Stack { s: v }
    }

    pub fn peek(self) -> Option<T> {
        match self.s.len() {
            0 => None,
            _ => Some(self.s[self.s.len()-1]),
        }
    }

    pub fn pop(&mut self) -> Option<T> {
        self.s.pop()
    }

    pub fn popn(&mut self, n: usize) -> Vec<T> {
        let mut popped = vec![];
        while popped.len() < n {
            match self.pop() {
                None => break,
                Some(t) => popped.push(t),
            };
        }

        popped
    }

    pub fn push(&mut self, t: T) {
        self.s.push(t);
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn popn() {
        let st = Stack::from(vec![1, 2, 3]);
        assert_eq!(st.peek(), Some(3));

        let st = Stack::from(vec![1, 2, 3]);
        assert_eq!(st.s, vec![1, 2, 3]);

        let mut st = Stack::from(vec![1, 2, 3]);
        let elements = st.popn(3);
        assert_eq!(elements, vec![3, 2, 1]);
        assert_eq!(st.peek(), None);

        let mut st = Stack::from(vec![1, 2, 3]);
        let elements = st.popn(100);
        assert_eq!(elements, vec![3, 2, 1]);
    }
}
