import heapq
import sys

class Graph(object):
	def __init__(self):
		self.totalSum = None
		self.nodes = []

class Node(object):
	def __init__(self):
		self.id = None
		self.neighbors = []
		self.dist = sys.maxint
		self.visited = False


n, e = raw_input().strip().split(' ')
n, e = [int(n), int(e)]

graph = Graph()
for i in range(n):
	node = Node()
	node.id = i+1
	node.visited = False
	graph.nodes.append(node)

for i in xrange(e):
	u, v, d = raw_input().strip().split(' ')
	u, v, d = [int(u), int(v), int(d)]
	graph.nodes[u-1].neighbors.append((graph.nodes[v-1],d))
	graph.nodes[v-1].neighbors.append((graph.nodes[u-1],d))

graph.nodes[0].dist = 0
node_queue = [(v.dist, v) for v in graph.nodes]
heapq.heapify(node_queue)

while len(node_queue):
	uv = heapq.heappop(node_queue)
	current = uv[1]
	if current.id == len(graph.nodes):
		break
	current.visited = True

	for nxt in current.neighbors:
		if nxt[0].visited:
			continue
		new_dist = current.dist if current.dist > nxt[1] else nxt[1]
		
		if new_dist < nxt[0].dist:
			nxt[0].dist = new_dist
			heapq.heappush(node_queue, (nxt[0].dist, nxt[0]))

if graph.nodes[len(graph.nodes)-1].dist == sys.maxint:
	print "NO PATH EXISTS"
else:
	print graph.nodes[len(graph.nodes)-1].dist
