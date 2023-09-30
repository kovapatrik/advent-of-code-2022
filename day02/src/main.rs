fn main() {
    let file = std::fs::read("input.txt").unwrap();
    let part1 = file
        .split(|c| *c == b'\n')
        .filter(|s| s.len() > 0)
        .map(|s| ((s[0] - b'A') as i16, (s[2] - b'X') as i16))
        .map(|(a,b)| 1 + b + 3 * (1 + b - a).rem_euclid(3))
        .sum::<i16>();

    println!("Part 1: {}", part1);

    let part2 = file
        .split(|c| *c == b'\n')
        .filter(|s| s.len() > 0)
        .map(|s| ((s[0] - b'A') as i16, (s[2] - b'X') as i16))
        .map(|(a,b)| 1 + b * 3 + (2 + a + b) % 3)
        .sum::<i16>();

    println!("Part 2: {}", part2);

}
