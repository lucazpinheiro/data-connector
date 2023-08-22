# Default go compiler
GC := go

# Entry point file
ENTRY := main.go

# Output file
OUTPUT := bin/app

# Makefile flags
MAKEFLAGS += "--silent"

# Script to generate binary
# output file.
build:
	@echo "Generating binary file..."
	@cd app && $(GC) build -o ../$(OUTPUT)
	@echo "Output file was succesfully generated."

run-dev:
	@echo "Running app as dev..."
	@go run app/*.go

run:
	@echo "Running app"
	@./bin/app
