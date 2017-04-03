### Moon Journey
Given N people, labeled 0 to N-1, and a number of pairs indicating that two persons are citizens of the same country, calculate
the number of possible ways to choose a pair of people from two different countries.

#### Input
First line: two integers N, P  
Next P lines: two space-separated integers denoting the labes of two peopple from different countries

#### Output
An integers denoting the number of ways two choose a pair of people from different countries

#### Notes
The problem can be solved as follows:  
The number of possible pairs that N people can form is N*(N-1)/2.  
If c people are citizens of the same country, this substracts c*(c-1)/2 pairs from the total possible pairs.  
Calculate the number of nodes in each connected component of the graph created by the given pairs, and apply the above
calculation.
