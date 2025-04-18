default: generate build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o livredger

generate:
	templ generate
	find ./internal/presentation \( ! -path './internal/presentation/views/*' -a -name '*_templ.go' \) -exec mv {} internal/presentation/views \;

run:
	air .
