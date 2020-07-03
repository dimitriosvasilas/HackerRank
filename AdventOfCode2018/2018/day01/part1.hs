import System.Environment

strToInt :: String -> Int
strToInt ('+':xs) = read xs
strToInt xs = read xs

main = do
    [f] <- getArgs
    input <- readFile f
    print . sum $ map strToInt (lines input)
