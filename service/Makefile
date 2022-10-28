build:
	protoc --go_out=. --proto_path=../proto ../proto/*.proto
	protoc --go-grpc_out=. --proto_path=../proto ../proto/*.proto

clean:
	rm -r gen

run:
	$(MAKE) build
	go run src/main.go

run_client:
	go run client/main.go