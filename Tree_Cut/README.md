### Cut the Tree

Given an N node tree, where each node has a data value, find the edge that when cut, it separates the tree into two trees
t<sub>1</sub> and t<sub>2</sub>, so that the difference |sum(t<sub>1</sub>) - sum(t<sub>2</sub>)| in minimal, where the sum
of a tree is the sum of the data values of its nodes.

#### Input
First line : N - the number of nodes in the tree  
Second line : N space-separated integers denoting the data values of the nodes  
Next N-1 lines : each line contains two space-separated integers, u and v, describing the edge u<->v

#### Output
The minimum difference for the given tree.

#### Notes
The solution can be computed by a recursice solution where the data value of each node is calculated as the sum of the data 
values of its subtree.
