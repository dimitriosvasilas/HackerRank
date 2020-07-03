import Data.Char
import System.Environment
import Data.List.Split

parse input = map (\x -> map toInt x) $ map (\x -> splitOn "\t" x) $ splitOn "\n" input
  where toInt x = read x :: Int

checksum spreadsheet = sum $ map minmax spreadsheet
  where minmax l = maximum l - minimum l

main :: IO ()
main = do
  args <- getArgs
  content <- readFile (args !! 0)
  let input = parse content
  print (checksum input)