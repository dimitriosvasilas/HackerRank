s = raw_input()
length = len(s)
sumsubs = int(s[0])
res = sumsubs
for i in range(1, length):
	sumsubs = int(s[i]) * (i+1) + sumsubs*10
	res = res + sumsubs

print res
