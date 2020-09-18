gen:
	protoc -I. --go_out=plugins=grpc:$GOPATH/src ./proto/workflow/v1/workflow.proto
	protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:. \          
  ./proto/workflow/v1/workflow.proto

clean:
	rm ./gen/go/*.go 

run:
	go run main.go