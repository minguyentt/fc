CMD = ./.bin/fc

build:
	@echo "COMPILING AND BUILDING..."
	@go build -o $(CMD) ./cmd/runner
	@echo "Executable CMD: $(CMD)"
