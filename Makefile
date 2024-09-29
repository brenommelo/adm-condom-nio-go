build:
	@go build -o bin/condominio.exe ./cmd

run: build
	@CompileDaemon --build="go build -o bin/condominio.exe ./cmd" --command="./bin/condominio.exe"

test:
	@go test -v ./...