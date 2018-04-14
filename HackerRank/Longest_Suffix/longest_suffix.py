def isSubSequence(str1,str2):
    j = 0
    for i in xrange(len(str2)):
        if str1[j] == str2[i]:
            j += 1
            if j == len(str1):
                return True
    return False

def maxSuffix(s, p):
    max_suff = len(p)
    for i in xrange(max_suff-1, -1, -1):
        #print p[i:], s
        if isSubSequence(p[i:], s) == False:
            #print 'false', len(p[i:])
            max_suff = len(p[i:]) - 1
            break        
    return max_suff    

s = raw_input()
q = int(raw_input())
for i in xrange(q):
    p = raw_input()
    print maxSuffix(s, p)
