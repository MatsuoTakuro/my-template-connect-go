build: ## TBU
	docker compose build

up: ## TBU
	docker compose up -d

up_only_app: ## Confirm error logs in stout when db does not run
	docker compose up app -d

up_only_db:
	docker compose up db -d

down: ## TBU
	docker compose down

db_cleanup: ## TBU
	psql -h 127.0.0.1 -p 5432 -d template-db -U postgres -W -f ./testutils/data/cleanup_stores.sql

db_setup: ## TBU
	psql -h 127.0.0.1 -p 5432 -d template-db -U postgres -W -f ./testutils/data/create_stores.sql -f ./testutils/data/insert_stores.sql

db_in: ## TBU
	psql -h 127.0.0.1 -p 5432 -d template-db -U postgres -W

test: ## Execute tests
  ## go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
	go test -race -shuffle=on ./...

call_Greet: ## Call Greet grpc request
	grpcurl -plaintext -proto ./proto/greet/v1/greet.proto -d '{"name": "test" }' localhost:9090 greet.v1.GreetService/Greet

call_StoreList: ## Call StoreList grpc request
	grpcurl -plaintext -proto ./proto/templateconnectgo/v1/store.proto  -d '{"search_query": "田", "company_cd": 1}' localhost:9090 templateconnectgo.v1.StoreService/ListStores


help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
