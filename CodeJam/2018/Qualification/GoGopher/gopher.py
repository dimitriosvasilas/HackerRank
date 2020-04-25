import sys

def upper_filled(Arr, center_x, center_y):
    if Arr[center_x-1][center_y-1] == 0:
        return False
    if Arr[center_x-1][center_y] == 0:
        return False
    if Arr[center_x-1][center_y+1] == 0:
        return False
    return True

def solve(Arr, A, center_x, center_y, count):
    if count >= A:
        return (center_x, center_y, count)
    if upper_filled(Arr, center_x, center_y):
        return (center_x+1, center_y, count+3)
    return (center_x, center_y, count)

def main():    
    T = int(input())
    for _ in range(T):
        A = int(input())
        
        Arr = [[0 for x in range(1000)] for y in range(1000)]
        center_x = 3
        center_y = 3
        count = 9
        
        while True:
            sys.stdout.write(str(center_x) + ' ' + str(center_y) + '\n')
            sys.stdout.flush()
            
            x, y = map(int, input().split(' '))
            if x == 0 and y == 0:
                break
            elif x == -1 and y == -1:
                sys.exit(0)
            else:
                Arr[x][y] = 1
                (center_x, center_y, count) = solve(Arr, A, center_x, center_y, count)            


if __name__ == '__main__':
    main()