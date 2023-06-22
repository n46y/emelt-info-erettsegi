#include <iostream>
#include <fstream>
#include <vector>

int main() {
    // 1. Feladat: Beolvasás
    std::vector<std::vector<int[3]>> adat;

    std::ifstream file("../../input/2023maj/kep.txt");

    if(file.is_open()){
        std::string szam;
        while (std::getline(file, szam, ' ')){
            std::cout << szam << ", ";
        }
    }

    return 0;
}