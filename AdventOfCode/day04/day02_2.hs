import Data.Char
import System.Environment
import Data.List.Split

parse input = map (\x -> map toInt x) $ map (\x -> splitOn "\t" x) $ splitOn "\n" input
  where toInt x = read x :: Int

pairs table = map pairsLine table
  where pairsLine l = [(x,y) | x <- l, y <- l, x < y]

checksum spreadsheet = sum (map checksumLine (pairs spreadsheet))
  where checksumLine l = sum (map (\(x,y) -> (if y `rem` x == 0 then y `div` x else 0)) l)

main :: IO ()
main = do
  args <- getArgs
  content <- readFile (args !! 0)
  let input = parse content
  print (checksum input)