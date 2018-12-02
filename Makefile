all: build

.PHONY: deps build init
deps:
	go get -d -v github.com/dustin/go-broadcast/...
	go get -d -v github.com/manucorporat/stats/...
	go get -d -v github.com/go-redis/redis
	go get -d -v github.com/gin-gonic/gin

# init redis
init:
	go run init/init.go

build: deps
	go build -o cics2
