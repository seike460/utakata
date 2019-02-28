build:
	env GOOS=linux go build -ldflags="-s -w" -o handlers/utakataNow cmd/utakata/main.go
	env GOOS=linux go build -ldflags="-s -w -X github.com/seike460/utakata/cmd/utakata/utakata.SlackType=Daily" -o handlers/utakataDaily cmd/utakata/main.go
