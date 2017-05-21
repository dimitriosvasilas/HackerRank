### BFS: Shortest Reach
Given an undirected graph of N nodes, labeled from 1 to N, and the edge between two nodes is always of length 6, calculate the
shortest distance from starting node s to all other nodes in the graph.

#### Input

First line: an integer denoting the number of queries  
The subsequent lines describe each query in the following format:  
First line: two space-separated integers describing the the number of nodes and the number of edges in the graph.  
Next m lines: two space-separated integers describing an edge connecting two nodes  
Last line: an integer denoting the starting node

### Output
For each of query, a line of space-separated integers denoting the shortest distances from starting node to each of the other
nodes. Distances should be listed sequentially by node number and should not include the starting node. If some node is unreachable, print -1 as the distance to that node.

