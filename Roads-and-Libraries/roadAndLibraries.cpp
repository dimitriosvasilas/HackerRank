#include <ostream>
#include <vector>
#include <iostream>
#include <stack>

using namespace std;

class Graph
{
	public:
		Graph() {}
		int vertexCount() const;
		void addVertices(int n);
		void addEdge(int start_id, int end_id, int cost = 0);
		
		const int connectedComponentCount(int start_id, vector<bool> &visited) const;
		vector<int> connectedComponents() const;
	private:
		class Edge
		{       
			public:
				Edge(int end_id, int cost): dest_id(end_id) {}
				const int getDestId() const;
			private:
				int dest_id;
		};
		class Vertex
		{
			public:
				Vertex(int id): id(id) {}
				void addEdge(int end_id, int cost);
				const int getId() const;
				const vector<Edge> & getEdges() const;

			private:
				int id;
				vector<Edge> edges;
		};
		vector<Vertex> vertices;
};
const int Graph::Vertex::getId() const
{
	return id;
}
const int Graph::Edge::getDestId() const
{
	return dest_id;
}

const vector<typename Graph::Edge> & Graph::Vertex::getEdges() const
{
	return edges;
}

void Graph::Vertex::addEdge(int end_id, int cost)
{
	edges.push_back(Edge(end_id, cost));
}

int Graph::vertexCount() const
{
	return vertices.size();
}

void Graph::addVertices(int n)
{
	for (int i = 0; i < n; i++) {
		int id = vertexCount();
		vertices.push_back(Vertex(id));
	}
}

void Graph::addEdge(int start_id, int end_id, int cost)
{
	vertices[start_id-1].addEdge(end_id-1, cost);
	vertices[end_id-1].addEdge(start_id-1, cost);
}

const int Graph::connectedComponentCount(int start_id, vector<bool> &visited) const
{
	stack<int> s;
	int count = 0;

	s.push(start_id);
	while(!s.empty()) {
		int id = s.top();
		s.pop();
		if (!visited[id])
		{
			visited[id] = true;
			count++;
			for (const Edge e : vertices[id].getEdges()) {
				int neighbor_id = e.getDestId();
				if (!visited[neighbor_id]) {
					s.push(neighbor_id);
				}
			}
		}
	}
	
	return count;
}
vector<int> Graph::connectedComponents() const
{
	vector<int> components;
	vector<bool> visited(vertexCount(), false);
	int visited_count = 0;

	for (int i = 0; i < vertexCount(); i++) {
                int c = connectedComponentCount(i, visited);
		if (c > 0)
			components.push_back(c);
		visited_count += c;
		if (visited_count == vertexCount())
			break;
        }
	return components;
}

int main()
{
	long q, n, m, c_lib, c_road, u, v;
	cin >> q;
	for (int q_i = 0; q_i < q; q_i++) {
		cin >> n >> m >> c_lib >> c_road;
	
		Graph g;
		g.addVertices(n);
	
		for(int m_i = 0; m_i < m; m_i++) {
			cin >> u >> v;
			g.addEdge(u, v);
		}
	
		if (c_road >= c_lib)
			cout << n*c_lib << "\n";
		else {
			unsigned long long cost = 0;
			vector<int> components = g.connectedComponents();
			for (vector<int>::const_iterator i = components.begin(); i != components.end(); ++i)
				cost += (*i - 1) * c_road;
			cost += components.size() * c_lib;
			cout << cost << "\n";
		}
	}
}
