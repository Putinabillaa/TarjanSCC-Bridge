package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	var filepath string
	var choose int

	color.New(color.FgMagenta, color.Bold).Println("Welcome to SCC and Bridge Finder!")
	color.New(color.FgCyan, color.Bold).Println("Choose input method:")
	fmt.Println("1. From file")
	fmt.Println("2. From terminal")
	color.New(color.FgGreen).Print(">>> ")
	fmt.Scanln(&choose)

	var gDir *DirGraph
	var gUnDir *UnDirGraph
	var err error

	for choose != 1 && choose != 2 {
		color.New(color.FgRed).Println("Invalid input. Please choose again.")
		color.New(color.FgGreen).Print(">>> ")
		fmt.Scanln(&choose)
	}

	if choose == 1 {
		color.New(color.FgCyan, color.Bold).Println("Enter filepath:")
		color.New(color.FgGreen).Print(">>> ")
		fmt.Scanln(&filepath)
		gDir, gUnDir, err = readFile(filepath)
		for err != nil {
			color.New(color.FgRed).Println("Error:", err)
			color.New(color.FgCyan, color.Bold).Println("Enter filepath:")
			color.New(color.FgGreen).Print(">>> ")
			fmt.Scanln(&filepath)
			gDir, gUnDir, err = readFile(filepath)
		}
	} else if choose == 2 {
		var E int

		gDir = NewDirGraph()
		gUnDir = NewUnDirGraph()

		for {
			color.New(color.FgCyan, color.Bold).Println("Enter number of edges:")
			color.New(color.FgGreen).Print(">>> ")
			_, err := fmt.Scanln(&E)
			if err == nil && E >= 0 {
				break
			}
			color.New(color.FgRed).Println("Please enter a valid non-negative integer.")
		}

		for i := 0; i < E; i++ {
			color.New(color.FgCyan, color.Bold).Println("Enter edge (format: A B):")
			color.New(color.FgGreen).Print(">>> ")
			var from, to string
			fmt.Scanf("%s %s", &from, &to)
			from = strings.TrimSpace(from)
			to = strings.TrimSpace(to)
			if len(from) != 1 || len(to) != 1 {
				color.New(color.FgRed).Println("Error: each line in the input must have exactly two values (from and to)")
				i--
				continue
			}
			gDir.AddEdge(string(from[0]), string(to[0]))
			gUnDir.AddEdge(string(from[0]), string(to[0]))
		}
	}

	startTime := time.Now()
	gDir.FindSCC()
	gUnDir.FindBridge()
	execTime := time.Since(startTime)

	visualizeBridge(gUnDir)
	visualizeGraph(gDir)
	visualizeSCC(gDir)

	color.New(color.FgMagenta, color.Bold).Print("Strongly connected components:")
	color.New(color.Bold).Println(gDir.SCCs)
	color.New(color.FgMagenta, color.Bold).Print("Bridges:")
	color.New(color.Bold).Println(gUnDir.Bridges)
	color.New(color.FgGreen).Println("Execution time:", execTime)
}
