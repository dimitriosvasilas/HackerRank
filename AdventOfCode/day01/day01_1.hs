import Data.Char
import System.Environment

pairs :: [Int] -> [(Int,Int)]
pairs l =
  let circ_l = l ++ [head l] in
    zip circ_l (tail circ_l)

captcha :: [Int] -> Int
captcha l = 
  sum (map (\(x,y) -> (if x == y then x else 0)) (pairs l))

main :: IO ()
main = do
  args <- getArgs
  str <- readFile (args !! 0)
  let input = map digitToInt str
  print (captcha input)