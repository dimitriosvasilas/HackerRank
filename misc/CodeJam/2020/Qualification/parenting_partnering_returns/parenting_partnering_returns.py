import sys

class Parent:
  def __init__(self):
    self.isFree = True
    self.endTime = -1

class Activity:
  def __init__(self, id, startTime, endTime):
    self.id = id
    self.startTime = startTime
    self.endTime = endTime
    self.parent = ""

class Input:
  def __init__(self):
    self.cameron = Parent()
    self.jamie = Parent()
    self.activities = []

  def schedule(self):
    schedule = ""
    for activity in self.activities:
      if not self.cameron.isFree and self.cameron.endTime <= activity.startTime:
        self.cameron.isFree = True
      if not self.jamie.isFree and self.jamie.endTime <= activity.startTime:
        self.jamie.isFree = True
      if self.cameron.isFree:
        activity.parent = "C"
        self.cameron.isFree = False
        self.cameron.endTime = activity.endTime
      elif self.jamie.isFree:
        activity.parent = "J"
        self.jamie.isFree = False
        self.jamie.endTime = activity.endTime
      else:
        return False
    return True

def getInput(t):
  inpArr = [Input() for i in range(t)]
  for inp in inpArr:
    line = input().strip()
    n = int(line)

    for i in range(n):
      line = input()
      activity = [int(x) for x in line.split(" ")]
      inp.activities.append(Activity(i, int(activity[0]), int(activity[1])))

    inp.activities = sorted(inp.activities, key=lambda x:x.startTime)
  return inpArr

def printOutput(tt, schedulePossible, inp):
  if not schedulePossible:
    print("Case #%d: %s" % ((tt+1), "IMPOSSIBLE"))
  else:
    inp.activities = sorted(inp.activities, key=lambda x:x.id)
    schedule = ""
    for act in inp.activities:
      schedule += act.parent
    print("Case #%d: %s" % ((tt+1), schedule))



def main():
  line = input().strip()
  t = int(line)

  inp = getInput(t)

  for tt in range(t):
    schedulePossible = inp[tt].schedule()
    printOutput(tt, schedulePossible, inp[tt])

if __name__ == "__main__":
    main()
