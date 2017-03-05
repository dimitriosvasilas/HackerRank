#!/bin/python
import sys

# Given n points on the x-axis and an integer k
# find the minimum number of intervals of the form [i-k, i+k]
# so that each point are contained within an interval
# and i is an existing point
# eg. for points 1 2 3 4 5 and k=1
# result is 2 (intervals are [1,3] and [3,5])
# for points 7 2 4 6 5 9 12 11 and k=2
# result is 3 (intervals are [2, 6[, [7, 11] and [10, 14])

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
        
