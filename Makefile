server:
	go run main.go

swag:
	export PATH=$(go env GOPATH)/bin:$PATH

docs:
	swag init