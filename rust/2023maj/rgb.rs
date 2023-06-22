use std::{fs, env, io};

#[derive(Debug)]
struct RGB (i32, i32, i32);

fn main() {
    // 1. Feladat: beolvasás
    let input_file = env::current_dir().unwrap().join("../../input/2023maj/kep.txt");

    let pixels: Vec<Vec<RGB>> = fs::read_to_string(input_file).unwrap()
        .trim()
        .split('\n')
        .map(|line| {
            let l: Vec<&str> = line.trim().split(' ').collect();
            l.chunks(3)
                .map(|chunk| {
                    RGB(
                        chunk[0].parse().unwrap(),
                        chunk[1].parse().unwrap(),
                        chunk[2].parse().unwrap()
                    )
                }).collect()
        }).collect();

    // 2. Feladat
    println!("2. feladat:\nKérem egy képpont adatait!");

    println!("Sor: ");
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let sor: usize = input.clone().trim().parse().unwrap();

    println!("Oszlop: ");
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let oszlop: usize = input.clone().trim().parse().unwrap();

    println!("A képpont színe {:?}", pixels[sor-1][oszlop-1]);


    // 3. Feladat
    println!("3. feladat:\nVilágos képpontok száma: {}",
           pixels.iter().flatten()
           .filter(|p| {
              p.0 + p.1 + p.2 > 600
           }).count()
    );


    // 4. Feladat
    let min = pixels.iter().flatten()
        .map(|p| {
            p.0 + p.1 + p.2
        }).min().unwrap();

    println!("4. feladat:\nA legsötétebb pont RGB összege: {}", min);

    println!("A legsötétebb pixelek színe:");
    for i in pixels.iter().flatten().filter(|p| p.0 + p.1 + p.2 == min)
    { println!("{:?}", i); }

    // 6. Feladat
    println!("6. feladat:");
    let (e, _) = pixels.iter().enumerate()
        .filter(|(_, l)| hatar(l, 10))
        .next().unwrap();
    let (u, _) = pixels.iter().enumerate()
        .filter(|(_, l)| hatar(l, 10))
        .last().unwrap();
    println!("A felhő legfelső sora: {}", e + 1);
    println!("A felhő legalsó sora: {}", u + 1);

}

// 5. Feladat
fn hatar(sor: &Vec<RGB>, elteres: i32) -> bool {
     if 0 < sor.iter()
        .map(|p| p.2 )
        .zip(sor[1..sor.len()].iter().map(|p| p.2))
        .filter(|(a, b)| (a-b).abs() > elteres)
        .count()
    { return true; }

    false
}
