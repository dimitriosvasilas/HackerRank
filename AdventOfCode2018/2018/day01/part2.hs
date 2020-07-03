import System.Environment
import qualified Data.Set as Set

strToInt :: String -> Int
strToInt ('+':xs) = read xs
strToInt xs = read xs

solve :: [Int] -> Int
solve = f Set.empty 0 . cycle
    where
        f :: Set.Set Int -> Int -> [Int] -> Int
        f s fr (x:xs)
            | fr `Set.member` s = fr
            | otherwise = f (Set.insert fr s) (fr+x) xs

main = do
    [f] <- getArgs
    input <- readFile f
    print . solve $ map strToInt (lines input)
