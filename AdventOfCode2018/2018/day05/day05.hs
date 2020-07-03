import System.Environment
import Data.Char

main = do
    [f] <- getArgs
    input <- filter isLetter <$> readFile f
    print $ part1 input
    print $ part2 input

remove :: [Char] -> Char -> [Char]
remove s x = filter (\c -> toLower c /= x) s

part1 :: [Char] -> Int
part1 s = length $ react' s

part2 :: [Char] -> Int
part2 s = foldl1 min
    . map length
    . map react'
    $ map (\c -> remove s' c)
    $ filter (\c -> toUpper c `elem` s' || toLower c `elem` s') ['a'..'z']
        where s' = react' s

react' :: [Char] -> [Char]
react' = foldl f []
    where
        f :: [Char] -> Char -> [Char]
        f [] c = [c]
        f (x:xs) c
            | x /= c && toUpper x == toUpper c = xs
            | otherwise = c:x:xs

react :: [Char] -> [Char]
react s
    | length (reactPass s) < length s = react $ reactPass s
    | otherwise = s

reactPass :: [Char] -> [Char]
reactPass (x:y:xs)
    | x /= y && toUpper x == toUpper y = reactPass xs
    | otherwise = x : reactPass (y:xs)
reactPass s = s