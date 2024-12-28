.PHONY: build dev

build:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/static/
	GOOS=js GOARCH=wasm go build -o web/static/app.wasm cmd/wasm/main.go
	go build -o bin/web cmd/web/main.go

dev:
	go install github.com/air-verse/air@latest
	air -c config/.air.toml
