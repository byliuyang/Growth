fmt:
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

test: fmt
	go test -v ./...
