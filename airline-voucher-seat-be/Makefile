
# Load .env file jika ada
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# Commands
run:
	@go run cmd/main/main.go