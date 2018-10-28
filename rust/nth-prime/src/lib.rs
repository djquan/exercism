struct Prime {
    curr: u32,
}

fn is_prime(n: u32) -> bool {
    for x in 2..(n / 2) + 1 {
        if n % x == 0 {
            return false;
        }
    }
    true
}

impl Iterator for Prime {
    type Item = u32;

    fn next(&mut self) -> Option<<Self as Iterator>::Item> {
        let mut target = self.curr + 1;
        loop {
            if is_prime(target) {
                self.curr = target;
                return Some(target);
            }
            target += 1;
        }
    }
}

fn primes() -> Prime {
    Prime { curr: 1 }
}

pub fn nth(n: u32) -> u32 {
    primes().nth(n as usize).unwrap()
}
