# Project structure
Client - /web folder \
Server - other folders (cmd - the entrypoints of the application, internal - incapsulated functionality, database - storages functonality, test - testing)

## Launch anonymous chat
```
$ make run
```
## Run tests
```
$ make tests
```
## Run tests coverage
```
$ make tests-cover
```
## Run formatters
```
$ make lint
```
## Run message count profiler
```
$ make profile
```
## Configuration environment variables
`SERVER_ADDRESS` - chat running address
`SQLITE_PATH` - sqlite file db path
`PROFILE_REQUEST_COUNT` - count of request for testing `/messages/count`

## Install required dependencies
```
$ go mod tidy
```
