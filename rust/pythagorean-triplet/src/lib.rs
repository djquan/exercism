pub fn find() -> Option<u32> {
    (1u32..1000)
        .flat_map(|a| {
            (a..1000 - a)
                .map(move |b| (b, 1000 - a - b))
                .filter(move |bc| a.pow(2) + bc.0.pow(2) == bc.1.pow(2))
                .map(move |bc| a * bc.0 * bc.1)
        }).nth(0)
}
