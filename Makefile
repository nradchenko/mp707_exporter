NAME = mp707_exporter

CGO_ENABLED = 1
GO111MODULE = on

export CGO_ENABLED
export GO111MODULE

GIT_COMMIT = $(shell git describe --always --dirty)

LDFLAGS += -w -s -X main.Version=$(GIT_COMMIT)

.PHONY: all
all: build

.PHONY: check
check: vet test

.PHONY: vet
vet:
	go vet -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -vf $(NAME)

.PHONY: build
build: $(NAME)

$(NAME):
	go build -v -ldflags "$(LDFLAGS)"
