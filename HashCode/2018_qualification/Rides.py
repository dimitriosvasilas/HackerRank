import numpy as np
import collections
from time import sleep
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
		scoresArr = np.zeros([vehicles,rides])
		for ride in range(rides):
			line = f.readline()
			fromX, fromY, toX, toY, earliest_start , latest_finish = [int(n) for n in line.split()]
			RideInfo = [(fromX,fromY),(toX,toY),earliest_start,latest_finish]
			RidesInfo.append(RideInfo)
			rideScore = abs(RideInfo[0][0] - 0) + abs(RideInfo[0][1] - 0) + abs(RideInfo[1][0] - RideInfo[0][0]) + abs(RideInfo[1][1] - RideInfo[0][1]) + RideInfo[2]
			scores.append(rideScore)
			RidesMap[index] = rideScore
			for i in range(vehicles):
				scoresArr[i][ride] = rideScore
			index += 1
	return vehicles, scoresArr

def write_file(schedule):
	with open('output.txt','w') as f:
		for key, value in schedule.items():
			f.write(str(key+1))
			for r in value:
				f.write(' ')
				f.write(str(r))
			f.write('\n')
		#for i in range(len(maxList)):
			#f.write(str(maxList[i][0]) +' '+str(maxList[i][1])+' '+str(maxList[i][2])+' '+ str(maxList[i][3])+'\n')


def main():
	vehicle_num, scoresArr = read_file('a_example.in')
	schedule = {}
	for i in range(vehicle_num):
		schedule[i] = []
		
	print(RidesMap)
	print(scoresArr)
	vehiclesScore = [0]*vehicle_num
	vehiclesPos = [(0,0)]*vehicle_num
	for i in range(len(RidesInfo)):
		
		minV = min(vehiclesScore)
		minVi = vehiclesScore.index(minV)
		print("vehicle:", minVi)

		while True:
			maxScore = np.amax(scoresArr)
			maxI,maxJ = np.unravel_index(scoresArr.argmax(), scoresArr.shape)
			print(scoresArr)
			print("vehiclesScore:", vehiclesScore[minVi])
			print("maxScore:", maxScore)
			print("x, y", maxI,maxJ)
			print("RidesInfo:", RidesInfo[maxI][3])
			if (vehiclesScore[minVi] + maxScore <= RidesInfo[maxI][3]):
				break
			else:
				scoresArr[minVi][maxJ] = -1
		print("-------------maxScore:", maxScore)
		print("-------------i:", maxI)
		print("-------------ride:", maxJ)
		
		vehiclesScore[minVi] += maxScore
		print("vehiclesScore", vehiclesScore)
	
		vehiclesPos[minVi] = (RidesInfo[maxJ][1][0], RidesInfo[maxJ][1][1])
		print("vehiclesPos", vehiclesPos)
		
		for i in range(vehicle_num):
			scoresArr[i][maxJ] = -1
		
		schedule[minVi].append(maxJ)
		print(schedule)	
		
		print("updating")
		for j in range(len(RidesInfo)):
				if scoresArr[minVi][j] == -1:
					continue
				scoresArr[minVi][j] += abs(RidesInfo[j][0][0] - vehiclesPos[minVi][0]) 
				scoresArr[minVi][j] += abs(RidesInfo[j][0][1] - vehiclesPos[minVi][1]) 
				scoresArr[minVi][j] += abs(RidesInfo[j][1][0] - RidesInfo[j][0][0])
				scoresArr[minVi][j] += abs(RidesInfo[j][1][1] - RidesInfo[j][0][1])
				scoresArr[minVi][j] -= abs(RidesInfo[j][0][0])
				scoresArr[minVi][j] -= abs(RidesInfo[j][0][1])
		sleep(1)
		print(scoresArr)
		sleep(1)
	write_file(schedule)

if __name__ == '__main__':
	main()