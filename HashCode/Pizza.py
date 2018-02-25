import numpy as np
import math

def read_file(filename):
	with open(filename,'r') as f:
		line = f.readline()
		rows, columns, min_ingr, max_cells = [int(n) for n in line.split()]

		pizza = np.zeros([rows, columns])
		for row in range(rows):
			for ing,col in zip(f.readline(),range(columns)):
				if ing == 'T':
					pizza[row,col] = 1
				else:
					pizza[row,col] = 0

	return pizza, min_ingr,max_cells


class Pizza():

	def __init__(self,pizza):
		self.pizza = pizza
		
class Tree():

	def __init__(self):
		pizza, L, H = read_file('example.txt')
		self.L = L
		self.H = H
		self.pizza = Pizza(pizza)

	def sliceDimensions(n, R, C) :
	    list = [] 
	    for i in range(1, int(math.sqrt(n) + 1)) :
	        if (n % i == 0) :
	            if (i <= R and int(n / i) <= C):
	                list.append((i, int(n / i)))
	                list.append((int(n / i), i))
	    return list


def main():
	tree = Tree()


if __name__ == '__main__':
	main()
