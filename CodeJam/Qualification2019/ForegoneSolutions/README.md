## Foregone Solutions
### Problem
Someone just won the Code Jam lottery, and we owe them N jamcoins! However, when we tried to print out an oversized check, we encountered a problem. The value of N, which is an integer, includes at least one digit that is a 4... and the 4 key on the keyboard of our oversized check printer is broken.

Fortunately, we have a workaround: we will send our winner two checks for positive integer amounts A and B, such that neither A nor B contains any digit that is a 4, and A + B = N. Please help us find any pair of values A and B that satisfy these conditions.

### Input
The first line of the input gives the number of test cases, T. T test cases follow; each consists of one line with an integer N.

### Output
For each test case, output one line containing Case #x: A B, where x is the test case number (starting from 1), and A and B are positive integers as described above.

It is guaranteed that at least one solution exists. If there are multiple solutions, you may output any one of them.

### Limits
1 ≤ T ≤ 100.
Time limit: 10 seconds per test set.
Memory limit: 1GB.
At least one of the digits of N is a 4.

### Sample

Input  
3  
4  
940  
4444

Output  
Case #1: 2 2  
Case #2: 852 88  
Case #3: 667 3777  