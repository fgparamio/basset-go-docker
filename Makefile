SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

UNAME := $(shell uname)
TARGET := $(shell echo $${PWD\#\#*/})

$(TARGET): $(SRC)
	@go build $(LDFLAGS) -o $(TARGET)

all: test

deps:

	go get -d -v github.com/gorilla/handlers
	go get -d -v github.com/gorilla/mux
	go get -d -v github.com/gorilla/context

updatedeps:

	go get -d -v -u -f github.com/gorilla/handlers
	go get -d -v -u -f github.com/gorilla/mux
	go get -d -v -u -f github.com/gorilla/context
	

testdeps:

	go get -d -v -t github.com/stretchr/testify/suite
	
updatetestdeps:

	go get -d -v -t -u -f github.com/stretchr/testify/suite
	
build: deps 

	go build github.com/gorilla/handlers
	go build github.com/gorilla/mux
	go build github.com/gorilla/context


install: deps
	
	go install github.com/gorilla/handlers
	go install github.com/gorilla/mux
	go install github.com/gorilla/context


test: deps testdeps 
	go test 

clean:

	go clean -i github.com/gorilla/handlers
	go clean -i github.com/gorilla/mux
	go clean -i github.com/gorilla/context
	go clean -i github.com/stretchr/testify/suite

run: install
	
	@$(TARGET)

.PHONY: \
	all \
	deps \
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	install \
	test \
	clean 

