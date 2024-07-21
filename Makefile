.PHONY: help
.DEFAULT_GOAL := help


help: ## Show a list of available commands
	grep "##.*" $(MAKEFILE_LIST) | grep -v ".*MAKEFILE_LIST.*" | sed -E "s/:.*##/:/g" | column -t -s :


go-help: ## Print go help screen
	go help

compile-and-run: ## Build the application
	go build ./runner/main.go
	./main

run: ## Run the application, requires the workspace (go.work) to be setup
	go run ./runner


add-module: ## Add a new module
	@read -p "Enter the name of the new module: " name && bash ./add-module.sh $$name

