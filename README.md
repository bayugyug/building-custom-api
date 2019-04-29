## rest-building

* [x] Sample golang rest api that simulates CRUD with in-memory-storage


### Pre-Requisite
	
	- Please run this in your command line to ensure packages are in-place.
	  (normally these will be handled when compiling the api binary)
	
		go get -u -v github.com/go-chi/chi
		go get -u -v github.com/go-chi/chi/middleware
		go get -u -v github.com/go-chi/cors
		go get -u -v github.com/go-chi/render
		go get -u -v github.com/onsi/ginkgo/ginkgo
		go get -u -v github.com/onsi/gomega/... 
		go get -u -v github.com/icrowley/fake
```sh


```

### Compile

```sh

     git clone https://github.com/bayugyug/rest-building.git && cd rest-building

     git pull && make clean && make

```

### Required Preparation



### List of End-Points-Url


```go
		#get 1 row
			curl -X GET    'http://127.0.0.1:8989/v1/api/building'

			{"code":200,"status":"Success","result":[{"id":"f2b1c1b85445b3767a3d86a677247a93","name":"building-2","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:04:39+08:00"}]}

		
		
		#create
			curl -X POST    'http://127.0.0.1:8989/v1/api/building' -d '{"name":"building-a","address":"address here","floors":["floor-1","floor-2"]}'
			
			{"code":200,"status":"Success","result":"2a2527d865a9979076e3f7e62e6e21e3"}


		
		#update
		 curl -X PUT    'http://127.0.0.1:8989/v1/api/building' -d '{"id":"2a2527d865a9979076e3f7e62e6e21e3","name":"building-a","address":"address here2","floors":["floor-a1","floor-a2","floor-a3"]}'
		 
			{"code":200,"status":"Success"}


		
		#get a record
		curl -X GET    'http://127.0.0.1:8989/v1/api/building/bb752d3573ca1679be6832f73ddb4e06'
			
			{"code":200,"status":"Success","results":{"id":"bb752d3573ca1679be6832f73ddb4e06","name":"building-b","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:12:54+08:00"}}

		
		#get all records
		curl -X GET    'http://127.0.0.1:8989/v1/api/building'

			{"code":200,"status":"Success","results":[{"id":"2a2527d865a9979076e3f7e62e6e21e3","name":"building-a","address":"address here2","floors":["floor-a1","floor-a2","floor-a3"],"created":"2019-04-29T23:09:55+08:00","modified":"2019-04-29T23:11:59+08:00"},{"id":"f2b1c1b85445b3767a3d86a677247a93","name":"building-2","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:04:39+08:00"},{"id":"bb752d3573ca1679be6832f73ddb4e06","name":"building-b","address":"address here","floors":["floor-1","floor-2"],"created":"2019-04-29T23:12:54+08:00"}]}

		#delete a record
		curl -X DELETE    'http://127.0.0.1:8989/v1/api/building/bb752d3573ca1679be6832f73ddb4e06'
			
			{"code":200,"status":"Success"}


```


### Mini-How-To on running the api binary

	[x] The api can accept a json format configuration
	
	[x] Fields:
	
		- port = port to run the http server (default: 8989)
	
		- showlog   = flag for dev't log on std-out
		
	[x] Sanity check
	    
		
	
	[x] Run from the console

```sh
		./rest-building --config '{"port":"8989","showlog":true}'

```
	[x] Check the log history from the redis-cache
	



### Notes

### Reference

### License

[MIT](https://bayugyug.mit-license.org/)

