.PHONY: install
install:
	@go mod download

.PHONY: update
update:
	@go get -u ./... && go mod tidy

.PHONY: listener
listener:
	@go run cmd/listener/main.go

.PHONY: publisher
publisher:
	@go run cmd/publisher/main.go

.PHONY: docker-compose
docker-compose:
	@docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down:
	@docker-compose down
