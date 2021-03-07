# Weather App Server

## Running without any setup
If you just want to run the server without installing anything, you may simply download the binary under the build folder (build/weather-app) and execute locally. In Linux/Mac, if you download the entire build folder, run from the command line:

```
./build/weather-app
```

It should start the server listening to port 8080.

## Running and cloning locally
If you want to clone and run the project locally, you may do so by following the steps below.

### Pre-requisites
* [Go](https://golang.org/dl/) version 1.16 or higher. 
* If you do not want to install Go locally, you may use the [official Docker image](https://hub.docker.com/_/golang) for Go.

### Getting started

#### Clone the repo to your favorite development directory
```
git clone https://github.com/gitterchris/weather-app-server.git
cd weather-app-server
```

### Run
```
go run ./cmd/...
```
