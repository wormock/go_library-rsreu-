.PHONY: test
test:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html

.PHONY: deps
deps:
	go mod tidy