### Jack goes to Rapture

[Problem on HackerRank](https://www.hackerrank.com/challenges/jack-goes-to-rapture)

Given the following information about a city's transportation system:
* the number of stations N
* the fare between each pair of connected stations  

calculate the most cost efficient path from the first to the last station.  
The fare from station A to station B, is the difference between the fare from A to B and the cumulative fare that paid to 
reach station A (fare(A,B) - total fare to reach station A). If the difference is negative, then the cost is zero.

#### Input
First line: two integers N, E.  
Next E lines: three integers, two stations that are connected and the fare between them.

#### Output
The minimum fare. If the last station cannot be reached from the first, print "NO PATH EXiSTS".
