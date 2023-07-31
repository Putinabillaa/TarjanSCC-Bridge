package main

import (
	"os"
	"os/exec"

	"github.com/awalterschulze/gographviz"
	"github.com/fatih/color"
)

/* Visualize graph */

func visualizeGraph(g *DirGraph) {
	graphAst, err := gographviz.ParseString(`digraph G {}`)
	if err != nil {
		color.New(color.FgRed).Println("Error parsing graph:", err)
		return
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.Vertex); i++ {
		nodeStr := g.Vertex[i]
		graph.AddNode("G", nodeStr, nil)
		for _, neighbor := range g.Adj[nodeStr] {
			graph.AddEdge(nodeStr, neighbor, true, nil)
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		color.New(color.FgRed).Println("Error analyzing graph:", err)
		return
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		color.New(color.FgRed).Println("Error creating DOT file:", err)
		return
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		color.New(color.FgRed).Println("Error writing to DOT file:", err)
		return
	}

	outputFile := "output/graph.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		color.New(color.FgRed).Println("Error generating graph:", err)
		return
	}

	color.New(color.FgGreen).Print("Graph visualization saved as ")
	color.New(color.Bold).Println(outputFile)
}

/* Visualize bridge */

func visualizeBridge(g *UnDirGraph) {
	graphAst, err := gographviz.ParseString(`digraph G {}`)
	if err != nil {
		color.New(color.FgRed).Println("Error parsing graph:", err)
		return
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.Vertex); i++ {
		nodeStr := g.Vertex[i]
		graph.AddNode("G", nodeStr, nil)
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
			}
			graph.AddEdge(nodeStr, neighbor, true, attrs)
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		color.New(color.FgRed).Println("Error analyzing graph:", err)
		return
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		color.New(color.FgRed).Println("Error creating DOT file:", err)
		return
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		color.New(color.FgRed).Println("Error writing to DOT file:", err)
		return
	}

	outputFile := "output/bridge.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		color.New(color.FgRed).Println("Error generating graph:", err)
		return
	}

	color.New(color.FgGreen).Print("Bridge visualization saved as ")
	color.New(color.FgWhite, color.Bold).Println(outputFile)
}

func visualizeSCC(g *DirGraph) {
	graphAst, err := gographviz.ParseString(`digraph G {}`)
	if err != nil {
		color.New(color.FgRed).Println("Error parsing graph:", err)
		return
	}

	graph := gographviz.NewGraph()

	for i := 0; i < len(g.SCCs); i++ {
		for j := 0; j < len(g.SCCs[i]); j++ {
			nodeStr := g.SCCs[i][j]
			graph.AddNode("G", nodeStr, nil)
			for _, neighbor := range g.Adj[nodeStr] {
				if contains(g.SCCs[i], neighbor) {
					graph.AddEdge(nodeStr, neighbor, true, nil)
				}
			}
		}
	}

	if err := gographviz.Analyse(graphAst, graph); err != nil {
		color.New(color.FgRed).Println("Error analyzing graph:", err)
		return
	}

	dot := graph.String()

	dotFile, err := os.CreateTemp("", "graphviz-*.dot")
	if err != nil {
		color.New(color.FgRed).Println("Error creating DOT file:", err)
		return
	}
	defer os.Remove(dotFile.Name())

	_, err = dotFile.WriteString(dot)
	if err != nil {
		color.New(color.FgRed).Println("Error writing to DOT file:", err)
		return
	}

	outputFile := "output/scc.png"
	cmd := exec.Command("dot", "-Tpng", "-o", outputFile, dotFile.Name())
	err = cmd.Run()
	if err != nil {
		color.New(color.FgRed).Println("Error generating graph:", err)
		return
	}

	color.New(color.FgGreen).Print("SCC visualization saved as ")
	color.New(color.Bold).Println(outputFile)
}
