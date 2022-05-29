# gRPC Training

For sharing session purposes only, this repo will evolve along with the training session

## Prerequisite

- Go version >= 1.16
- Postgresql running on local machine or you can use [Docker](https://hub.docker.com/_/postgres/)
- Install Protocol Buffer Compiler
  - Windows : <https://github.com/protocolbuffers/protobuf/releases/download/v3.20.0/protoc-3.20.0-win64.zip>
  - Linux (Debian Family, Ubuntu)  :  
  `apt install -y protobuf-compiler`
  - MacOS (Using HomeBrew) :  
  
    ```bash
    brew install protobuf
    ```

- Go plugins for the protocol compiler  

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```

Update your PATH so that the `protoc` compiler can find the plugins in Unix environment you can do this:  

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

And on windows can follow this step <https://helpdeskgeek.com/windows-10/add-windows-path-environment-variable/>

## Sample Case

We will create a simple scenario where there are 3 services with following dependencies order  

```txt
Order Service ---call--> Product Service ---call--> Inventory Service
```

## How to Run

- First, make sure the `.proto` files compiled whenever you change it, by running `make protoc` inside this root directory
- Open 3 terminals, and on each of it `cd` to `service-inventory`, `service-order` and `service-product` directory
- Change `DB_READ_URL` and `DB_WRITE_URL` on `Makefile` at `service-inventory` to your running local database
- And then on each terminal run `go mod tidy` to get the dependency if this is your 1st run, follow with `make run` to run the gRPC server for each respective service
- Try to hit the order gRPC, which will call the product and also the inventory service

## IMPORTANT NOTE

The proto buffer shared via symbolic link in each service, which is not an appropriate way, in the real implementation better expose the proto into an importable package