[WIP] KoSimMic : **Ko**ko **Sim**ple **Mic**roservice

My #002 Mini Project on learning Go.

Based on Nicholas Jackson's [tutorial](https://github.com/nicholasjackson/building-microservices-youtube)

## Run the Apps
* $ go run main.go

## Hit the endpoints
### (playing around endpoints)
* $ curl -d 'yourname' localhost:9090/hello
* $ curl localhost:9090/goodbye

### (REST endpoints)
* $ curl localhost:9090

* $ curl localhost:9090 -d '{"name": "Ice Tea", "description": "Indonesian Traditional Ice Tea", "price": 15000, "sku": "tea007"}'

* $ curl localhost:9090/1 -X PUT -d '{"name": "New Latte", "description": "New Description", "price": 250, "sku": "newlatte001"}'