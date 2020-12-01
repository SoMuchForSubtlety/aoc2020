main :: IO ()
main = do
  infile <- readFile "input/day01.txt"
  let xs = [read x :: Int | x <- lines infile]
  print (part1 xs)
  print (part2 xs)

part1 xs = do
  let xsi = zip [0 ..] xs
  head [x * y | (i, x) <- xsi, (j, y) <- xsi, i /= j, x + y == 2020]

part2 xs = do
  let xsi = zip [0 ..] xs
  head [x * y * z | (i, x) <- xsi, (j, y) <- xsi, i /= j, x + y < 2020, (l, z) <- xsi, l /= i, x + y + z == 2020]
