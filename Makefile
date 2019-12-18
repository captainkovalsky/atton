make:
	@echo "Default make"

clean:
	rm -rf dist
	mkdir -p dist

run:
	@./dist/atton parse

fmt:
	go fmt ./...

build: clean win macos

macos:
	@go build -o publish/atton main.go

win:
	@GOOS=windows GOARCH=386 go build -o publish/atton.exe main.go
