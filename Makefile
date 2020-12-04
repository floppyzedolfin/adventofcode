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

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

run: build
run:
	./adventofcode.out $(RUN_ARGS)
