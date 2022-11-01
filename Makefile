ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  DIRECTION := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(DIRECTION):;@:)
endif

.PHONY: migrate run
migrate:
	@go run cmd/migrations/main.go -migrate $(DIRECTION)

run:
	@go run cmd/rest/main.go