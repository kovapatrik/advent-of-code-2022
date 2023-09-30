fn main() {
    let file = std::fs::read("input.txt").unwrap();
    let part1 = file
        .split(|c| *c == b'\n')
        .filter(|s| s.len() > 0)
        .map(|s| s.split_at(s.len() / 2))
        .map(|(p1, p2)| {
            p1.iter()
                .filter(|c| p2.contains(c))
                .map(|c| {
                    if *c >= b'a' {
                        (c - b'a' + 1) as i16
                    } else {
                        (c - b'A' + 27) as i16
                    }
                })
                .next()
                .unwrap()
        })
        .sum::<i16>();

    println!("Part 1: {}", part1);

    let part2 = file
        .split(|c| *c == b'\n')
        .filter(|s| s.len() > 0)
        .collect::<Vec<&[u8]>>()
        .chunks(3)
        .map(|w| {
            let mut c = w.to_vec();
            c.sort_by(|a: &&[u8], b| a.len().cmp(&b.len()));
            c[0].iter()
                .find(|b| c[1].contains(b) && c[2].contains(b))
                .unwrap()
        })
        .map(|c| {
            if *c >= b'a' {
                (c - b'a' + 1) as i16
            } else {
                (c - b'A' + 27) as i16
            }
        })
        .sum::<i16>();

    println!("Part 2: {}", part2);
}
