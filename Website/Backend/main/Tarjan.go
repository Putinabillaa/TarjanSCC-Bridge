/* Tarjan Algorithm for finding SCC and Bridge */
package main

/*
FindSCC finds all strongly connected components in a
directed graph using Tarjan's algorithm
*/

func (g *DirGraph) FindSCC() {
	g.Stack = nil
	for _, v := range g.Vertex {
		if g.Disc[v] == 0 {
			g.TarjanSCC(v)
		}
	}
}

func (g *DirGraph) TarjanSCC(v string) {
	g.Index++
	g.Disc[v] = g.Index
	g.LowLink[v] = g.Index
	g.Stack = append(g.Stack, v)
	g.OnStack[v] = true

	for _, w := range g.Adj[v] {
		if g.Disc[w] == 0 {
			g.TarjanSCC(w)
			g.LowLink[v] = min(g.LowLink[v], g.LowLink[w])
		} else if g.OnStack[w] {
			g.LowLink[v] = min(g.LowLink[v], g.Disc[w])
		}
	}
	if g.LowLink[v] == g.Disc[v] {
		var scc []string
		for {
			w := g.Stack[len(g.Stack)-1]
			g.Stack = g.Stack[:len(g.Stack)-1]
			g.OnStack[w] = false
			scc = append(scc, w)
			if w == v {
				break
			}
		}
		g.SCCs = append(g.SCCs, scc)
	}
}

/*
FindBridge finds all bridges in an undirected
graph using Tarjan's algorithm
*/

func (g *UnDirGraph) FindBridge() {
	for _, v := range g.Vertex {
		if !g.Visited[v] {
			g.TarjanBridge(v, "")
		}
	}
}

func (g *UnDirGraph) TarjanBridge(v, parent string) {
	g.Visited[v] = true
	g.Index++
	g.Disc[v] = g.Index
	g.LowLink[v] = g.Index
	g.Parent[v] = parent

	for _, w := range g.Adj[v] {
		if !g.Visited[w] {
			g.TarjanBridge(w, v)
			g.LowLink[v] = min(g.LowLink[v], g.LowLink[w])
			if g.LowLink[w] > g.Disc[v] {
				g.Bridges = append(g.Bridges, [2]string{v, w})
			}
		} else if w != parent {
			g.LowLink[v] = min(g.LowLink[v], g.Disc[w])
		}
	}
}
