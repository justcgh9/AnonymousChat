### Launch anonymous chat
```
$ make run
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