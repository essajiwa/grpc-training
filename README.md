# gRPC Training
For sharing session purposes only, this repo will evolve along with the training session

## Prerequisite
- Go version >= 1.16
- Install Protocol Buffer Compiler
    - Windows : https://github.com/protocolbuffers/protobuf/releases/download/v3.20.0/protoc-3.20.0-win64.zip
    - Linux (Debian Family, Ubuntu)  : <br />
    `apt install -y protobuf-compiler`
    - MacOS (Using HomeBrew) : <br />
    `brew install protobuf`
- Go plugins for the protocol compiler <br /> 
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Update your PATH so that the protoc compiler can find the plugins:
`export PATH="$PATH:$(go env GOPATH)/bin"`

## Sample Case
We will create a simple scenario where there are 3 services in with following dependencies order  
```
Order Service --> Product Service --> Inventory Service
```

## How to Run
- First, make sure the `.proto` files compiled whenever you change it, by running `make protoc` inside this root directory
- Open 3 terminals, and on each of it `cd` to `service-inventory`, `service-order` and `service-product` directory
- And then on each terminal run `make run` to run the gRPC server for each respective service
- Try to hit the order gRPC, which will call the product and also the inventory service

## IMPORTANT NOTE
The proto buffer shared via simbolic link in each service, which is not an appropriate way, in the real implementation better expose the proto into an importable package