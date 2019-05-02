BUILD_DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_TIME := $(shell date +%Y%m%d.%H%M%S)
BUILD_HASH := $(shell git log -1 2>/dev/null| head -n 1 | cut -d ' ' -f 2)
BUILD_NAME := building-custom-api
TEST_FILES := $(shell go list ./... | grep -v /vendor/)

all: build

build :
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(BUILD_NAME) -a -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-X main.BuildTime=$(BUILD_TIME) " .

test : 
	go test -v ./... > testrun.txt
	golint  $(TEST_FILES) > lint.txt
	go vet -v $(TEST_FILES) > vet.txt
	gocov test github.com/bayugyug/building-custom-api | gocov-xml > coverage.xml
	go test $(TEST_FILES) -bench=. -test.benchmem -v 2>/dev/null | gobench2plot > benchmarks.xml
	ginkgo -v  ./... > gink.txt

testginkgo : build
	ginkgo -v  ./...

testrun : clean test
	time go test -v -bench=. -benchmem -dummy > testrun.txt 2>&1

prepare : build

clean:
	rm -f $(BUILD_NAME) bin/$(BUILD_NAME)
	rm -f benchmarks.xml coverage.xml vet.txt lint.txt testrun.txt gink.txt

re: clean all

