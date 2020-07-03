import System.Environment
import Data.List

commonSublist :: [Char] -> [Char] -> [Char]
commonSublist [] [] = []
commonSublist (x:xs) (y:ys)
    | x == y = [x] ++ commonSublist xs ys
    | otherwise = commonSublist xs ys

solve :: [[Char]] -> [Char]
solve [_] = ""
solve (x:y:xs)
    | length(x) - length(xy) == 1 = xy
    | otherwise = solve (y:xs)
    where
        xy = commonSublist x y

main = do
    [f] <- getArgs
    input <- readFile f
    print . solve . sort $ lines input
