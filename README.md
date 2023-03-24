# CoinListAPI

   This project provides a gRPC service for accessing cryptocurrency data using the CoinGecko API. 

# Requirements and Installation To run this project: 

   1) Go programming language to be installed. https://go.dev/doc/install
   
      To confirm that the command prints the installed version of Go.
   
          go version

   2) Protocol Buffers to be installed. https://grpc.io/docs/protoc-installation/
   
         To confirm that the command prints the installed version of protoc.
   
          protoc --version

# Setting up a gRPC-Go project:

1) Clone the repository: 

       git clone git@github.com:Yongwen6580/CoinListAPI.git

2) Navegate to the project diretory(For this project: CoinListAPI):

   Navigate to the cloned directory. Execute the following command to to enable the module.
    
       export GO111MODULE="on"
       
3) At the same directory, Installing the gRPC Go plugin:

       go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

       go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
       
4) change the directory to proto

       cd pkg/gRPC/proto
       
5) Create the proto file with the required services and messages in the proto directory(In this case coinService directory)

       protoc --go_out=. coin_service.proto
     
       protoc --go-grpc_out=. coin_service.proto
       
if you have an error showing "Please specify a program using absolute path" (skip this step if you don't have error)

    export PATH="$PATH:$(go env GOPATH)/bin"
       
6) Initialize a Go module

       go mod tidy
      

# Usage:

1) Start the server(change the directory to gRPC : /CoinListAPI/pkg/gRPC): 

      Execute the following command, if successful it will shows Starting server on port : 50053...
      
       go run server.go 

2) Use the client to call the gRPC methods:
      
      2.1) Start a new terminal and follow Setting up a gRPC-Go project step 2.
      
      2.2) Change the directory to client.
      
       cd cmd/client
            
      2.3) Follow Setting up a gRPC-Go project step 6

      2.4) Under the dictory client. Execute the following command
      
       go run client.go 

     






