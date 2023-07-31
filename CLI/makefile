# Variables
GOCMD=go
MAINFILE=main/main.go
GRAPHFILE=main/Graph.go
VISUALIZERFILE=main/Visualizer.go
TARJANFILE=main/Tarjan.go
FILEREADERFILE=main/FileReader.go
HELPERFILE=main/Helper.go

# Targets
.PHONY: build run clean

build:
	mkdir -p output
	$(GOCMD) build -o cli $(MAINFILE) $(GRAPHFILE) $(VISUALIZERFILE) $(TARJANFILE) $(FILEREADERFILE) $(HELPERFILE) 

run:
	mkdir -p output
	$(GOCMD) run $(MAINFILE) $(GRAPHFILE) $(VISUALIZERFILE) $(TARJANFILE) $(FILEREADERFILE) $(HELPERFILE) 

clean:
	rm -f cli
