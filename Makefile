ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: run
run:
	@echo "Running application..."
	go run . $(RUN_ARGS) -tags sdl

.PHONY: update
update:
	@echo "Updating Carcassonne-Engine..."
	go get -u github.com/YetAnotherSpieskowcy/Carcassonne-Engine@main
	@echo "Updating raylib-go..."
	go get -v -u github.com/gen2brain/raylib-go/raylib
