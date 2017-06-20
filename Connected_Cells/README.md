### Connected Cells in a Grid

[Problem on HacherRank](https://www.hackerrank.com/challenges/connected-cell-in-a-grid)

Given an N x M matrix, where:
- each cell of the grid contains either 0 or 1. Cells containing 1 are called filled cells.
- two cells are connected if they are adjacent to each other horizontally, vertically, or diagonally
- connected cells that are also filled form a region


find the number of cells in the largest region in the matrix

#### Input
First line : N - the number of row in the matrix  
Second line : M - the number of columns in the matrix  
Next N lines : the matrix

#### Output
Number of cells in the largest region in the given array

#### Notes
The problem is equivalent with the following:  
Find the number of nodes in the largest connected component of an undirected graph
