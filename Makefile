build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o handlers/utakata.handler src/utakata/utakata.go
