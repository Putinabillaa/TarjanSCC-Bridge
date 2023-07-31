package main

/* Directed Graph for SCC */

type DirGraph struct {
	Vertex  []string
	Adj     map[string][]string
	Index   int
	Disc    map[string]int
	LowLink map[string]int
	OnStack map[string]bool
	Stack   []string
	SCCs    [][]string
}

func NewDirGraph() *DirGraph {
	return &DirGraph{
		Vertex:  []string{},
		Adj:     make(map[string][]string),
		Index:   0,
		Disc:    make(map[string]int),
		LowLink: make(map[string]int),
		OnStack: make(map[string]bool),
		Stack:   []string{},
		SCCs:    [][]string{},
	}
}

func (g *DirGraph) AddEdge(u, v string) {
	g.Adj[u] = append(g.Adj[u], v)
	if !contains(g.Vertex, u) {
		g.Vertex = append(g.Vertex, u)
	}
	if !contains(g.Vertex, v) {
		g.Vertex = append(g.Vertex, v)
	}
}

/* Undirected Graph for Bridge */

type UnDirGraph struct {
	Vertex  []string
	Adj     map[string][]string
	Index   int
	Disc    map[string]int
	LowLink map[string]int
	Visited map[string]bool
	Parent  map[string]string
	Bridges [][2]string
}

func NewUnDirGraph() *UnDirGraph {
	return &UnDirGraph{
		Vertex:  []string{},
		Adj:     make(map[string][]string),
		Index:   0,
		Disc:    make(map[string]int),
		LowLink: make(map[string]int),
		Visited: make(map[string]bool),
		Parent:  make(map[string]string),
		Bridges: [][2]string{},
	}
}

func (g *UnDirGraph) AddEdge(u, v string) {
	if !contains(g.Adj[u], v) {
		g.Adj[u] = append(g.Adj[u], v)
	}
	if !contains(g.Adj[v], u) {
		g.Adj[v] = append(g.Adj[v], u)
	}

	if !contains(g.Vertex, u) {
		g.Vertex = append(g.Vertex, u)
	}
	if !contains(g.Vertex, v) {
		g.Vertex = append(g.Vertex, v)
	}
}
