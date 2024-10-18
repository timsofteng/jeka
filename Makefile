.DELETE_ON_ERROR:

.PHONY:

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

codegen-telegram-http-client:
	oapi-codegen \
		-config ./api/telegram-http-client/cfg.yaml \
		https://raw.githubusercontent.com/PaulSonOfLars/telegram-bot-api-spec/refs/heads/main/api.json


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

