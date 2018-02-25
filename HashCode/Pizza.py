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

	return pizza, min_ingr,max_cells, rows, columns

		
class Tree():

	def __init__(self):
		self.pizza = None
		self.L = None
		self.H = None
		self.R = None
		self.C = None
		self.children = []
		self.score = 0
		
	def from_Input(self):
		pizza, L, H, R, C = read_file('example.txt')
		self.L = L
		self.H = H
		self.R = R
		self.C = C
		self.pizza = pizza
	
	def from_State(self, state, x, y, dimX, dimY):
		self.L = state.L
		self.H = state.H
		self.R = state.R
		self.C = state.C
		self.pizza = state.pizza.copy()
		self.score = state.score
		for r in range(x, x+dimX):
			for c in range(y, y+dimY):
				self.pizza[r][c] = None
				self.score += 1
	
	def show(self):
		print(self.pizza)

	#Calulate maximum possible dimensions for a slice containg n cells
	def sliceDimensions(self,n, R, C):
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
		
		for size in range(self.H, 2*self.L-1, -1):
			flag = False
			FittingSlices = self.sliceDimensions(size, self.R, self.C)
			for i in range(len(FittingSlices)):
				for r in range(self.R):
					for c in range(self.C):
						StartingPosition = (r,c) #needs changes to be general
						if (self.CorrectSlice(StartingPosition[0],StartingPosition[1],FittingSlices[i][0],FittingSlices[i][1])):
							flag = True
							print (FittingSlices[i][0], FittingSlices[i][1], r, c, "Correct Slice")
							newChild = Tree()
							newChild.from_State(self, StartingPosition[0], StartingPosition[1], FittingSlices[i][0], FittingSlices[i][1])
							newChild.show()
							print(newChild.score)
							self.children.append(newChild)
			if flag:
				break
		for c in self.children:
			c.Check()
			


def main():
	tree = Tree()
	tree.from_Input()
	tree.Check()


if __name__ == '__main__':
	main()
