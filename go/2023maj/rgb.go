package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pixel struct {
    red int
    green int
    blue int
}

func (p Pixel) sum() int {
    return p.red + p.blue + p.green
}

func main()  {
    // 1. Feladat: beolvasás
    file, _ := os.Open("../../input/2023maj/kep.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

    pixels := make([][]Pixel, 0)

    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        pixels = append(pixels, make([]Pixel, 0))
        for i := 0; i < len(line); i+=3 {
            red, _ := strconv.Atoi(line[i])
            green, _ := strconv.Atoi(line[i+1])
            blue, _ := strconv.Atoi(line[i+2])
            pixels[len(pixels)-1] = append(pixels[len(pixels)-1], Pixel{red, green, blue})
        }
    }

    // 2. Feladat
    fmt.Println("Kérem egy képpont adatait!")
    var input string;

    fmt.Print("Sor: ")
    fmt.Scanln(&input)
    sor, _ := strconv.Atoi(strings.TrimSpace(input))

    fmt.Print("Oszlop: ")
    fmt.Scanln(&input)
    oszlop, _ := strconv.Atoi(strings.TrimSpace(input))

    fmt.Printf("A képpont színe RGB(%v, %v, %v)\n",
        pixels[sor-1][oszlop-1].red,
        pixels[sor-1][oszlop-1].green,
        pixels[sor-1][oszlop-1].blue,
    )

    // 3. Feladat
    vilagos := 0
    for _, s := range pixels {
        for _, p := range s {
            if p.sum() > 600 {
                vilagos++
            }
        }
    }

    fmt.Printf("3. feladat:\nVilágos képpontok száma: %v\n", vilagos)

    // 4. Feladat
    min := math.MaxInt
    for _, s := range pixels {
        for _, p := range s {
            if m := p.sum(); m < min {
                min = m
            }
        }
    }

    fmt.Printf("4. feladat:\nA legsötétebb pont RGB összege: %v\n", min)
    fmt.Println("A legsötétebb pixelek színe:")

    for _, s := range pixels {
        for _, p := range s {
            if p.sum() == min {
                fmt.Printf("RGB(%v, %v, %v)\n", p.red, p.green, p.blue)
            }
        }
    }

    // 6. Feladat
    var legfelso int;
    var legalso int;

    for i := 0; i < len(pixels); i++{
        if hatar(&pixels, i, 10) {
            legfelso = i + 1
            break
        }
    }

    for i := len(pixels)-1; i > 0; i--{
        if hatar(&pixels, i, 10) {
            legalso = i + 1
            break
        }
    }

    fmt.Printf("6. feladat:\nA felhő legfelső sora: %v\nA felhő legalsó sora: %v\n", legfelso, legalso)
}

// 5. Feladat 
func hatar(pixels *[][]Pixel, sor int, elteres int) bool {
    for i := 0; i < len((*pixels)[sor])-1; i++{
        if int(math.Abs(float64((*pixels)[sor][i].blue - (*pixels)[sor][i+1].blue))) > elteres {
            return true
        }
    }    
    return false
}
