// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sccandbridge", handleSCCAndBridge)

	c := cors.Default().Handler(r)
	http.Handle("/", c)

	fs := http.FileServer(http.Dir("static/images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func handleSCCAndBridge(w http.ResponseWriter, r *http.Request) {
	returnError := func(message string, statusCode int) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	}
	if r.Method != http.MethodPost {
		returnError("Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	type InputData struct {
		Input string `json:"input"`
	}

	var inputData InputData
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		returnError("Error decoding request body", http.StatusInternalServerError)
		return
	}

	var gDir *DirGraph
	var gUnDir *UnDirGraph
	var execTime time.Duration

	lines := strings.Split(inputData.Input, "\n")
	gDir = NewDirGraph()
	gUnDir = NewUnDirGraph()

	for _, line := range lines {
		edge := strings.Fields(line)
		if len(edge) != 2 {
			returnError("Error input text format", http.StatusInternalServerError)
			return
		}
		gDir.AddEdge(edge[0], edge[1])
		gUnDir.AddEdge(edge[0], edge[1])
	}

	startTime := time.Now()
	gDir.FindSCC()
	gUnDir.FindBridge()
	execTime = time.Since(startTime)

	if err := visualizeGraph(gDir); err != nil {
		returnError(fmt.Sprintf("Error visualizing graph: %s", err), http.StatusInternalServerError)
		return
	}

	if err := visualizeSCC(gDir); err != nil {
		returnError(fmt.Sprintf("Error visualizing graph: %s", err), http.StatusInternalServerError)
		return
	}

	if err := visualizeBridge(gUnDir); err != nil {
		returnError(fmt.Sprintf("Error visualizing graph: %s", err), http.StatusInternalServerError)
		return
	}
	outputData := struct {
		SCCs     [][]string
		Bridges  [][2]string
		ExecTime string
		ImageSrc []string
	}{gDir.SCCs, gUnDir.Bridges, execTime.String(),
		[]string{
			"http://localhost:8080/images/graph.png",
			"http://localhost:8080/images/scc.png",
			"http://localhost:8080/images/bridge.png",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputData)
	if err != nil {
		returnError("Error encoding response body", http.StatusInternalServerError)
		return
	}
}
