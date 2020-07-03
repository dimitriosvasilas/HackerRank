import Data.Char
import System.Environment
import Data.List.Split
import Data.List (nub)

validPassphrases = length . filter (\x -> isValid x)
  where isValid pass = length (pass) == length (nub pass)
   
parse str = map (\x -> splitOn " " x) $ splitOn "\n" str

main :: IO ()
main = do
  args <- getArgs
  content <- readFile (args !! 0)
  let input = parse content
  print (validPassphrases input)