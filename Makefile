run:
	go run main.go
test:
	go test -v ./... -coverpkg ./... -coverprofile cover.out
build:
	go build