import sys


class Flavor(object):
    def __init__(self):
        self.available = True
        self.frequency = 0


def read_jundge_input():
    inpt = [int(n) for n in input().split()]
    if inpt[0] == -1:
        sys.exit(0)
    return inpt[1:]


def sell(customer, flavors):
    minF = None
    toSell = -1
    for preference in customer:
        flavors[preference].frequency += 1
        if flavors[preference].available:
            if minF == None or flavors[preference].frequency < minF:
                minF = flavors[preference].frequency
                toSell = preference
    if toSell != -1:
        flavors[toSell].available = False
    return toSell


def main():
    T = int(input())
    for _ in range(T):
        N = int(input())

        flavors = []
        for _ in range(N):
            flavors.append(Flavor())

        for _ in range(N):
            customer = read_jundge_input()
            sys.stdout.write(str(sell(customer, flavors)) + '\n')
            sys.stdout.flush()


if __name__ == '__main__':
    main()