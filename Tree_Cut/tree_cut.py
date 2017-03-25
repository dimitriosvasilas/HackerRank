import sys

class Graph(object):
	def __init__(self):
		self.totalSum = None
		self.nodes = []

class Node(object):
	def __init__(self):
		self.neighbors = []
		self.data = None
		self.id = None
		self.visited = False

N = int(raw_input())
data = map(int,raw_input().strip().split())
graph = Graph()

for i in range(N):
	node = Node()
	node.data = data[i]
	node.id = i
	graph.nodes.append(node)

for i in range(N-1):
	edge = map(int,raw_input().strip().split())
	graph.nodes[edge[0]-1].neighbors.append(graph.nodes[edge[1]-1])
	graph.nodes[edge[1]-1].neighbors.append(graph.nodes[edge[0]-1])


def dfs_sum(node):
	node.visited = True
	for n in node.neighbors:
		if n.visited == False:
			node.data += dfs_sum(n)
	return node.data


graph.totalSum = dfs_sum(graph.nodes[0])

mind = graph.totalSum
for n in graph.nodes:
	d = abs(graph.totalSum - 2*n.data)
	if d < mind:
		mind = d

print mind
