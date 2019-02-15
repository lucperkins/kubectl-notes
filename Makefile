fmt:
	gofmt -w .

tidy:
	go mod tidy

imports:
	goimports -w .

spruce-up: fmt tidy imports
