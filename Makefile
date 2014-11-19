all: deps test install
	@echo "--> Cross compiling"
	@goxc

deps:
	@echo "--> Installing dependencies"
	@go get ./...

test:
	@echo "--> Running unit tests"
	@go test -cover ./...

install:
	@echo "--> Installing"
	@go install -v

publish:
	@echo "--> Publising to Bintray"
	@goxc bintray