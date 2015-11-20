TARGET := sudoku
SOURCES := $(shell find . -name "*.go")

all: build


.PHONY: build
build: $(TARGET)


$(TARGET): $(SOURCES)
	go get ./...
	go build -o $@ ./cmd/$@


.PHONY: test
test:
	go get -t ./...
	go test -v .


.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out


.PHONY: goapp_serve
goapp_serve:
	goapp serve ./cmd/appspot/app.yaml


.PHONY: goapp_deploy
goapp_deploy:
	goapp deploy -application sudoku-as-a-service ./cmd/appspot/app.yaml
