package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/awalterschulze/gographviz"
)

/* Visualize graph */

func visualizeGraph(g *DirGraph) error {
	graphAst, err := gographviz.ParseString(`digraph G {bgcolor="#0E1116";}`)
	if err != nil {
		return fmt.Errorf("error parsing graph: %w", err)
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.Vertex); i++ {
		nodeStr := g.Vertex[i]
		attrs := make(map[string]string)
		attrs["fontcolor"] = "white"
		attrs["color"] = "white"
		graph.AddNode("G", nodeStr, attrs)
		for _, neighbor := range g.Adj[nodeStr] {
			attrs := make(map[string]string)
			attrs["color"] = "white"
			graph.AddEdge(nodeStr, neighbor, true, attrs)
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return fmt.Errorf("error analyzing graph: %w", err)
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		return fmt.Errorf("error generating dot file: %w", err)
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		return fmt.Errorf("error writing to dot file: %w", err)
	}

	outputFile := "static/images/graph.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating graph: %w", err)
	}

	return nil
}

/* Visualize bridge */

func visualizeBridge(g *UnDirGraph) error {
	graphAst, err := gographviz.ParseString(`digraph G {bgcolor="#0E1116";}`)
	if err != nil {
		return fmt.Errorf("error parsing graph: %w", err)
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.Vertex); i++ {
		nodeStr := g.Vertex[i]
		attrs := make(map[string]string)
		attrs["fontcolor"] = "white"
		attrs["color"] = "white"
		graph.AddNode("G", nodeStr, attrs)
		for _, neighbor := range g.Adj[nodeStr] {
			isBridge := false
			for _, bridge := range g.Bridges {
				if (bridge[0] == nodeStr && bridge[1] == neighbor) || (bridge[0] == neighbor && bridge[1] == nodeStr) {
					isBridge = true
					break
				}
			}

			attrs := make(map[string]string)
			if isBridge {
				attrs["color"] = "blue"
			} else {
				attrs["color"] = "darkgrey"
			}
			graph.AddEdge(nodeStr, neighbor, true, attrs)
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return fmt.Errorf("error analysing graph: %w", err)
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		return fmt.Errorf("error generating dot file: %w", err)
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		return fmt.Errorf("error writing to dot file: %w", err)
	}

	outputFile := "static/images/bridge.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating graph: %w", err)
	}

	return nil
}

func visualizeSCC(g *DirGraph) error {
	graphAst, err := gographviz.ParseString(`digraph G {bgcolor="#0E1116";}`)
	if err != nil {
		return fmt.Errorf("error parsing graph: %w", err)
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.SCCs); i++ {
		for j := 0; j < len(g.SCCs[i]); j++ {
			nodeStr := g.SCCs[i][j]
			attrs := make(map[string]string)
			attrs["fontcolor"] = "white"
			attrs["color"] = "white"
			graph.AddNode("G", nodeStr, attrs)
			for _, neighbor := range g.Adj[nodeStr] {
				if contains(g.SCCs[i], neighbor) {
					attrs := make(map[string]string)
					attrs["color"] = "white"
					graph.AddEdge(nodeStr, neighbor, true, attrs)
				}
			}
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return fmt.Errorf("error analysing graph: %w", err)
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		return fmt.Errorf("error generating dot file: %w", err)
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		return fmt.Errorf("error writing to dot file: %w", err)
	}

	outputFile := "static/images/scc.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error generating graph: %w", err)
	}

	return nil
}
