clean:
	@go clean .

build: clean
	@go build -ldflags "-X main.buildVersion=`date "+%Y-%m-%d_%H-%M.%S"`" .
