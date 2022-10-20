ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  DIRECTION := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(DIRECTION):;@:)
endif

.PHONY: migrate 
migrate:
	@go run cmd/migrations/main.go -migrate $(DIRECTION)