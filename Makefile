run:
	@go build
	@./atton parse

fmt:
	go fmt ./...
