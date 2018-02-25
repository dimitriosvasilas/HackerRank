import numpy as np
import math

# Read the input file
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

		
class Tree():

	def __init__(self):
		pizza, L, H = read_file('example.txt')
		self.L = L
		self.H = H
		self.pizza = pizza

	#Calulate maximum possible dimensions for a slice containg n cells
	def sliceDimensions(self,n, R, C) :
	    Slice = [] 
	    for i in range(1, int(math.sqrt(n) + 1)) :
	        if (n % i == 0) :
	            if (i <= R and int(n / i) <= C):
	                Slice.append((i, int(n / i)))
	                Slice.append((int(n / i), i))
	    return Slice

	# Check if slice inside pizza
	def SliceInside(self,r1,c1,dimR,dimC):
		try:
			self.pizza[r1,c1] #C heck if upper left corner inside pizza
			self.pizza[r1+dimR-1,c1+dimC-1] # Cehck of lower right corner inside pizza
			return True
		except:
			return False

	# Check if slice has L cells of each ingredient
	def SatisfyLCond(self,r1,c1,dimR,dimC):
		PizzaSlice = self.pizza[r1:r1+dimR,c1:c1+dimC]
		tomatoes = np.sum(PizzaSlice)
		mushrooms = np.size(PizzaSlice) - tomatoes

		if ((tomatoes >= self.L) and (mushrooms >= self.L)):
			return True
		else:
			return False


	# Check if slize has H cells and doesn't include other slices
	def SatisfyHCond(self,r1,c1,dimR,dimC):
		if (self.SliceInside(r1,c1,dimR,dimC)):
			PizzaSlice = self.pizza[r1:r1+dimR,c1:c1+dimC]
			if (PizzaSlice.size <= self.H and not math.isnan(PizzaSlice.sum())):
				return True
		return False



	# See if slice is inside pizza, has at least L of each ingredient and H maximum cells
	def CorrectSlice(self,r1,c1,dimR,dimC):
		# Check if it has at least L ingredients
		Cond1 = self.SatisfyLCond(r1,c1,dimR,dimC)
		# Check if it has maximum H cells and doesn't include other slices
		Cond2 = self.SatisfyHCond(r1,c1,dimR,dimC)
		if (Cond1 and Cond2):
			return True
		else:
			return False


	def Check(self):
		FittingSlices = self.sliceDimensions(6,3,5)
		for i in range(len(FittingSlices)):
			StartingPosition = (0,0) #needs changes to be general
			if (self.CorrectSlice(StartingPosition[0],StartingPosition[1],FittingSlices[i][0],FittingSlices[i][1])):
				print (FittingSlices[i][0],FittingSlices[i][1],"Correct Slice")






def main():
	tree = Tree()
	tree.Check()


if __name__ == '__main__':
	main()
