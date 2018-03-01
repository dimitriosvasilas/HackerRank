import numpy as np

RidesInfo = [] 

# Read the input file
def read_file(filename):
	with open(filename,'r') as f:
		line = f.readline()
		rows, cols, vehicles, rides, bonus, simsteps = [int(n) for n in line.split()]
		city = np.zeros([rows,cols])
		for ride in range(rides):
			fromX, fromY, toX, toY, earliest_start , latest_finish = [int(n) for n in line.split()]
			RideInfo = [(fromX,fromY),(toX,toY),earliest_start,latest_finish]
			RidesInfo.append(RideInfo)


def main(): 
	read_file('a_example.in')

if __name__ == '__main__':
	main()