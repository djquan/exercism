pub fn find() -> Option<u32> {
    for a in 1u32..1000 {
        for b in a..(1000 - a) {
            let c = 1000 - a - b;
            if a.pow(2) + b.pow(2) == c.pow(2) {
                return Some(a * b * c);
            }
        }
    }
    None
}
