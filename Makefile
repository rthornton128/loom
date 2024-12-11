.PHONY: cmd
all: app cmd

.PHONY: app
app:
	go build -v -o server ./app

.PHONY: cmd
cmd:
	go build -v -o bin/ ./cmd/...

.PHONY: test
test:
	go test -v ./...
