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

1) Start the server: 

      Execute the following two command
      
         go run server/server.go 

     OR navigate to server.go and click the run button

2) Use the client to call the gRPC methods:

      Execute the following two command
      
         go run client/client.go 

      OR navigate to client.go and click the run button
 







