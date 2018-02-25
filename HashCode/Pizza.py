import numpy as np


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

	def __init__(self,pizza,L,H):
		self.pizza = pizza
		self.L = L
		self.H = H



def main():
	pizza, L, H = read_file('example.txt')
	NewPizza = Pizza(pizza,L,H)




if __name__ == '__main__':
	main()
