import math
def findLevel(x):
    return int(math.log(x,2))
    
def findLca(a, b):
    while a != b:
        if a < b:
            b /= 2
        else:
            a /= 2
    return a

def findDist(a, b):
    lca = findLca(a, b)
    d_lca = findLevel(lca)
    d_a = findLevel(a)
    d_b = findLevel(b)
    return d_a + d_b - 2 * d_lca

q = int(raw_input())
for i in xrange(q):
    a, b = map(int,raw_input().split())
    print findDist(a,b)
