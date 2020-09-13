gen:
	protoc -I. --go_out=plugins=grpc:./gen/go/ ./proto/workflow/v1/workflow.proto

clean:
	rm ./gen/go/*.go 

run:
	go run main.go