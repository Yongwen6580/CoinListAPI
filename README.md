# CoinListAPI

This project provides a gRPC service for accessing cryptocurrency data using the CoinGecko API. 

# Requirements To run this project: 

you will need Go programming language to be installed. https://go.dev/doc/install

# Installation:

1) Clone the repository: 

   Git clone git@github.com:Yongwen6580/CoinListAPI.git

2) Install dependencies:

   Navigate to the cloned directory and navigate to the coinService folder : cd coinService

    Execute the following two command to generate gRPC dependencies.
    protoc --go_out=. Coin_service.proto
    protoc --go-grpc_out=. coin_service.proto 

# Usage:

1) Start the server: go run server/server.go 

      or Navigate to server.go and click the run button

2) Use the client to call the gRPC methods: go run client/client.go 

      Or navigate to client.go and click the run button
 







