build-protos:
	./cmd/protoc --go_out protos --proto_path ./ components/sync/protocol/*.proto

run:
	go run cmd/main.go

go-tidy:
	go mod tidy

