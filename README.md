# Protobuffers

This module saves all protobuffers of the project, run the following commands to generate the files in GO:

- protoc --go_out=. --go-grpc_out=. information/*.proto
- protoc -I. --go_out=paths=source_relative:./order --go-grpc_out=paths=source_relative:./order order/*.proto
- protoc -I. --go_out=paths=source_relative:./address --go-grpc_out=paths=source_relative:./address address/*.proto
- protoc --go_out=. --go-grpc_out=. classification/*.proto
- python -m grpc_tools.protoc -I. --grpc_python_out=. --python_out=. classification/*.proto 
