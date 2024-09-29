DB_URL=postgres://postgres:123456@localhost:5432/condominio?sslmode=disable
GOOSE_DIR=./migrations
BINARY_OUTPUT=bin/condominio.exe
CMD_DIR=./cmd

build:
	@go build -o $(BINARY_OUTPUT) $(CMD_DIR)

run: build
	@CompileDaemon --build="go build -o $(BINARY_OUTPUT) $(CMD_DIR)" --command="$(BINARY_OUTPUT)"

# Target para rodar as migrações de banco de dados
migrate-up:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" up

# Target para reverter a última migração
migrate-down:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" down $(STEPS)

# Target para exibir o status das migrações
migrate-status:
	goose -dir $(GOOSE_DIR) postgres "$(DB_URL)" status

test:
	@go test -v ./...
