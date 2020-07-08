import sys

class Input:
  def __init__(self):
    self.arr = []

  def nest(self):
    string = ""
    currentDepth = 0
    for i in self.arr:
      while currentDepth < i:
        string += "("
        currentDepth += 1
      while currentDepth > i:
        string += ")"
        currentDepth -= 1
      string += str(i)
    while currentDepth > 0:
        string += ")"
        currentDepth -= 1
    return string

def getInput(t):
  inpArr = [Input() for i in range(t)]
  for inp in inpArr:
    line = input().strip()
    inp.arr = [int(c) for c in line]

  return inpArr

def printOutput(tt, string):
  print("Case #%d: %s" % ((tt+1), string))


def main():
  line = input().strip()
  t = int(line)

  inp = getInput(t)

  for tt in range(t):
    printOutput(tt, inp[tt].nest())

if __name__ == "__main__":
    main()
