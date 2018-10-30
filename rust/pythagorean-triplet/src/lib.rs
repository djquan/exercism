pub fn find() -> Option<u32> {
    (1u32..1000)
        .flat_map(|a| (a..1000 - a).map(move |b| (a, b, 1000 - a - b)))
        .filter(|(a, b, c)| a.pow(2) + b.pow(2) == c.pow(2))
        .map(|(a, b, c)| a * b * c)
        .nth(0)
}
