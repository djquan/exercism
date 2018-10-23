use std::collections::BTreeMap;

pub fn raindrops(n: u32) -> String {
    let mut sounds = Vec::new();
    let mut sound_map = BTreeMap::new();
    sound_map.insert(3, "Pling");
    sound_map.insert(5, "Plang");
    sound_map.insert(7, "Plong");

    for (drops, sound) in sound_map {
        if n % drops == 0 {
            sounds.push(sound);
        }
    }

    if sounds.len() == 0 {
        return n.to_string();
    }

    sounds.concat()
}
