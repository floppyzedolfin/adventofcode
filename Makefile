init:
	go mod vendor

test:
	go test ./...

cover:
	go test --cover ./...

clean:
	go clean -testcache -modcache -cache

build:
	go build -o adventofcode.out cmd/main.go

run: build
run:
	if [ -z "$(DOOR)" ]; D="-door $(DOOR)"; fi
	if [ -z "$(PARTS)" ]; P="-parts $(PARTS)"; fi
	@echo "Running Advent of Code for December $(filter-out $@,$(MAKECMDGOALS))"
	./adventofcode.out $(D) $(P)
