import System.Environment
import qualified Data.Map as Map
import Data.List.Split

main = do
    [f] <- getArgs
    input <- readFile f
    print $ part1 input
    print $ part2 input

part1 :: String -> Int
part1 s = Map.size
    . Map.filter (\a -> length a > 1)
    . claims
    $ lines s

part2 :: String -> Int
part2 s = head
    . Map.keys
    $ solve (Map.toList overlapingClaims) ids
        where
            ids = foldr (\ (id,(_,_),(_,_)) -> Map.insert id 0) Map.empty $ map parse (lines s)
            overlapingClaims = Map.filter (\a -> length a > 1) . claims $ lines s

solve :: [((Int,Int),[Int])] -> Map.Map Int Int -> Map.Map Int Int
solve [] ids = ids
solve (((_,_),lid):xs) ids = solve xs (del lid ids)
    where
        del :: [Int] -> Map.Map Int Int -> Map.Map Int Int
        del [] m = m
        del (x:xs) m = del xs (Map.delete x m)

claims :: [String] -> Map.Map (Int,Int) [Int]
claims l = foldr (\ (id,(x,y)) -> Map.insertWith (++) (x,y) [id]) Map.empty
    . foldr (\ x y -> toRect x ++ y) []
    $ map parse l

parse :: String -> (Int, (Int, Int), (Int, Int))
parse s = f. map read $ filter (\s -> length s > 0) $ splitOneOf "# @,:x" s
    where
        f [id, x, y, w, h] = (id, (x, y), (w, h))

toRect :: (Int, (Int, Int), (Int, Int)) -> [(Int, (Int, Int))]
toRect (id, (x, y),(w, h)) = [ (id,(x,y)) | x <- [x .. x + w - 1], y <- [y .. y + h - 1]]
