# 5. Feladat
def hatar(sor, elteres):
    for i in range(len(adat[sor]) - 1):
        kulonbseg = abs(adat[sor][i][2] - adat[sor][i+1][2])
        if kulonbseg > elteres:
            return True

    return False


if __name__ == "__main__":
    # 1. Feladat: beolvasás
    adat = []

    with open("kep.txt", "r", encoding="utf-8") as file:
        for sor in file:
            adat.append([])
            for idx, szam in enumerate(sor.strip().split(" ")):
                if idx % 3 == 0:
                    adat[-1].append([int(szam)])
                else:
                    adat[-1][-1].append(int(szam))

    # 2. Feladat
    print("2. feladat:")
    print("Kérem egy képpont adatait!")

    sor = int(input("Sor: ")) - 1
    oszlop = int(input("Oszlop: ")) - 1
    print(f"A képpont színe RGB({adat[sor][oszlop][0]},{adat[sor][oszlop][1]},{adat[sor][oszlop][2]})")

    # 3. Feladat
    print("3. feladat:")

    vilagos_pontok = 0
    for sor in adat:
        for pixel in sor:
            if sum(pixel) > 600:
                vilagos_pontok += 1

    print(f"A világos képpontok száma: {vilagos_pontok}")

    # 4. Feladat
    print("4. feladat:")

    legsotetebb_pont = 3 * 255
    for sor in adat:
        for pixel in sor:
            if sum(pixel) < legsotetebb_pont:
                legsotetebb_pont = sum(pixel)

    print(f"A legsötétebb pont RGB összege: {legsotetebb_pont}")

    print("A legsötétebb pixelek színe:")
    for sor in adat:
        for pixel in sor:
            if sum(pixel) == legsotetebb_pont:
                print(f"RGB({pixel[0]},{pixel[1]},{pixel[2]})")

    # 6. Feladat
    print("6. feladat:")

    i = 0
    while not hatar(i, 10):
        i += 1
    print(f"A felhő legfelső sora: {i + 1}")

    i = len(adat) - 1
    while not hatar(i, 10):
        i -= 1
    print(f"A felhő legalsó sora: {i + 1}")
