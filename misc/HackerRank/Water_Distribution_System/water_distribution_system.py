class Graph(object):
    def __init__(self):
        self.totalSum = None
        self.nodes = []

class Node(object):
    def __init__(self):
        self.id = None
        self.neighbors = []
        self.visited = False
        self.color = 0

def dfs(root, came_from):
    root.color = 1
    #print "dfs", root.id, "color", root.color
    for i in reversed(xrange(len(root.neighbors))):
        #print "examine", root.neighbors[i].id, "color", root.neighbors[i].color
        if root.neighbors[i].id == root.id:
            return True
        if root.neighbors[i].color == 1 and came_from != None and root.neighbors[i].id != came_from.id:
            return True
        if root.neighbors[i].color == 0 and dfs(root.neighbors[i], root):
            return True
    root.color = 2
    return False

def is_cyclic(graph):
  for node in graph.nodes:
      if (node.color == 0):
          if (dfs(node, None)):
              return True
  return False

t = int(raw_input())
for i in xrange(t):
    n, m = map(int,raw_input().split())
    e = map(int, raw_input().split(' '))
    
    graph = Graph()
    
    for i in range(n):
        node = Node()
        node.id = i
        graph.nodes.append(node)
    
    i = 0
    while (i < 2*m-1):
        graph.nodes[e[i]].neighbors.append(graph.nodes[e[i+1]])
        graph.nodes[e[i+1]].neighbors.append(graph.nodes[e[i]])
        i += 2
    
    if(is_cyclic(graph)):
        print "1"
    else:
        print "0"

