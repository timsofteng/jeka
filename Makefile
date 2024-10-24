.DELETE_ON_ERROR:

.PHONY:
build:
	CGO_ENABLED=0 go build -o bin/${APP_NAME} .

.PHONY:
app:
	./${APP_NAME}

.PHONY: e
e:
	nvim

.PHONY: dev
dev:
	air

.PHONY:
run:
		go run .

.PHONY:
test:
	go test ./...

.PHONY:
lint:
	golangci-lint run --concurrency=2

.PHONY:
schema-gen:
	pg_dump --schema-only --no-owner --no-comments --no-privileges --format=p --file=./api/schema.sql ${DATABASE_URL}

.PHONY:
queries-gen:
	cd tool && sqlc generate

.PHONY:
proto-gen:
	protoc \
		--proto_path=./api/proto \
		--go_out=./services/grpcserver/pb --go_opt=paths=source_relative \
		--go-grpc_out=./services/grpcserver/pb --go-grpc_opt=paths=source_relative \
		./api/proto/*.proto	

.PHONY:
codegen-http-server-handler:
	oapi-codegen \
		-config ./tool/openapi-server.cfg.yaml \
		 ./api/openapi.yaml

## CONTAINER #####################
.PHONY:
container-build:
	docker build \
		--pull \
		--build-arg APP_NAME=$(APP_NAME) \
		--build-arg HTTP_SERVER_PORT=$(HTTP_SERVER_PORT) \
		-t $(APP_NAME) .

.PHONY:
container-run:
	docker run \
		-e HOST=$(HOST) \
		-e HTTP_SERVER_PORT=$(HTTP_SERVER_PORT) \
		-e VAULT_API_KEY=$(VAULT_API_KEY) \
		-e VAULT_DEFAULT_LEDGER=$(VAULT_DEFAULT_LEDGER) \
		-e VAULT_DEFAULT_COLLECTION=$(VAULT_DEFAULT_COLLECTION) \
		-it -p $(HTTP_SERVER_PORT):$(HTTP_SERVER_PORT) $(APP_NAME)
#########################################

