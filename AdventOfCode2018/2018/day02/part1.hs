import System.Environment
import qualified Data.Map as Map

exactlyTwoThree :: [(Char, Int)] -> (Int, Int)
exactlyTwoThree = f (0, 0)
    where
        f :: (Int, Int) -> [(Char, Int)] -> (Int, Int)
        f tt [] = tt
        f (a, b) ((_,n):xs)
            | n == 2 = f (1, b) xs
            | n == 3 = f (a, 1) xs
            | otherwise = f (a, b) xs

freq :: String -> [(Char, Int)]
freq s = Map.toList $ Map.fromListWith (+) [(c, 1) | c <- s]

main = do
    [f] <- getArgs
    input <- readFile f
    let (b2, b3) = foldr sum (0,0) $ map exactlyTwoThree $ map freq (lines input)
    print (b2*b3)
    where
        sum = (\(a,b) (x, y) -> (a+x, b+y))