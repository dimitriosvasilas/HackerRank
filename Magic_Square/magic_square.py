n = int(raw_input())
square = []
for i in range(n):
	row = map(int,raw_input().strip().split(' '))
	square.append(row)

magic_sum = 0
anti_diag = 0
cnt = 0
lst = []

for i in range(n):
	magic_sum += square[i][i]
	anti_diag += square[i][n-i-1]

if (anti_diag != magic_sum):
	cnt += 1
	lst.append(0)

for i in range(n):
	rsum = 0
	csum = 0
	for j in range(n):
		rsum += square[i][j]
		csum += square[j][i]
	if (rsum != magic_sum):
		cnt += 1
		lst.append(i+1)
	if (csum != magic_sum):
		cnt += 1
		lst.append(-i-1)
print cnt
if (cnt > 0):
	lst.sort()
	for x in lst:
		print x
