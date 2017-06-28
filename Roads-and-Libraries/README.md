### Roads and Libraries

[Problem on HackerRank](https://www.hackerrank.com/challenges/torque-and-development/problem)


Given an undirected graph that represents a network of unconstructed roads between cities, the cost of
building a library in one city, and the cost of reconstructing one of the roads, calculate the minimum 
cost so that all cities have access to a library.

#### Input

First line: an integer denoting the number of queries.
The subsequent lines describe each query in the following format:  
First line: four space-separated integers describing the number of cities (n), the number of roads (m), 
the cost to build a library, and the cost to build a road.
Next m lines: two space-separated integers u and v describing a road connecting the cities u and v.

#### Output

An integer denoting the minimum cost.

#### Notes

The problem is equivalent to the following:
Find the number of cities in each connected component of the graph.
If cost to build a road > cost to build a road then the minimum cost is calculated by 
building one library in each city.
Otherwise, the minimum cost is calculated by building one library in each connected component
and building the minimum number of roads so that all the cities in the component have access to the
library.
