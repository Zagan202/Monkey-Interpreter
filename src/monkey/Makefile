# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=monkey

all: run
build: 
		$(GOBUILD) $(BINARY_NAME)
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) $(BINARY_NAME)
		./$(BINARY_NAME)
