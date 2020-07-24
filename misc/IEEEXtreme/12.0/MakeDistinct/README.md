## Make Distinct

Your are given an array of N integers, A. An operation involves taking a number from the array and either adding or subtracting 1 from it. What is the minimum number of operations to make A have only unique elements.

### Standard input

The first line contains an integer N.
The second line contains N elements, representing A.

### Standard output

Print the answers on the first line

### Constraints and notes

1 ≤ N ≤ 2 * 10^5

1 ≤ A<sub>i</sub> ≤ N for all valid i

    Input           Output  Explanation
    7               7       Transform into 2 4 5 -1 1 0 3.
    2 3 4 2 2 1 3

    4               4       Transform into 5 4 3 2.
    3 3 3 3

    5               4       Transform 6 3 4 5 2.
    5 4 4 5 4