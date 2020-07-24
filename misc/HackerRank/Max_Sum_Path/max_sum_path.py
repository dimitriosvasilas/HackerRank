forest = []
r, c = map(int,raw_input().split())
for i in xrange(r):
    forest.append(map(int, raw_input().split(' ')))
maxValues = []
pathValues = []
for i in xrange(r):
    maxValues.append([None]*c)
for i in xrange(r):
    pathValues.append([None]*c)

for i in xrange(len(forest)):
    for j in xrange(len(forest[i])):
        if i == 0 and j == 0:
            maxValues[i][j] = forest[i][j]
        elif i == 0:
            pathValues[i][j] = False
            maxValues[i][j] = maxValues[i][j-1] + forest[i][j]
        elif j == 0:
            pathValues[i][j] = True
            maxValues[i][j] = maxValues[i-1][j] + forest[i][j]
        else:
            if maxValues[i-1][j] >= maxValues[i][j-1]:
                pathValues[i][j] = True
                maxValues[i][j] = maxValues[i-1][j] + forest[i][j]
            else:
                pathValues[i][j] = False
                maxValues[i][j] = maxValues[i][j-1] + forest[i][j]

i = r-1
j = c-1
path = []
while True:
    if pathValues[i][j] == True:
        i -= 1
    else:
        j -= 1
    if i == 0 and j == 0:
        break
    path.append((i,j))
    
path.reverse()
#print path

health = 1
sumadd = 0
for move in path:
    health += forest[move[0]][move[1]]
    #print move[0], move[1], 'health', health
    if (health < 1):
        add = abs(1-health)
        sumadd += add
        health += add

print sumadd+1

