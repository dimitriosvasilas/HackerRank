n = int(raw_input())
m = int(raw_input())
grid = []
for i in range(n):
	row = map(int,raw_input().strip().split(' '))
	grid.append(row)

discovered = []
for i in range(n):
	discovered.append([0]*m)

def dfs_connected_component(posx, posy):
	if discovered[posx][posy] == 1:
		return 0
	discovered[posx][posy] = 1
	if grid[posx][posy] == 0:
		return 0
	else:
		count = 1
		if posx-1 >= 0 and posy-1 >= 0 and grid[posx-1][posy-1] == 1:
			count += dfs_connected_component(posx-1, posy-1)
		if posx-1 >= 0 and grid[posx-1][posy] == 1:
			count += dfs_connected_component(posx-1, posy)
		if posx-1 >= 0 and posy+1 < m and grid[posx-1][posy+1] == 1:
			count += dfs_connected_component(posx-1, posy+1)
		if posy-1 >= 0 and grid[posx][posy-1] == 1:
			count += dfs_connected_component(posx, posy-1)
		if posy+1 < m and grid[posx][posy+1] == 1:
			count += dfs_connected_component(posx, posy+1)
		if posx+1 < n and posy-1 >= 0 and grid[posx+1][posy-1] == 1:
			count += dfs_connected_component(posx+1, posy-1)
		if posx+1 < n and grid[posx+1][posy] == 1:
			count += dfs_connected_component(posx+1, posy)
		if posx+1 < n and posy+1 < m and grid[posx+1][posy+1] == 1:
			count += dfs_connected_component(posx+1, posy+1)
		return count

max_count=0
for i in range(n):
	for j in range(m):
		count = dfs_connected_component(i,j)
		if count > max_count:
			max_count = count

print max_count
