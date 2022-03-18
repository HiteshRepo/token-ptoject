# Setup steps

## Start server
1. go run cmd/server/main.go -configFile <file-path>
2. file is preset at token-project/config.default.yaml

## Start Client
1. go run cmd/client/main.go -configFile <file-path>
2. file is preset at token-project/config.default.yaml

### Start server and client in different terminals
### Commands (run on client terminal)
1. tokenclient -create -id 1234 -host localhost -port 50051
2. tokenclient -write -id 1234 -name abc -low 0 -mid 10 -high 100 -host localhost -port 50051
3. tokenclient -read -id 1234 -host localhost -port 50051
4. tokenclient -drop -id 1234 -host localhost -port 50051