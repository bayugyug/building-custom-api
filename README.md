## building-custom-api



* [x] Sample golang rest api that simulates CRUD with in-memory-storage



### Dependencies
- Dependencies manager: Dep - https://github.com/golang/dep
- Testing framework: ginkgo - https://github.com/onsi/ginkgo

### Compile

```sh

     git clone https://github.com/bayugyug/building-custom-api.git && cd building-custom-api

     git pull && make clean && make

```
 


### End-Points-Url


```go

#create
curl -X POST    'http://127.0.0.1:8989/v1/api/building' -d '{"name":"building-a","address":"address here","floors":["floor-1","floor-2"]}'
{"status":"success","result":"2a2527d865a9979076e3f7e62e6e21e3"}

#update
curl -X PUT    'http://127.0.0.1:8989/v1/api/building' -d '{"id":"2a2527d865a9979076e3f7e62e6e21e3","name":"building-a","address":"address here2","floors":["floor-a1","floor-a2","floor-a3"]}'
{"status":"success"}

#get a record
curl -X GET    'http://127.0.0.1:8989/v1/api/building/2a2527d865a9979076e3f7e62e6e21e3'	
{"status":"success","results":{"id":"2a2527d865a9979076e3f7e62e6e21e3","name":"building-a","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:12:54+08:00"}}

#get all records
curl -X GET    'http://127.0.0.1:8989/v1/api/building'
{"status":"success","results":[{"id":"2a2527d865a9979076e3f7e62e6e21e3","name":"building-a","address":"address here2","floors":["floor-a1","floor-a2","floor-a3"],"created":"2019-04-29T23:09:55+08:00","modified":"2019-04-29T23:11:59+08:00"},{"id":"f2b1c1b85445b3767a3d86a677247a93","name":"building-2","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:04:39+08:00"},{"id":"bb752d3573ca1679be6832f73ddb4e06","name":"building-b","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:12:54+08:00"}]}

#delete a record
curl -X DELETE    'http://127.0.0.1:8989/v1/api/building/bb752d3573ca1679be6832f73ddb4e06'
{"status":"success"}
```


### Run

- The api can accept a json format configuration
	- Fields:
		- port      = port to run the http server (default: 8989)

- Sanity check
	- Either
		- ginkgo ./...	
		- make test

- From console

```sh

./bin/building-custom-api --config '{"port":"8989"}'

```


### Notes

### Reference

### License

[MIT](https://bayugyug.mit-license.org/)

