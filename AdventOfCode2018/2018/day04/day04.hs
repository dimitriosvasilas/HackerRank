import System.Environment

main = do
    [f] <- getArgs
    input <- readFile f
    let parsed = map parseLogEntry (lines input)
    print parsed

parseLogEntry :: String -> (Int, Int)
parseLogEntry _ = (0,0)