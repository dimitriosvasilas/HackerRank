class Graph(object):
	def __init__(self):
		self.totalSum = None
		self.nodes = []

class Node(object):
	def __init__(self):
		self.id = None
		self.neighbors = []
		self.visited = False


def dfs_connected_comp(root):
	stack = []
        stackid = []
        count = 0
        stack.append(node)
        stackid.append(node.id)
        while len(stack) > 0:
                n = stack.pop(0)
                stackid.pop(0)
                if n.visited == False:
                        n.visited = True
                        count += 1
                        for i in reversed(xrange(len(n.neighbors))):
                                if n.neighbors[i].visited == False:
                                        stack.insert(0, n.neighbors[i])
                                        stackid.insert(0, n.neighbors[i].id)
	return count

N,l = map(int,raw_input().split())
graph = Graph()
res = N*(N-1)/2

for i in range(N):
	node = Node()
	node.id = i
	graph.nodes.append(node)

for i in range(l):
	a,b = map(int,raw_input().split())
	graph.nodes[a].neighbors.append(graph.nodes[b])
	graph.nodes[b].neighbors.append(graph.nodes[a])

components = []
for node in graph.nodes:
	count = dfs_connected_comp(node)
	res -= count*(count-1)/2

print res
