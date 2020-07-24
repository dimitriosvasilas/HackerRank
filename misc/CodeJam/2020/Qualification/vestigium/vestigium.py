import sys

class Input:
  def __init__(self):
    self.n = 0
    self.arr = []

  def computeTrace(self):
    trace = 0
    for i in range(self.n):
      trace += self.arr[i][i]
    return trace

  def isRowNatural(self, rowID):
      freq = {}
      for i in range(self.n):
        if (self.arr[rowID][i] in freq):
            return 1
        else:
            freq[self.arr[rowID][i]] = 1
      return 0

  def isColNatural(self, colID):
      freq = {}
      for i in range(self.n):
        if (self.arr[i][colID] in freq):
            return 1
        else:
            freq[self.arr[i][colID]] = 1
      return 0


  def countRowsRepeating(self):
    count = 0
    for i in range(self.n):
      count += self.isRowNatural(i)
    return count

  def countColsRepeating(self):
    count = 0
    for i in range(self.n):
      count += self.isColNatural(i)
    return count


def getInput(t):
  inpArr = [Input() for i in range(t)]
  for inp in inpArr:
    line = input().strip()
    inp.n = int(line)

    for i in range(inp.n):
      line = input()
      inp.arr.append([int(x) for x in line.split(" ")])

  return inpArr

def printOutput(tt, trace, rowsRepeating, colsRepeating):
  print("Case #%d: %d %d %d" % ((tt+1), trace, rowsRepeating, colsRepeating))


def main():
  line = input().strip()
  t = int(line)

  inp = getInput(t)

  for tt in range(t):
      printOutput(tt, inp[tt].computeTrace(), inp[tt].countRowsRepeating(), inp[tt].countColsRepeating())

if __name__ == "__main__":
    main()
