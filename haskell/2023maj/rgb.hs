import System.IO
import Data.List (findIndex)
import Data.Maybe (fromJust)

type Pixel = (Int, Int, Int)

toPixels:: [Int] -> [Pixel]
toPixels [] = []
toPixels (x:y:z:xs) = (x, y, z): toPixels xs

toInt :: [String] -> [Int]
toInt = map read

-- 5. Feladat
hatar :: [Pixel] -> Int -> Bool
hatar [] _ = False
hatar [_] _ = False
hatar (x:y:xs) elteres
    | abs (blue x - blue y) > elteres = True
    | otherwise = hatar (y:xs) elteres
        where
            blue (_, _, z) = z


main :: IO()
main = do
    -- 1. Feladat: Beolvasás
    handle <- openFile "../../input/2023maj/kep.txt" ReadMode
    contents <- hGetContents handle
    let pixels = map (toPixels . toInt . words) $ lines contents

    -- 2. Feladat
    putStrLn "2. feladat:\nKérem egy képpont adatait!"
    putStrLn "Sor: "
    sor <- readLn :: IO Int
    putStrLn "Oszlop: "
    oszlop <- readLn :: IO Int
    putStrLn ("A képpont színe RGB" ++ show (pixels !! (sor - 1) !! (oszlop - 1)))

    -- 3. Feladat
    putStr "3. feladat:\nA világos képpontok száma: "
    print (length $ filter (\(x,y,z) -> x + y + z > 600) $ concat pixels)

    -- 4. Feladat
    putStr "4. feladat:\nA legsötétebb pont RGB összege: "
    let legkisebb = minimum $ map (\(x,y,z) -> x + y + z) $ concat pixels
    print legkisebb

    putStrLn "A legsötétebb pixelek színe:"
    mapM_ print $ filter (\(x,y,z) -> x+y+z == legkisebb) $ concat pixels

    -- 5. Feladat a 'main' fölött

    -- 6. Feladat
    putStrLn "6. feladat:\nA felhő legfelső sora:"
    print (1 + fromJust (findIndex (`hatar` 10) pixels))

    putStrLn "A felhő legalsó sora:"
    print (length pixels - fromJust (findIndex (`hatar` 10) $ reverse pixels))


    hClose handle