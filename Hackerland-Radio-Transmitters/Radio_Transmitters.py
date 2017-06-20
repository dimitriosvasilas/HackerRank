n,k = raw_input().strip().split(' ')
n,k = [int(n),int(k)]
x = map(int,raw_input().strip().split(' '))
x.sort()
cnt = 0
curr = 0
i = 0
placed = False
while (i<n):
    if (x[i]-x[curr] > k):
        
        if (not placed):
            cnt += 1
            placed = True
            curr = i-1
        else:
            placed = False
            curr = i
        i -= 1
    i += 1

if (not placed):
    cnt += 1
    
print cnt
        
