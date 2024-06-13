all:goweb

goweb:./cmd/main.go
	go build -o goweb ./cmd/main.go
