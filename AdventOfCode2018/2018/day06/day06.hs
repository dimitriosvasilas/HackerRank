import System.Environment
import qualified Data.Map as Map
import Data.List
import Data.Function (on)
import Data.Ord
import Data.List.Split

main = do
    [f] <- getArgs
    input <- readFile f
    print $ part1 input

part1 :: String -> Int
part1 s = snd
    . maximumBy (comparing snd)
    -- compute how many coordinates are closest to each input coordiate
    . freq
    -- for each coordinate in the rectangle, find which input coordinate it is
    -- closest to
    . map (\p -> closest coordsId p)
    -- generate all the coordinates contained the rectangle
    . region
    -- TODO: need top left, bottom right coordinates
    -- computute the conrers of the rectangle that the input coordinates form
    $ bounds coordsId
    where
        -- give each input coordinate a unique id
        coordsId = zip coords [0 .. length coords]
            where
                -- parse the input to a list of input coordinates
                coords = map parse $ lines s

-- Parse a comma delimited coordinate
parse :: [Char] -> (Int, Int)
parse s = f . map read $ filter (\s -> length s > 0) $ splitOneOf ", " s
    where
        f [a, b] = (b, a)

-- Given a list of coordinate, computes the corners of a the smallest rectangle
-- that contains all the coordinates
bounds :: [((Int, Int), Int)] -> (Int, Int, Int, Int)
bounds l = (minimum (map (fst . fst) l), minimum (map (snd . fst) l), maximum (map (fst . fst) l), maximum (map (snd . fst) l))

-- Given the corners of a rectangle, generates all the coordinates contained
-- in that rectangle
region :: (Int, Int, Int, Int) -> [(Int, Int)]
region (minx, miny, maxx, maxy) = [(x,y) | x <- [minx .. maxx], y <- [miny .. maxy]]

-- Computes the manhattan distance between two points
manhattan :: (Int, Int) -> (Int, Int) -> Int
manhattan (x,y) (a,b) = abs (x-a) + abs (y-b)

-- Given a list of input coordinates and another coordinate,
-- finds the input coordinate that is closent to the given one.
closest :: [((Int, Int), Int)] -> (Int, Int) -> ((Int, Int), Int)
closest p (x,y) = f . sortBy (compare `on` fst) $ map (\((a,b),id) -> (manhattan (x,y) (a,b), id)) p
    where
        f ((d1,id1):(d2,id2):xs)
            | d1 == d2 = ((x, y), -1)
            | otherwise = ((x, y), id1)

-- Given a list of coordinates, and which input coordinate is closest to each
-- coordinate, computes the number of coordinates that are closest to each of
-- the input coordinates
freq :: [((Int,Int),Int)] -> [(Int, Int)]
freq l = Map.toList $ Map.fromListWith (+) [(id, 1) | ((_,_),id) <- l]