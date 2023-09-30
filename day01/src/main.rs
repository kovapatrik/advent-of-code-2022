fn main() {
    let file = std::fs::read_to_string("input.txt").unwrap();

    let part1 = file
        .split("\n\n")
        .map(|s| s.lines().map(|n| n.parse::<i32>().unwrap()).sum::<i32>())
        .max()
        .unwrap();

    println!("Part 1: {}", part1);

    let mut part2 = file
        .split("\n\n")
        .map(|s| s.lines().map(|n| n.parse::<i32>().unwrap()).sum::<i32>())
        .collect::<Vec<i32>>();
    part2.sort();

    println!("Part 2: {}", part2.iter().rev().take(3).sum::<i32>());
}
