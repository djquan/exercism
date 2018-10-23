static SOUND_MAP: [(u32, &str);3] = [(3, "Pling"), (5, "Plang"), (7, "Plong")];

pub fn raindrops(n: u32) -> String {
    let mut sounds = Vec::new();
    for (drops, sound) in SOUND_MAP.iter() {
        if n % drops == 0 {
            sounds.push(*sound);
        }
    }

    if sounds.len() == 0 {
        return n.to_string();
    }

    sounds.concat()
}
