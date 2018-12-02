
go run api/main.go
go run service/main.go

to update proto : protoc -I proto/ proto/car.proto  --go_out=plugins=grpc:proto


Api runs http server, 
when request comes to api, 
it contact the service via rpc.

This is example, how it in general works. Have fun!