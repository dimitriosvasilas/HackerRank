import System.Environment

manhattan 1 = 0
manhattan n | n > (2*k+1)^2-2*k = downEdge n k
            | n > (2*k+1)^2-4*k = leftEdge n k
            | n > (2*k+1)^2-6*k = upEdge n k
            | otherwise = rightEdge n k
            where k = floor $ (sqrt (fromIntegral n - 1) + 1) / 2

downEdge n k = dist (-k) (-k + (2*k+1)^2 - n)

leftEdge n k = dist (-k + (2*k+1)^2-2*k - n) k

upEdge n k = dist k (k - (2*k+1)^2+4*k + n)

rightEdge n k = dist (k - (2*k+1)^2+6*k + n) (-k)

dist x y = abs x + abs y

parse str = toInt str
  where toInt x = read x :: Int

main :: IO ()
main = do
  args <- getArgs
  str <- readFile (args !! 0)
  let input = parse str
  print (manhattan input)