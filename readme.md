## Steps to setup

### Prerequisite
1. Install grpccurl: brew install grpccurl
2. cd token_project
3. chmod +x start_server.sh
4. chmpd +x client_commands.sh

### Steps to start server only and execute commands
1. cd into token_project
2. ./start_server.sh
3. to create a token: 
```
grpcurl -plaintext -d  '{ "id": "1234" }' localhost:50051 token.v1.TokenService/CreateToken
```
4. to write a token
```
grpcurl -plaintext -d  '{"token": {"id": "1234", "high": 10, "mid": 5, "low": 1, "name": "test_token"}}' localhost:50051 token.v1.TokenService/WriteToken
```
5. to read a token
```
grpcurl -plaintext -d  '{ "id": "1234" }' localhost:50051 token.v1.TokenService/ReadToken
```
6. to drop a token
```
grpcurl -plaintext -d  '{ "id": "1234" }' localhost:50051 token.v1.TokenService/DropToken
```
7. to stop server: 'ctlr + C' on the terminal

### Steps to start server, execute commands and stop server by itself
1. cd into token_project
2. ./client_commands.sh