def calculate_damage_arr(P, D):
    damage = []
    damage_cur = (0, 1)
    s_count = 0
    for c in P:
        if c == 'C':
            damage_cur = (damage_cur[0], damage_cur[1]*2)
            
        else:
            damage_cur = (damage_cur[0]+damage_cur[1], damage_cur[1])
            s_count += 1
        damage.append(damage_cur)
    if s_count > D:
        return None
    else:
        return damage
        
def damage_after_swap(P, D, damage, index):
    damage_cur = (0, 1)
    if index-2 >= 0:
        damage_cur = (damage[index-2][0], damage[index-2][1])
    j = index-1
    while j < len(P):
        if P[j] == 'C':
            damage_cur = (damage_cur[0], damage_cur[1]*2)
        else:
            damage_cur = (damage_cur[0]+damage_cur[1], damage_cur[1])
        damage[j] = damage_cur
        j += 1
    return damage

def solve(damage, P, D):
    hacks = 0
    while damage[len(damage)-1][0] > D:
        i = len(P)-1
        while i >= 0:
            if P[i] == 'S' and P[i-1] == 'C':
                P[i-1] = 'S'; P[i] = 'C'; hacks += 1
                
                damage = damage_after_swap(P, D, damage, i)
                break
            i -= 1
    return hacks

def main():
    T = int(input())
    for t in range(T):
        D,P = input().split()
        D = int(D); P = list(P)

        damage = calculate_damage_arr(P, D)

        if damage == None:
            print("Case #%d: IMPOSSIBLE" % (t+1))
        else:
            hacks = solve(damage, P, D)
            print("Case #%d: %d" % (t+1, hacks))

        
if __name__ == '__main__':
    main()