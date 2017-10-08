#!/bin/python

import sys

if __name__ == "__main__":
    t = int(raw_input().strip())
    for a0 in xrange(t):
        n = int(raw_input().strip())
        prices = map(int, raw_input().strip().split(' '))
        maxPr = 0
        profit = 0
        bought = 0
        boughtN = 0
        for p in reversed(prices):
            if (p >= maxPr):
                profit += maxPr*boughtN - bought
                maxPr = p
                bought = 0
                boughtN = 0
            else:
                bought += p
                boughtN += 1
        profit += maxPr*boughtN - bought
        print profit
