TARGET := sudoku
SOURCES := $(shell find . -name "*.go")

all: build


.PHONY: build
build: $(TARGET)


$(TARGET): $(SOURCES)
	go build -o $@ ./cmd/$@


.PHONY: test
test:
	go get -t ./...
	go test -v .


.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out
