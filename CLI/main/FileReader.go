package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

/* read input from file */

func readFile(filepath string) (*DirGraph, *UnDirGraph, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var edges [][]string
	var maxCurrEdges int
	maxEdges := math.MinInt64
	for scanner.Scan() {
		line := scanner.Text()
		edge := strings.Split(line, " ")
		if len(edge) != 2 {
			return nil, nil, fmt.Errorf("each line in the file must have exactly two values (from and to)")
		}
		maxCurrEdges = int(math.Max(float64(edge[0][0]), float64(edge[1][0])))
		maxEdges = int(math.Max(float64(maxEdges), float64(maxCurrEdges)))
		edges = append(edges, edge)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	gDir := NewDirGraph()
	gUnDir := NewUnDirGraph()
	for _, edge := range edges {
		from := string(edge[0][0])
		to := string(edge[1][0])
		gDir.AddEdge(from, to)
		gUnDir.AddEdge(from, to)
	}

	return gDir, gUnDir, nil
}
