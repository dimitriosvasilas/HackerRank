import numpy as np
import collections

RidesInfo = [] 
scores = []
RidesMap = {}

# Read the input file
def read_file(filename):
	with open(filename,'r') as f:
		line = f.readline()
		rows, cols, vehicles, rides, bonus, simsteps = [int(n) for n in line.split()]
		city = np.zeros([rows,cols])
		index = 0
		for ride in range(rides):
			line = f.readline()
			fromX, fromY, toX, toY, earliest_start , latest_finish = [int(n) for n in line.split()]
			RideInfo = [(fromX,fromY),(toX,toY),earliest_start,latest_finish]
			RidesInfo.append(RideInfo)
			rideScore = abs(RideInfo[0][0] - 0) + abs(RideInfo[0][1] - 0) + abs(RideInfo[1][0] - RideInfo[0][0]) + abs(RideInfo[1][1] - RideInfo[0][1]) + RideInfo[2]
			scores.append(rideScore)
			RidesMap[index] = rideScore
			#sorted(RidesMap.iteritems(), key=lambda key_value: key_value[0])
			index += 1
	return vehicles

def write_file(schedule):
	with open('output.txt','w') as f:
		for key, value in schedule.items():
			f.write(str(len(value)))
			for r in value:
				f.write(' ')
				f.write(str(r))
			f.write('\n')
		#for i in range(len(maxList)):
			#f.write(str(maxList[i][0]) +' '+str(maxList[i][1])+' '+str(maxList[i][2])+' '+ str(maxList[i][3])+'\n')


def main(): 
	vehicle_num = read_file('c_no_hurry.in')
	schedule = {}
	for i in range(vehicle_num):
		schedule[i] = []
		
	#print(RidesMap)
	vehiclesScore = [0]*vehicle_num
	vehiclesPos = [(0,0)]*vehicle_num
	for i in range(len(RidesInfo)):
		maxR = max(scores)
		maxRi = scores.index(maxR)
		scores[maxRi] = -1
		#print("max ride:", maxRi)
		#print("max ride score:", maxR)
		
		
		minV = min(vehiclesScore)
		minVi = vehiclesScore.index(minV)
		#print("min vehicle:", minVi)
		#print("min vehicle score:", minV)
		
		vehiclesScore[minVi] += maxR
		vehiclesPos[minVi] = (RidesInfo[maxRi][1][0], RidesInfo[maxRi][1][1])
		
		#cp c(schedule)
		schedule[minVi].append(maxRi)
		#print(schedule)
		write_file(schedule)
	

if __name__ == '__main__':
	main()