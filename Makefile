build: ## TBU
	docker compose build

up: ## TBU
	docker compose up -d

up_fore: ## TBU
	docker compose up

up_only_app: ## Confirm error logs in stout when db does not run
	docker compose up app -d

up_only_db: ## TBU
	docker compose up db -d

log_app: ## TBU
	docker logs template-connect-go --follow

log_db: ## TBU
	docker logs template-db --follow

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

test_v: ## Execute tests in details
	go test -v -race -shuffle=on ./...

call_greet: ## Call Greet request on grpc
	grpcurl -plaintext -proto ./proto/greet/v1/greet.proto -d '{"name": "test" }' localhost:9090 greet.v1.GreetService/Greet

call_store_list_grpc: ## Call StoreList request on grpc
	grpcurl -plaintext -proto ./proto/templateconnectgo/v1/store.proto  -d '{"search_query": "田", "company_cd": 1}' localhost:9090 templateconnectgo.v1.StoreService/ListStores

call_hello: ## Call Hello request on http
	curl localhost:8080/hello

call_store_list_http: ## Call StoreList request on http
	curl --get -d search_query=田 -d company_cd=1 localhost:8080/store

call_all: call_greet call_store_list_grpc call_hello call_store_list_http

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
