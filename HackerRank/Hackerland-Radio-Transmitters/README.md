### Hackerland Radio Transmitters

[Problem on HackerRank](https://www.hackerrank.com/challenges/hackerland-radio-transmitters)

Given an one-dimensional city with n houses, where each house i is located at some point x<sub>i</sub> on the x-axis, and a set
of radio transmitters with range k than be located at the roofs of city's houses, find the minimum number of transmitters needed
to cover every house in the city.

#### Input
First line: two space-separated integers describing the values of n and k  
Second line: space-separated integers describing the location of each house

#### Output
An integer denoting the minimum number of transmitters needed to cover all the houses.

#### Notes
The problem is equivalent with the following:  
Given n points on the x-axis and an integer k, find the minimum number of intervals of the form [x<sub>i</sub>-k, 
x<sub>i</sub>+k], so that each point x<sub>i</sub> is contained in an interval
