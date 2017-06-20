import sys

class Graph(object):
	def __init__(self):
		self.totalSum = None
		self.nodes = []

class Node(object):
	def __init__(self):
		self.id = None
		self.dist = -1
		self.neighbors = []
		self.visited = False

def bfs_distance(graph, root):
	queue = []
	root = graph.nodes[root-1]
	root.dist = 0
	queue.append(root)
	root.visited = True
	while len(queue) > 0:
		node = queue.pop(0)
		for n in node.neighbors:
			if n.visited == False:
				n.dist = node.dist+6
				queue.append(n)
				n.visited = True

q = int(raw_input())
for i in range(q):
	n,m = raw_input().strip().split(' ')
	n,m = [int(n),int(m)]
	graph = Graph()
	for j in range(n):
		node = Node()
		node.id = j+1
		graph.nodes.append(node)
	for j in range(m):
		u, v = raw_input().strip().split(' ')
		u, v = [int(u),int(v)]
		graph.nodes[u-1].neighbors.append(graph.nodes[v-1])
		graph.nodes[v-1].neighbors.append(graph.nodes[u-1])
	
	s = int(raw_input())
	
	bfs_distance(graph, s)

	for node in graph.nodes:
		if node.dist != 0:
			sys.stdout.write(str(node.dist) + ' ')
	sys.stdout.write('\n')
