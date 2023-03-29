.PHONY: package

build-static:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vhupoc server/server.go
