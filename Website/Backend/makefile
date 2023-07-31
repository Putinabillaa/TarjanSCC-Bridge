# Variables
GOCMD=go
MAINFILE=main/main.go
GRAPHFILE=main/Graph.go
VISUALIZERFILE=main/Visualizer.go
TARJANFILE=main/Tarjan.go
HELPERFILE=main/Helper.go

# Targets
.PHONY: build run clean

build:
	$(GOCMD) build -o server $(MAINFILE) $(GRAPHFILE) $(VISUALIZERFILE) $(TARJANFILE) $(HELPERFILE)

run:
	$(GOCMD) run $(MAINFILE) $(GRAPHFILE) $(VISUALIZERFILE) $(TARJANFILE) $(HELPERFILE)

clean:
	rm -f server
