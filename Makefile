protoc:
	@echo "Generating Proto Buffer"
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/inventory/*.proto proto/order/*.proto proto/product/*.proto

.PHONY: protoc
