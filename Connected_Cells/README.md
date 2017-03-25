### Connected Cells in a Grid

Given an N x M matrix, find the number of cells in the largest region in the matrix
- each cell of the grid contains either 0 or 1. Cells containing 1 are called filled cells.
- two cells are connected if they are adjacent to each other horizontally, vertically, or diagonally
- connected cells that are also filled form a region

#### Input

First line : N - the number of row in the matrix
Second line : M - the number of columns in the matrix
Next N lines : the matrix

#### Output

Number of cells in the largest region in the given array

#### Notes

The problem can be expressed as follows:
- the matrix represents a graph. Connected and filled cells represent graph edges
- regions represent connected components of the graph
- the task is to count the number of nodes in the largest connected component
