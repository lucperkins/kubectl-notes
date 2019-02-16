fmt:
	gofmt -w .

tidy:
	go mod tidy

imports:
	goimports -w .

spruce-up: fmt tidy imports

release:
	goreleaser release --rm-dist

test-release:
	goreleaser --snapshot --skip-publish --rm-dist
