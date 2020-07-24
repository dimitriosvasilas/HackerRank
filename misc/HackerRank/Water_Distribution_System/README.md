Good evening master Wayne.

Joker and his gang attended Black Hat USA 2017 where they learned of a new way on how to damage our city! Specifically, tomorrow night they will try to damage the water pumps of Gotham using bubbles!

The bubbles cause corrosion to the pumps and in a few hours they will damage them with catastrophic results! To dash Joker’s plan, besides stopping him, you need to make sure that the city network does not contain loops.

If Joker manages to inject bubbles to the network and they enter a loop, they will still cause damage to that area even though you would have already arrested Joker and his gang.

Given the map of the water distribution system, you need to make sure that the map does not contain loops.

Standard input
On the first line there will be an integer tt, the number of test cases to follow.

For each test case, there will be 22 input lines:

On the first line of the test case, there will be 2 integers nn and mm, where nn is the number of vertices and mm is the number of edges.
On the second line, there will be mm pairs of integers separated by a space character. Each pair shows a two way connection between vertex aa and vertex bb.

Standard output
For each test case you will have to write one line that contains an integer, in the case where there is a loop you will write the number 11 or else you will write the number 00.

Constraints and notes
1 \le t \le 1\,0001≤t≤1000 
1 \le n \le 1\,0001≤n≤1000 
1 \le m \le 10\,0001≤m≤10000 
0 \le a,b \le n-10≤a,b≤n−1 
There can be multiple edges or self-loops. In this case we consider the graph to contain a loop.


