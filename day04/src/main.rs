fn main() {
    let file = std::fs::read_to_string("input.txt").unwrap();

    let part1 = file
        .lines()
        .map(|line| {
            let (l, r) = line.split_once(',').unwrap();
            let ((l1, l2), (r1, r2)) = (
                l.split_once('-').unwrap(),
                r.split_once('-').unwrap(),
            );
            (
                l1.parse::<u32>().unwrap(),
                l2.parse::<u32>().unwrap(),
                r1.parse::<u32>().unwrap(),
                r2.parse::<u32>().unwrap(),
            )
        })
        .filter(|(l1, l2, r1, r2)| (l1 >= r1 && l2 <= r2) || (r1 >= l1 && r2 <= l2))
        .count();

    println!("Part 1: {}", part1);

    let part2 = file
        .lines()
        .map(|line| {
            let (l, r) = line.split_once(',').unwrap();
            let ((l1, l2), (r1, r2)) = (
                l.split_once('-').unwrap(),
                r.split_once('-').unwrap(),
            );
            (
                l1.parse::<u32>().unwrap(),
                l2.parse::<u32>().unwrap(),
                r1.parse::<u32>().unwrap(),
                r2.parse::<u32>().unwrap(),
            )
        })
        .filter(|(l1, l2, r1, r2)| l1 <= r2 && r1 <= l2)
        .count();

    println!("Part 2: {}", part2);
}
