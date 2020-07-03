import Data.Char
import System.Environment

rotate :: [Int] -> Int -> [Int]
rotate [] _ = []
rotate xs 0 = xs
rotate (x:xs) n = rotate (xs ++ [x]) (n-1)

pairs :: [Int] -> [(Int,Int)]
pairs l =
  let rot_l = rotate l (length(l) `div` 2) in
    zip l rot_l

captcha :: [Int] -> Int
captcha l = 
  sum (map (\(x,y) -> (if x == y then x else 0)) (pairs l))

main :: IO ()
main = do
  args <- getArgs
  str <- readFile (args !! 0)
  let input = map digitToInt str
  print (captcha input)