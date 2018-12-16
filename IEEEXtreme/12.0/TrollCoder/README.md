## Troll Coder

You have found a huge treasure on an island, but the only access point to that island is a bridge guarded by a giant Troll! In order to cross the bridge, you have to guess a sequence of N bits by submitting queries. For each query, the Troll will tell you how many bits you guesses correctly until you guess the correct sequence.

### Interaction

Your program must exchange information with the Troll by submitting queries and reading answers.

Note that you must flush the buffer so that the output reaches the Troll. Here we illustrate it for several languages.

At the beginning of each test case, the Troll will give you a single integer N which will represent the length of the sequence.

To submit a query, your program should output the letter Q followed by a space and by a binary sequence of length N with each bit separated by a space. After each query you will receive an integer denoting the number of correct bits. The last submission will be your final answer and it should start with an A followed by a space and by a binary sequence of length N with each bit separated by a space.

### Constraints and notes
- This task is NOT adaptive
- 1 ≤ N ≤ 100
- Your program can submit at most N + 1 queries before arriving at the correct answer.

Example

    Interaction
    6
                Q 0 0 0 0 0 0
    2
                A 1 0 1 1 0 1
    Explanation
    This sequence has 6 bits and the solution is 101101.
    
    The first query is Q 0 0 0 0 0 0 which has 2 correct bits.
    
    The correct answer is given in the form of A 1 0 1 1 0 1 .