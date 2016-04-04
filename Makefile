.DEFAULT_GOAL=build

PKG := $(GOPATH)/pkg

run: generate
	go run bandit-server.go

rerun: generate
	$(GOPATH)/bin/rerun github.com/peleteiro/bandit-server

get:
	go get github.com/skelterjohn/rerun/...
	go get github.com/jteeuwen/go-bindata/...

generate: get
	$(GOPATH)/bin/go-bindata -o assets/assets.go -pkg=assets -prefix=assets -ignore=.\*.go assets

build: generate 
	@GOOS=darwin GOARCH=amd64 go build -o $(PKG)/darwin_amd64/bandit-server bandit-server.go
	@GOOS=linux GOARCH=amd64 go build -o $(PKG)/linux_amd64/bandit-server bandit-server.go
	@GOOS=linux GOARCH=386 go build -o $(PKG)/linux_386/bandit-server bandit-server.go

fmt:
	go fmt ./...

test: get
	go test ./...
