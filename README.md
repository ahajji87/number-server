## Number-server

### Description
number-server is a sample go websocket server that handle numbers with 9 digits length.
It handles by default:
- 5 concurrent client connections
- Save received numbers on file ensuring no duplicity
- Every 10 seconds by default it print a report of new add, duplicated numbers received during this 10 secs and the total sum of all numbers received during runtime.
- server can be shutdown by client calling endpoint /terminate

### Requirements
This application was created under go version 1.12, un use Go dependency management tool golang/dep 
- go version <=1.12
- go dep - [installation guid](https://github.com/golang/dep)

## Usage
After colin the repo to your computer, follow next steps:

```
dep ensure
```

In the config.yml we have the default configuration

- **server.port**: apllication port, by default 4000
- **server.timeouts** read and write:  server time outs to read and write 
- **server.ws**: hold the configuration of websocket engine
- **ws.connections**: max concurrent websocket connections 
- **ws.endpoint**: the enpoint for handling the received numbers
- **server.shutdown**: hold the configuration for the shutdown by api requst
- **shutdown.endpoint**: endpoint name to let client shutdown the server,
- **shutdown.timeout**: time of graceful shutdown
- **storage**: hold the information of storage
- **path**: root to file where we save the numbers
- **app.report.time**: Report log time, by default each 10 seconds the application will log a report about handled numbers during this time.

#### Build
After installing the dependencies and modifing the config file lets build the application with command:

```
go build -o ./build/number-server .
```

#### Run
To run the application just run the build with the command:

```
./bild/number-server
```

In case you want to run the application with a diferent config run the following command targiting your configuration file: 
```
./build/number-server -config=PATH TO CONFIG YML FILE
```

To use the application send a string number with 9 digits length exp: "123456789" to endoint:
```
ws:\\localhost:4000\ws
```

To terminate the application:
```
ws:\\localhost:4000\terminate
```
##Tests:
To test the application use:

``` 
go test -v number-server/app/usercase
go test -v number-server/app/usercase
```
###Benchmark
There is a basic benchmark test in the path, in the root project tap:

```
go test -bench=. -benchmem
```  
There is a step tests for [Artillery](https://artillery.io/docs/getting-started/) to do a load test, tap:
```
artillery run loadben.yml 
```  
